package baodian

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBaodianServerSdkConfigGetAPIRequest
获取宝点SDK的配置项（已迁移） API请求
taobao.baodian.server.sdk.config.get

获取SDK各种配置项（已迁移） */
type TaobaoBaodianServerSdkConfigGetAPIRequest struct {
	model.Params
	// appKey
	_appkey string
	// 渠道
	_channel string
	// sdk版本号
	_sdkVer string
	// 与后端约定
	_type int64
}

// New
