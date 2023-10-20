package alsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmCardBindcardAPIRequest 绑定物理卡 API请求
// alibaba.alsc.crm.card.bindcard
//
// 将会员卡和实例物理卡绑定在一起
type AlibabaAlscCrmCardBindcardAPIRequest struct {
	model.Params
	// 请求参数
	_paramBindPhysicalCardOpenReq *BindPhysicalCardOpenReq
}

// NewAlibabaAlscCrmCardBindcardRequest 初始化AlibabaAlscCrmCardBindcardAPIRequest对象
func NewAlibabaAlscCrmCardBindcardRequest() *AlibabaAlscCrmCardBindcardAPIRequest {
	return &AlibabaAlscCrmCardBindcardAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlscCrmCardBindcardAPIRequest) Reset() {
	r._paramBindPhysicalCardOpenReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmCardBindcardAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.card.bindcard"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmCardBindcardAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscCrmCardBindcardAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBindPhysicalCardOpenReq is ParamBindPhysicalCardOpenReq Setter
// 请求参数
func (r *AlibabaAlscCrmCardBindcardAPIRequest) SetParamBindPhysicalCardOpenReq(_paramBindPhysicalCardOpenReq *BindPhysicalCardOpenReq) error {
	r._paramBindPhysicalCardOpenReq = _paramBindPhysicalCardOpenReq
	r.Set("param_bind_physical_card_open_req", _paramBindPhysicalCardOpenReq)
	return nil
}

// GetParamBindPhysicalCardOpenReq ParamBindPhysicalCardOpenReq Getter
func (r AlibabaAlscCrmCardBindcardAPIRequest) GetParamBindPhysicalCardOpenReq() *BindPhysicalCardOpenReq {
	return r._paramBindPhysicalCardOpenReq
}

var poolAlibabaAlscCrmCardBindcardAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlscCrmCardBindcardRequest()
	},
}

// GetAlibabaAlscCrmCardBindcardRequest 从 sync.Pool 获取 AlibabaAlscCrmCardBindcardAPIRequest
func GetAlibabaAlscCrmCardBindcardAPIRequest() *AlibabaAlscCrmCardBindcardAPIRequest {
	return poolAlibabaAlscCrmCardBindcardAPIRequest.Get().(*AlibabaAlscCrmCardBindcardAPIRequest)
}

// ReleaseAlibabaAlscCrmCardBindcardAPIRequest 将 AlibabaAlscCrmCardBindcardAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlscCrmCardBindcardAPIRequest(v *AlibabaAlscCrmCardBindcardAPIRequest) {
	v.Reset()
	poolAlibabaAlscCrmCardBindcardAPIRequest.Put(v)
}
