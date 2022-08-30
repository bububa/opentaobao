package alihealthpw

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthPwGmDetailAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.pw.gm.detail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthPwGmDetailAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
