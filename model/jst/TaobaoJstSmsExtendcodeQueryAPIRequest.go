package jst

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstSmsExtendcodeQueryAPIRequest 聚石塔扩展码查询 API请求
// taobao.jst.sms.extendcode.query
//
// 聚石塔扩展码查询
type TaobaoJstSmsExtendcodeQueryAPIRequest struct {
	model.Params
	// 扩展码查询请求
	_extendCodeQueryRequest *ExtendCodeQueryRequest
}

// NewTaobaoJstSmsExtendcodeQueryRequest 初始化TaobaoJstSmsExtendcodeQueryAPIRequest对象
func NewTaobaoJstSmsExtendcodeQueryRequest() *TaobaoJstSmsExtendcodeQueryAPIRequest {
	return &TaobaoJstSmsExtendcodeQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJstSmsExtendcodeQueryAPIRequest) GetApiMethodName() string {
	return "taobao.jst.sms.extendcode.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJstSmsExtendcodeQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetExtendCodeQueryRequest is ExtendCodeQueryRequest Setter
// 扩展码查询请求
func (r *TaobaoJstSmsExtendcodeQueryAPIRequest) SetExtendCodeQueryRequest(_extendCodeQueryRequest *ExtendCodeQueryRequest) error {
	r._extendCodeQueryRequest = _extendCodeQueryRequest
	r.Set("extend_code_query_request", _extendCodeQueryRequest)
	return nil
}

// GetExtendCodeQueryRequest ExtendCodeQueryRequest Getter
func (r TaobaoJstSmsExtendcodeQueryAPIRequest) GetExtendCodeQueryRequest() *ExtendCodeQueryRequest {
	return r._extendCodeQueryRequest
}
