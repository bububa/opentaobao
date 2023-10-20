package mei

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallmeicrmmembergetbypaycodeAPIRequest 支付码获取会员信息 API请求
// tmall.mei.crm.member.getbypaycode
//
// 通过支付码获取会员信息
type TmallmeicrmmembergetbypaycodeAPIRequest struct {
	model.Params
	// 会员码
	_payCode string
}

// NewTmallmeicrmmembergetbypaycodeRequest 初始化TmallmeicrmmembergetbypaycodeAPIRequest对象
func NewTmallmeicrmmembergetbypaycodeRequest() *TmallmeicrmmembergetbypaycodeAPIRequest {
	return &TmallmeicrmmembergetbypaycodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallmeicrmmembergetbypaycodeAPIRequest) GetApiMethodName() string {
	return "tmall.mei.crm.member.getbypaycode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallmeicrmmembergetbypaycodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallmeicrmmembergetbypaycodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPayCode is PayCode Setter
// 会员码
func (r *TmallmeicrmmembergetbypaycodeAPIRequest) SetPayCode(_payCode string) error {
	r._payCode = _payCode
	r.Set("pay_code", _payCode)
	return nil
}

// GetPayCode PayCode Getter
func (r TmallmeicrmmembergetbypaycodeAPIRequest) GetPayCode() string {
	return r._payCode
}
