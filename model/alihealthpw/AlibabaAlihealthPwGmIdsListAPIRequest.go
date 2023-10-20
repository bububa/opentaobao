package alihealthpw

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthPwGmIdsListAPIRequest 同情用药根据申请单列表查询申请单 API请求
// alibaba.alihealth.pw.gm.ids.list
//
// 同情用药根据申请单列表查询申请单
type AlibabaAlihealthPwGmIdsListAPIRequest struct {
	model.Params
	// 入参
	_body *ListByApplyIdsForBreq
}

// NewAlibabaAlihealthPwGmIdsListRequest 初始化AlibabaAlihealthPwGmIdsListAPIRequest对象
func NewAlibabaAlihealthPwGmIdsListRequest() *AlibabaAlihealthPwGmIdsListAPIRequest {
	return &AlibabaAlihealthPwGmIdsListAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthPwGmIdsListAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.pw.gm.ids.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthPwGmIdsListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthPwGmIdsListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBody is Body Setter
// 入参
func (r *AlibabaAlihealthPwGmIdsListAPIRequest) SetBody(_body *ListByApplyIdsForBreq) error {
	r._body = _body
	r.Set("body", _body)
	return nil
}

// GetBody Body Getter
func (r AlibabaAlihealthPwGmIdsListAPIRequest) GetBody() *ListByApplyIdsForBreq {
	return r._body
}
