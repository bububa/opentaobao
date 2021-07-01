package jstinteractive

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoJstInteractiveAssetsConfigAPIRequest
任务素材配置接口 API请求
taobao.jst.interactive.assets.config

任务素材配置接口 */
type TaobaoJstInteractiveAssetsConfigAPIRequest struct {
	model.Params
	// []
	_assetsConfigList []AssetsConfig
	// 小程序id
	_miniAppId string
}

// New
