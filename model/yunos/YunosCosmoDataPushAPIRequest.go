package yunos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosCosmoDataPushAPIRequest
COSMO-PUSH模式数据接入 API请求
yunos.cosmo.data.push

YunOS提供外部数据源接入，并输出到多端设备上，该接口提供了PUSH模式的数据接入 */
type YunosCosmoDataPushAPIRequest struct {
	model.Params
	// 业务方数据源唯一标识，由COSMO平台颁发
	_appId string
	// 业务方推送数据，List结构的JSON序列化字符串
	_jsonModel string
}

// New
