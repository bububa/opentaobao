package film

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFilmKoubeiCinemaGetactivityAPIRequest
口碑-影院营销数据查询 API请求
taobao.film.koubei.cinema.getactivity

口碑-影院营销数据查询 */
type TaobaoFilmKoubeiCinemaGetactivityAPIRequest struct {
	model.Params
	// 用户账号
	_userId string
	// 账号类型
	_accountType string
	// 城市编码
	_cityCode int64
	// 平台
	_platform int64
	// 影院ID集合
	_cinemaIds []int64
	// 附加参数
	_params string
}

// New
