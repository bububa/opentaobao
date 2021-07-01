package jstinteractive

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJstInteractiveAssetsQueryAPIRequest
查询可配置任务素材接口 API请求
taobao.jst.interactive.assets.query

查询可配置任务素材列表，用以配置任务素材 */
type TaobaoJstInteractiveAssetsQueryAPIRequest struct {
	model.Params
	// 小程序id
	_miniAppId string
}

// New
