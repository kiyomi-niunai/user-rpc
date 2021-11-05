package model

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	usersFieldNames          = builderx.RawFieldNames(&Users{})
	usersRows                = strings.Join(usersFieldNames, ",")
	usersRowsExpectAutoSet   = strings.Join(stringx.Remove(usersFieldNames, "`create_time`", "`update_time`"), ",")
	usersRowsWithPlaceHolder = strings.Join(stringx.Remove(usersFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheXBackendUsersIdPrefix = "cache:xBackend:users:id:"
)

type (
	UsersModel interface {
		Insert(data Users) (sql.Result, error)
		FindOne(id int64) (*Users, error)
		Update(data Users) error
		Delete(id int64) error
	}

	defaultUsersModel struct {
		sqlc.CachedConn
		table string
	}

	Users struct {
		Id           int64          `db:"id"`
		Name         string         `db:"name"`
		Oid          sql.NullString `db:"oid"`
		Id1          sql.NullString `db:"id1"`
		Id2          sql.NullString `db:"id2"`
		Balance      int64          `db:"balance"`
		Level        int64          `db:"level"`
		Xp           int64          `db:"xp"`
		CreatedAt    sql.NullTime   `db:"created_at"`
		UpdatedAt    sql.NullTime   `db:"updated_at"`
		SsSyncAt     sql.NullTime   `db:"ss_sync_at"`
		Avatar       sql.NullString `db:"avatar"`
		AvatarName   sql.NullString `db:"avatar_name"` // 头像名
		VipLevel     int64          `db:"vip_level"`
		VipXp        int64          `db:"vip_xp"`
		LastLoginAt  sql.NullTime   `db:"last_login_at"`
		Source       int64          `db:"source"`
		FbId         sql.NullInt64  `db:"fb_id"`
		FbToken      sql.NullString `db:"fb_token"`
		Ext          sql.NullString `db:"ext"`
		MoneyBox     int64          `db:"money_box"`
		Break        int64          `db:"break"`
		Inbox        sql.NullString `db:"inbox"`
		Shop         sql.NullString `db:"shop"`
		Task         sql.NullString `db:"task"`
		BigWinTimes  int64          `db:"big_win_times"`
		JackpotTimes int64          `db:"jackpot_times"`
		TotalWin     int64          `db:"total_win"`
		BiggestWin   int64          `db:"biggest_win"`
		SubExpiredAt sql.NullTime   `db:"sub_expired_at"`
		Extend       sql.NullString `db:"extend"`   // 扩展字段
		Baggage      sql.NullString `db:"baggage"`  // 背包
		Card         sql.NullString `db:"card"`     // 卡牌
		AppleId      sql.NullString `db:"apple_id"` // apple id
		Mail         string         `db:"mail"`     // 用户绑定邮箱
	}
)

func NewUsersModel(conn sqlx.SqlConn, c cache.CacheConf) UsersModel {
	return &defaultUsersModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`users`",
	}
}

func (m *defaultUsersModel) Insert(data Users) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, usersRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.Id, data.Name, data.Oid, data.Id1, data.Id2, data.Balance, data.Level, data.Xp, data.CreatedAt, data.UpdatedAt, data.SsSyncAt, data.Avatar, data.AvatarName, data.VipLevel, data.VipXp, data.LastLoginAt, data.Source, data.FbId, data.FbToken, data.Ext, data.MoneyBox, data.Break, data.Inbox, data.Shop, data.Task, data.BigWinTimes, data.JackpotTimes, data.TotalWin, data.BiggestWin, data.SubExpiredAt, data.Extend, data.Baggage, data.Card, data.AppleId, data.Mail)

	return ret, err
}

func (m *defaultUsersModel) FindOne(id int64) (*Users, error) {
	xBackendUsersIdKey := fmt.Sprintf("%s%v", cacheXBackendUsersIdPrefix, id)
	var resp Users
	err := m.QueryRow(&resp, xBackendUsersIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", usersRows, m.table)
		return conn.QueryRow(v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUsersModel) Update(data Users) error {
	xBackendUsersIdKey := fmt.Sprintf("%s%v", cacheXBackendUsersIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, usersRowsWithPlaceHolder)
		return conn.Exec(query, data.Name, data.Oid, data.Id1, data.Id2, data.Balance, data.Level, data.Xp, data.CreatedAt, data.UpdatedAt, data.SsSyncAt, data.Avatar, data.AvatarName, data.VipLevel, data.VipXp, data.LastLoginAt, data.Source, data.FbId, data.FbToken, data.Ext, data.MoneyBox, data.Break, data.Inbox, data.Shop, data.Task, data.BigWinTimes, data.JackpotTimes, data.TotalWin, data.BiggestWin, data.SubExpiredAt, data.Extend, data.Baggage, data.Card, data.AppleId, data.Mail, data.Id)
	}, xBackendUsersIdKey)
	return err
}

func (m *defaultUsersModel) Delete(id int64) error {

	xBackendUsersIdKey := fmt.Sprintf("%s%v", cacheXBackendUsersIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, xBackendUsersIdKey)
	return err
}

func (m *defaultUsersModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheXBackendUsersIdPrefix, primary)
}

func (m *defaultUsersModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", usersRows, m.table)
	return conn.QueryRow(v, query, primary)
}
