package cainiaohandover

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoGlobalLogisticsCarrierQuerylistAPIRequest
实际承运商查询 API请求
cainiao.global.logistics.carrier.querylist

查询出所有的实际承运商 */
type CainiaoGlobalLogisticsCarrierQuerylistAPIRequest struct {
	model.Params
	// 多语言(暂不支持，保留入参)
	_locale string
}

// New
