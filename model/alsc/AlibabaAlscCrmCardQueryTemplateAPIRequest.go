package alsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmCardQueryTemplateAPIRequest 查询卡模板详情 API请求
// alibaba.alsc.crm.card.query.template
//
// 查询卡模板详情
type AlibabaAlscCrmCardQueryTemplateAPIRequest struct {
	model.Params
	// 请求对象
	_paramQueryCardTemplateOpenReq *QueryCardTemplateOpenReq
}

// NewAlibabaAlscCrmCardQueryTemplateRequest 初始化AlibabaAlscCrmCardQueryTemplateAPIRequest对象
func NewAlibabaAlscCrmCardQueryTemplateRequest() *AlibabaAlscCrmCardQueryTemplateAPIRequest {
	return &AlibabaAlscCrmCardQueryTemplateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlscCrmCardQueryTemplateAPIRequest) Reset() {
	r._paramQueryCardTemplateOpenReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmCardQueryTemplateAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.card.query.template"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmCardQueryTemplateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscCrmCardQueryTemplateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamQueryCardTemplateOpenReq is ParamQueryCardTemplateOpenReq Setter
// 请求对象
func (r *AlibabaAlscCrmCardQueryTemplateAPIRequest) SetParamQueryCardTemplateOpenReq(_paramQueryCardTemplateOpenReq *QueryCardTemplateOpenReq) error {
	r._paramQueryCardTemplateOpenReq = _paramQueryCardTemplateOpenReq
	r.Set("param_query_card_template_open_req", _paramQueryCardTemplateOpenReq)
	return nil
}

// GetParamQueryCardTemplateOpenReq ParamQueryCardTemplateOpenReq Getter
func (r AlibabaAlscCrmCardQueryTemplateAPIRequest) GetParamQueryCardTemplateOpenReq() *QueryCardTemplateOpenReq {
	return r._paramQueryCardTemplateOpenReq
}

var poolAlibabaAlscCrmCardQueryTemplateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlscCrmCardQueryTemplateRequest()
	},
}

// GetAlibabaAlscCrmCardQueryTemplateRequest 从 sync.Pool 获取 AlibabaAlscCrmCardQueryTemplateAPIRequest
func GetAlibabaAlscCrmCardQueryTemplateAPIRequest() *AlibabaAlscCrmCardQueryTemplateAPIRequest {
	return poolAlibabaAlscCrmCardQueryTemplateAPIRequest.Get().(*AlibabaAlscCrmCardQueryTemplateAPIRequest)
}

// ReleaseAlibabaAlscCrmCardQueryTemplateAPIRequest 将 AlibabaAlscCrmCardQueryTemplateAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlscCrmCardQueryTemplateAPIRequest(v *AlibabaAlscCrmCardQueryTemplateAPIRequest) {
	v.Reset()
	poolAlibabaAlscCrmCardQueryTemplateAPIRequest.Put(v)
}
