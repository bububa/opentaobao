package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJstMiniappCrowdUserAddAPIRequest
小程序添加用户到指定的活动 API请求
taobao.jst.miniapp.crowd.user.add

小程序添加用户到指定的活动 */
type TaobaoJstMiniappCrowdUserAddAPIRequest struct {
	model.Params
	// 活动code
	_crowdCode string
	// 小程序id
	_mcGwSourceMiniAppId string
	// 小程序appkey
	_mcGwSourceAppKey string
}

// New
