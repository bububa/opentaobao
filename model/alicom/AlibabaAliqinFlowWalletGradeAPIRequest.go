package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaaliqinflowwalletgradeAPIRequest 获取流量档位 API请求
// alibaba.aliqin.flow.wallet.grade
//
// 获取直充流量档位
type AlibabaaliqinflowwalletgradeAPIRequest struct {
	model.Params
	// 手机号码
	_phoneNum string
}

// NewAlibabaaliqinflowwalletgradeRequest 初始化AlibabaaliqinflowwalletgradeAPIRequest对象
func NewAlibabaaliqinflowwalletgradeRequest() *AlibabaaliqinflowwalletgradeAPIRequest {
	return &AlibabaaliqinflowwalletgradeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaaliqinflowwalletgradeAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.flow.wallet.grade"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaaliqinflowwalletgradeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaaliqinflowwalletgradeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPhoneNum is PhoneNum Setter
// 手机号码
func (r *AlibabaaliqinflowwalletgradeAPIRequest) SetPhoneNum(_phoneNum string) error {
	r._phoneNum = _phoneNum
	r.Set("phone_num", _phoneNum)
	return nil
}

// GetPhoneNum PhoneNum Getter
func (r AlibabaaliqinflowwalletgradeAPIRequest) GetPhoneNum() string {
	return r._phoneNum
}
