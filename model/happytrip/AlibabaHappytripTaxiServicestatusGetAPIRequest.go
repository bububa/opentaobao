package happytrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaHappytripTaxiServicestatusGetAPIRequest
供应商服务开通状态 API请求
alibaba.happytrip.taxi.servicestatus.get

获取服务供应商在每个地区的服务开通状态、支持的车型等 */
type AlibabaHappytripTaxiServicestatusGetAPIRequest struct {
	model.Params
	// 成本中心代码，用于区分不同的分账账号
	_costCenter string
}

// New
