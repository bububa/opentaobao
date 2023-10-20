package alihealthpw

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthPwApplynodeUpdatenameAPIRequest) Reset() {
	r._body = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthPwApplynodeUpdatenameAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.pw.applynode.updatename"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthPwApplynodeUpdatenameAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthPwApplynodeUpdatenameAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBody is Body Setter
// 回调入参
func (r *AlibabaAlihealthPwApplynodeUpdatenameAPIRequest) SetBody(_body *ModifyNameRo) error {
	r._body = _body
	r.Set("body", _body)
	return nil
}

// GetBody Body Getter
func (r AlibabaAlihealthPwApplynodeUpdatenameAPIRequest) GetBody() *ModifyNameRo {
	return r._body
}

var poolAlibabaAlihealthPwApplynodeUpdatenameAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthPwApplynodeUpdatenameRequest()
	},
}

// GetAlibabaAlihealthPwApplynodeUpdatenameRequest 从 sync.Pool 获取 AlibabaAlihealthPwApplynodeUpdatenameAPIRequest
func GetAlibabaAlihealthPwApplynodeUpdatenameAPIRequest() *AlibabaAlihealthPwApplynodeUpdatenameAPIRequest {
	return poolAlibabaAlihealthPwApplynodeUpdatenameAPIRequest.Get().(*AlibabaAlihealthPwApplynodeUpdatenameAPIRequest)
}

// ReleaseAlibabaAlihealthPwApplynodeUpdatenameAPIRequest 将 AlibabaAlihealthPwApplynodeUpdatenameAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthPwApplynodeUpdatenameAPIRequest(v *AlibabaAlihealthPwApplynodeUpdatenameAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthPwApplynodeUpdatenameAPIRequest.Put(v)
}
