package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users"`
	Id            uuid.UUID  `bun:"type:uuid(36),pk"`          // ユーザーID
	Email         string     `bun:"type:varchar(255),notnull"` // メールアドレス
	Password      string     `bun:"type:varchar(255),notnull"` // パスワード
	Name          string     `bun:"type:varchar(255),notnull"` // ユーザー名
	CreatedAt     *time.Time `bun:"type:timestamp"`            // 作成日時
	UpdatedAt     *time.Time `bun:"type:timestamp"`            // 更新日時
}
