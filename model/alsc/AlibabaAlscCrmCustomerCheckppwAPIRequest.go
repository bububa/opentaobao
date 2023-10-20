package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmcustomercheckppwAPIRequest 校验支付密码 API请求
// alibaba.alsc.crm.customer.checkppw
//
// 校验支付密码
type AlibabaalsccrmcustomercheckppwAPIRequest struct {
	model.Params
	// 请求参数
	_checkRequest *CheckPayPasswdReq
}

// NewAlibabaalsccrmcustomercheckppwRequest 初始化AlibabaalsccrmcustomercheckppwAPIRequest对象
func NewAlibabaalsccrmcustomercheckppwRequest() *AlibabaalsccrmcustomercheckppwAPIRequest {
	return &AlibabaalsccrmcustomercheckppwAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalsccrmcustomercheckppwAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.customer.checkppw"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalsccrmcustomercheckppwAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalsccrmcustomercheckppwAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCheckRequest is CheckRequest Setter
// 请求参数
func (r *AlibabaalsccrmcustomercheckppwAPIRequest) SetCheckRequest(_checkRequest *CheckPayPasswdReq) error {
	r._checkRequest = _checkRequest
	r.Set("check_request", _checkRequest)
	return nil
}

// GetCheckRequest CheckRequest Getter
func (r AlibabaalsccrmcustomercheckppwAPIRequest) GetCheckRequest() *CheckPayPasswdReq {
	return r._checkRequest
}
