package baodian

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoDegUserGamegiftQueryAPIRequest
用户数娱游戏礼包查询 API请求
taobao.deg.user.gamegift.query

查询用户数娱礼包列表 */
type TaobaoDegUserGamegiftQueryAPIRequest struct {
	model.Params
	// 状态，1为待发放，2为已发放，3为过期
	_status int64
	// cp item id列表
	_cpItemIds []string
}

// New
