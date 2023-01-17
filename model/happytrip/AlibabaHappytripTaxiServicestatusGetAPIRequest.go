package happytrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHappytripTaxiServicestatusGetAPIRequest 供应商服务开通状态 API请求
// alibaba.happytrip.taxi.servicestatus.get
//
// 获取服务供应商在每个地区的服务开通状态、支持的车型等
type AlibabaHappytripTaxiServicestatusGetAPIRequest struct {
	model.Params
	// 成本中心代码，用于区分不同的分账账号
	_costCenter string
}

// NewAlibabaHappytripTaxiServicestatusGetRequest 初始化AlibabaHappytripTaxiServicestatusGetAPIRequest对象
func NewAlibabaHappytripTaxiServicestatusGetRequest() *AlibabaHappytripTaxiServicestatusGetAPIRequest {
	return &AlibabaHappytripTaxiServicestatusGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHappytripTaxiServicestatusGetAPIRequest) GetApiMethodName() string {
	return "alibaba.happytrip.taxi.servicestatus.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHappytripTaxiServicestatusGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHappytripTaxiServicestatusGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCostCenter is CostCenter Setter
// 成本中心代码，用于区分不同的分账账号
func (r *AlibabaHappytripTaxiServicestatusGetAPIRequest) SetCostCenter(_costCenter string) error {
	r._costCenter = _costCenter
	r.Set("cost_center", _costCenter)
	return nil
}

// GetCostCenter CostCenter Getter
func (r AlibabaHappytripTaxiServicestatusGetAPIRequest) GetCostCenter() string {
	return r._costCenter
}
