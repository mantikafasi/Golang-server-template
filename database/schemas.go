package database

import (
	"context"

	"github.com/uptrace/bun"
)

type UserBadge struct { //leaved as example
	bun.BaseModel `bun:"table:userbadges"`

	ID               int32  `bun:"id,pk,autoincrement" json:"-"`
	DiscordID        string `bun:"discordid,type:numeric" json:"-"`
	BadgeName        string `bun:"badge_name" json:"badge_name"`
	BadgeIcon        string `bun:"badge_icon" json:"badge_icon"`
	RedirectURL      string `bun:"redirect_url" json:"redirect_url"`
	BadgeType        int32  `bun:"badge_type" json:"badge_type"`
	BadgeDescription string `bun:"badge_description" json:"badge_description"`
}

func createSchema() error {
	models := []any{
		//list of models
	}

	for _, model := range models {
		if _, err := DB.NewCreateTable().IfNotExists().Model(model).Exec(context.Background()); err != nil {
			return err
		}
	}
	return nil
}
