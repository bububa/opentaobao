package btrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripBtripCostCenterModifyAPIRequest
修改外部成本中心 API请求
alitrip.btrip.cost.center.modify

修改外部成本中心，设置成员，设置支付宝账号，设置名称，编号等 */
type AlitripBtripCostCenterModifyAPIRequest struct {
	model.Params
	// 请求对象
	_rq *OpenCostCenterModifyRq
}

// NewAlitripBtripCostCenterModifyRequest 初始化AlitripBtripCostCenterModifyAPIRequest对象
func NewAlitripBtripCostCenterModifyRequest() *AlitripBtripCostCenterModifyAPIRequest {
	return &AlitripBtripCostCenterModifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripBtripCostCenterModifyAPIRequest) GetApiMethodName() string {
	return "alitrip.btrip.cost.center.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripBtripCostCenterModifyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Rq Setter
// 请求对象
func (r *AlitripBtripCostCenterModifyAPIRequest) SetRq(_rq *OpenCostCenterModifyRq) error {
	r._rq = _rq
	r.Set("rq", _rq)
	return nil
}

// Get Rq Getter
func (r AlitripBtripCostCenterModifyAPIRequest) GetRq() *OpenCostCenterModifyRq {
	return r._rq
}
