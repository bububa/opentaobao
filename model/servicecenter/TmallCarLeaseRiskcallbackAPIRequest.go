package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallCarLeaseRiskcallbackAPIRequest 整车租赁风控模型回调 API请求
// tmall.car.lease.riskcallback
//
// 租赁公司回调风控结果
type TmallCarLeaseRiskcallbackAPIRequest struct {
	model.Params
	// 授信结果
	_creditInfo *CreditInfoTopDto
}

// NewTmallCarLeaseRiskcallbackRequest 初始化TmallCarLeaseRiskcallbackAPIRequest对象
func NewTmallCarLeaseRiskcallbackRequest() *TmallCarLeaseRiskcallbackAPIRequest {
	return &TmallCarLeaseRiskcallbackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallCarLeaseRiskcallbackAPIRequest) GetApiMethodName() string {
	return "tmall.car.lease.riskcallback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallCarLeaseRiskcallbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallCarLeaseRiskcallbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCreditInfo is CreditInfo Setter
// 授信结果
func (r *TmallCarLeaseRiskcallbackAPIRequest) SetCreditInfo(_creditInfo *CreditInfoTopDto) error {
	r._creditInfo = _creditInfo
	r.Set("credit_info", _creditInfo)
	return nil
}

// GetCreditInfo CreditInfo Getter
func (r TmallCarLeaseRiskcallbackAPIRequest) GetCreditInfo() *CreditInfoTopDto {
	return r._creditInfo
}
