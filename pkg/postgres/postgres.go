package postgres

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgconn/stmtcache"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

func NewPoolConnection(
	ctx context.Context,
	db,
	dbAddr,
	user,
	password string,
	maxIdleLifetime,
	maxLifetime time.Duration,
	prepareCacheCap,
	maxConn int,
) (*pgxpool.Pool, error) {
	// TODO add logger
	addr, port, err := parseDbAddressAndPort(dbAddr)
	cfg, err := pgxpool.ParseConfig(fmt.Sprintf(
		"host=%s port=%d dbname=%s sslmode=disable user=%s password=%s pool_max_conns=%d",
		addr, port, db, user, password, maxConn,
	))
	if err != nil {
		return nil, errors.Wrapf(err, "failed parse postgres dsn: %s:%v", dbAddr, dbAddr)
	}

	cfg.MaxConnIdleTime = maxIdleLifetime
	cfg.MaxConnLifetime = maxLifetime

	cfg.ConnConfig.BuildStatementCache = func(conn *pgconn.PgConn) stmtcache.Cache {
		return stmtcache.New(conn, stmtcache.ModeDescribe, prepareCacheCap)
	}

	pool, err := pgxpool.ConnectConfig(ctx, cfg)
	if err != nil {
		return nil, errors.Wrapf(err, "failed connect to pg: %s:%v", addr, port)
	}

	return pool, nil
}

func parseDbAddressAndPort(conn string) (string, int, error) {
	splits := strings.Split(conn, ":")
	address := splits[0]
	port, err := strconv.Atoi(splits[1])
	if err != nil {
		return "", 0, errors.Wrapf(err, "failed parse db address and port, connection string=%s", conn)
	}
	return address, port, nil
}
