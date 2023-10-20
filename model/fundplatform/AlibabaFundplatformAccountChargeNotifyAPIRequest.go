package fundplatform

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabafundplatformaccountchargenotifyAPIRequest 账户充值成功通知 API请求
// alibaba.fundplatform.account.charge.notify
//
// 通知外部业务方充值成功
type AlibabafundplatformaccountchargenotifyAPIRequest struct {
	model.Params
	// 入参对象
	_request *AlibabafundplatformaccountchargenotifyStruct
}

// NewAlibabafundplatformaccountchargenotifyRequest 初始化AlibabafundplatformaccountchargenotifyAPIRequest对象
func NewAlibabafundplatformaccountchargenotifyRequest() *AlibabafundplatformaccountchargenotifyAPIRequest {
	return &AlibabafundplatformaccountchargenotifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabafundplatformaccountchargenotifyAPIRequest) GetApiMethodName() string {
	return "alibaba.fundplatform.account.charge.notify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabafundplatformaccountchargenotifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabafundplatformaccountchargenotifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
// 入参对象
func (r *AlibabafundplatformaccountchargenotifyAPIRequest) SetRequest(_request *AlibabafundplatformaccountchargenotifyStruct) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r AlibabafundplatformaccountchargenotifyAPIRequest) GetRequest() *AlibabafundplatformaccountchargenotifyStruct {
	return r._request
}
