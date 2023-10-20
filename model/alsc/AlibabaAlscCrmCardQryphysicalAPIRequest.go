package alsc

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscCrmCardQryphysicalAPIRequest 查询物理卡 API请求
// alibaba.alsc.crm.card.qryphysical
//
// 查询物理卡
type AlibabaAlscCrmCardQryphysicalAPIRequest struct {
	model.Params
	// 入参
	_paramQueryPhyCardOpenReq *QueryPhyCardOpenReq
}

// NewAlibabaAlscCrmCardQryphysicalRequest 初始化AlibabaAlscCrmCardQryphysicalAPIRequest对象
func NewAlibabaAlscCrmCardQryphysicalRequest() *AlibabaAlscCrmCardQryphysicalAPIRequest {
	return &AlibabaAlscCrmCardQryphysicalAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlscCrmCardQryphysicalAPIRequest) Reset() {
	r._paramQueryPhyCardOpenReq = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmCardQryphysicalAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.crm.card.qryphysical"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmCardQryphysicalAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscCrmCardQryphysicalAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamQueryPhyCardOpenReq is ParamQueryPhyCardOpenReq Setter
// 入参
func (r *AlibabaAlscCrmCardQryphysicalAPIRequest) SetParamQueryPhyCardOpenReq(_paramQueryPhyCardOpenReq *QueryPhyCardOpenReq) error {
	r._paramQueryPhyCardOpenReq = _paramQueryPhyCardOpenReq
	r.Set("param_query_phy_card_open_req", _paramQueryPhyCardOpenReq)
	return nil
}

// GetParamQueryPhyCardOpenReq ParamQueryPhyCardOpenReq Getter
func (r AlibabaAlscCrmCardQryphysicalAPIRequest) GetParamQueryPhyCardOpenReq() *QueryPhyCardOpenReq {
	return r._paramQueryPhyCardOpenReq
}

var poolAlibabaAlscCrmCardQryphysicalAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlscCrmCardQryphysicalRequest()
	},
}

// GetAlibabaAlscCrmCardQryphysicalRequest 从 sync.Pool 获取 AlibabaAlscCrmCardQryphysicalAPIRequest
func GetAlibabaAlscCrmCardQryphysicalAPIRequest() *AlibabaAlscCrmCardQryphysicalAPIRequest {
	return poolAlibabaAlscCrmCardQryphysicalAPIRequest.Get().(*AlibabaAlscCrmCardQryphysicalAPIRequest)
}

// ReleaseAlibabaAlscCrmCardQryphysicalAPIRequest 将 AlibabaAlscCrmCardQryphysicalAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlscCrmCardQryphysicalAPIRequest(v *AlibabaAlscCrmCardQryphysicalAPIRequest) {
	v.Reset()
	poolAlibabaAlscCrmCardQryphysicalAPIRequest.Put(v)
}
