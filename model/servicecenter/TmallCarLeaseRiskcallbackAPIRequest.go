package servicecenter

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallcarleaseriskcallbackAPIRequest 整车租赁风控模型回调 API请求
// tmall.car.lease.riskcallback
//
// 租赁公司回调风控结果
type TmallcarleaseriskcallbackAPIRequest struct {
	model.Params
	// 授信结果
	_creditInfo *CreditInfoTopDto
}

// NewTmallcarleaseriskcallbackRequest 初始化TmallcarleaseriskcallbackAPIRequest对象
func NewTmallcarleaseriskcallbackRequest() *TmallcarleaseriskcallbackAPIRequest {
	return &TmallcarleaseriskcallbackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallcarleaseriskcallbackAPIRequest) GetApiMethodName() string {
	return "tmall.car.lease.riskcallback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallcarleaseriskcallbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallcarleaseriskcallbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCreditInfo is CreditInfo Setter
// 授信结果
func (r *TmallcarleaseriskcallbackAPIRequest) SetCreditInfo(_creditInfo *CreditInfoTopDto) error {
	r._creditInfo = _creditInfo
	r.Set("credit_info", _creditInfo)
	return nil
}

// GetCreditInfo CreditInfo Getter
func (r TmallcarleaseriskcallbackAPIRequest) GetCreditInfo() *CreditInfoTopDto {
	return r._creditInfo
}
