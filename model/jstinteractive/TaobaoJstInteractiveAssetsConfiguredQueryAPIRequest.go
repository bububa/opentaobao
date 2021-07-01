package jstinteractive

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJstInteractiveAssetsConfiguredQueryAPIRequest
查询已配置的任务素材列表接口 API请求
taobao.jst.interactive.assets.configured.query

查询已配置任务素材列表 */
type TaobaoJstInteractiveAssetsConfiguredQueryAPIRequest struct {
	model.Params
	// 小程序id
	_miniAppId string
}

// New
