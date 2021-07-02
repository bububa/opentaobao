package alihealthpw

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthPwApplynodeUpdatenameAPIRequest 回调变更患者姓名 API请求
// alibaba.alihealth.pw.applynode.updatename
//
// 回调变更患者姓名
type AlibabaAlihealthPwApplynodeUpdatenameAPIRequest struct {
	model.Params
	// 回调入参
	_body *ModifyNameRo
}

// NewAlibabaAlihealthPwApplynodeUpdatenameRequest 初始化AlibabaAlihealthPwApplynodeUpdatenameAPIRequest对象
func NewAlibabaAlihealthPwApplynodeUpdatenameRequest() *AlibabaAlihealthPwApplynodeUpdatenameAPIRequest {
	return &AlibabaAlihealthPwApplynodeUpdatenameAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthPwApplynodeUpdatenameAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.pw.applynode.updatename"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthPwApplynodeUpdatenameAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Body Setter
// 回调入参
func (r *AlibabaAlihealthPwApplynodeUpdatenameAPIRequest) SetBody(_body *ModifyNameRo) error {
	r._body = _body
	r.Set("body", _body)
	return nil
}

// Get Body Getter
func (r AlibabaAlihealthPwApplynodeUpdatenameAPIRequest) GetBody() *ModifyNameRo {
	return r._body
}
