package jstinteractive

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJstInteractiveActivityQueryAPIRequest
互动任务活动查询接口 API请求
taobao.jst.interactive.activity.query

互动任务活动查询接口 */
type TaobaoJstInteractiveActivityQueryAPIRequest struct {
	model.Params
	// 小程序id
	_miniAppId string
}

// New
