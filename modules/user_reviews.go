package modules

import (
	"context"
	"server-go/common"
	"server-go/database"

	"github.com/patrickmn/go-cache"
)

func GetAllBadges() (badges []database.UserBadge, err error) {

	cachedBadges, found := common.Cache.Get("badges")
	if found {
		badges = cachedBadges.([]database.UserBadge)
		return
	}

	badges = []database.UserBadge{}
	err = database.DB.NewSelect().Model(&badges).Scan(context.Background(), &badges)

	common.Cache.Set("badges", badges, cache.DefaultExpiration)
	return
}
