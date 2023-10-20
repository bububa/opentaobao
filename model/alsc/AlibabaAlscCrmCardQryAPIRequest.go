package alsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmCardQryAPIRequest 查询卡实例 API请求
// alibaba.alsc.crm.card.qry
//
// 查询卡实例（优先使用卡实例ID查询，没有则用物理卡号查询）
type AlibabaAlscCrmCardQryAPIRequest struct {
	model.Params
	// 请求对象
	_paramQueryCardOpenReq *QueryCardOpenReq
}

// NewAlibabaAlscCrmCardQryRequest 初始化AlibabaAlscCrmCardQryAPIRequest对象
func NewAlibabaAlscCrmCardQryRequest() *AlibabaAlscCrmCardQryAPIRequest {
	return &AlibabaAlscCrmCardQryAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlscCrmCardQryAPIRequest) Reset() {
	r._paramQueryCardOpenReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmCardQryAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.card.qry"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmCardQryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscCrmCardQryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamQueryCardOpenReq is ParamQueryCardOpenReq Setter
// 请求对象
func (r *AlibabaAlscCrmCardQryAPIRequest) SetParamQueryCardOpenReq(_paramQueryCardOpenReq *QueryCardOpenReq) error {
	r._paramQueryCardOpenReq = _paramQueryCardOpenReq
	r.Set("param_query_card_open_req", _paramQueryCardOpenReq)
	return nil
}

// GetParamQueryCardOpenReq ParamQueryCardOpenReq Getter
func (r AlibabaAlscCrmCardQryAPIRequest) GetParamQueryCardOpenReq() *QueryCardOpenReq {
	return r._paramQueryCardOpenReq
}

var poolAlibabaAlscCrmCardQryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlscCrmCardQryRequest()
	},
}

// GetAlibabaAlscCrmCardQryRequest 从 sync.Pool 获取 AlibabaAlscCrmCardQryAPIRequest
func GetAlibabaAlscCrmCardQryAPIRequest() *AlibabaAlscCrmCardQryAPIRequest {
	return poolAlibabaAlscCrmCardQryAPIRequest.Get().(*AlibabaAlscCrmCardQryAPIRequest)
}

// ReleaseAlibabaAlscCrmCardQryAPIRequest 将 AlibabaAlscCrmCardQryAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlscCrmCardQryAPIRequest(v *AlibabaAlscCrmCardQryAPIRequest) {
	v.Reset()
	poolAlibabaAlscCrmCardQryAPIRequest.Put(v)
}
