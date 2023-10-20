package alihealthpw

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthPwGmDetailAPIRequest 同情用药申请单详情接口 API请求
// alibaba.alihealth.pw.gm.detail
//
// 同情用药申请单详情接口，提供给合作方查询申请单详情
type AlibabaAlihealthPwGmDetailAPIRequest struct {
	model.Params
	// 入参
	_body *DetailForBReq
}

// NewAlibabaAlihealthPwGmDetailRequest 初始化AlibabaAlihealthPwGmDetailAPIRequest对象
func NewAlibabaAlihealthPwGmDetailRequest() *AlibabaAlihealthPwGmDetailAPIRequest {
	return &AlibabaAlihealthPwGmDetailAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthPwGmDetailAPIRequest) Reset() {
	r._body = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthPwGmDetailAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.pw.gm.detail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthPwGmDetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthPwGmDetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBody is Body Setter
// 入参
func (r *AlibabaAlihealthPwGmDetailAPIRequest) SetBody(_body *DetailForBReq) error {
	r._body = _body
	r.Set("body", _body)
	return nil
}

// GetBody Body Getter
func (r AlibabaAlihealthPwGmDetailAPIRequest) GetBody() *DetailForBReq {
	return r._body
}

var poolAlibabaAlihealthPwGmDetailAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthPwGmDetailRequest()
	},
}

// GetAlibabaAlihealthPwGmDetailRequest 从 sync.Pool 获取 AlibabaAlihealthPwGmDetailAPIRequest
func GetAlibabaAlihealthPwGmDetailAPIRequest() *AlibabaAlihealthPwGmDetailAPIRequest {
	return poolAlibabaAlihealthPwGmDetailAPIRequest.Get().(*AlibabaAlihealthPwGmDetailAPIRequest)
}

// ReleaseAlibabaAlihealthPwGmDetailAPIRequest 将 AlibabaAlihealthPwGmDetailAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthPwGmDetailAPIRequest(v *AlibabaAlihealthPwGmDetailAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthPwGmDetailAPIRequest.Put(v)
}
