package legalsuit

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalSuitCourtEntrustGetAPIRequest 委托开庭服务查询 API请求
// alibaba.legal.suit.court.entrust.get
//
// 查询委托开庭信息
type AlibabaLegalSuitCourtEntrustGetAPIRequest struct {
	model.Params
	// 案件ID
	_suitId int64
	// 委托ID
	_entrustId int64
}

// NewAlibabaLegalSuitCourtEntrustGetRequest 初始化AlibabaLegalSuitCourtEntrustGetAPIRequest对象
func NewAlibabaLegalSuitCourtEntrustGetRequest() *AlibabaLegalSuitCourtEntrustGetAPIRequest {
	return &AlibabaLegalSuitCourtEntrustGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaLegalSuitCourtEntrustGetAPIRequest) Reset() {
	r._suitId = 0
	r._entrustId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLegalSuitCourtEntrustGetAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.suit.court.entrust.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLegalSuitCourtEntrustGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLegalSuitCourtEntrustGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSuitId is SuitId Setter
// 案件ID
func (r *AlibabaLegalSuitCourtEntrustGetAPIRequest) SetSuitId(_suitId int64) error {
	r._suitId = _suitId
	r.Set("suit_id", _suitId)
	return nil
}

// GetSuitId SuitId Getter
func (r AlibabaLegalSuitCourtEntrustGetAPIRequest) GetSuitId() int64 {
	return r._suitId
}

// SetEntrustId is EntrustId Setter
// 委托ID
func (r *AlibabaLegalSuitCourtEntrustGetAPIRequest) SetEntrustId(_entrustId int64) error {
	r._entrustId = _entrustId
	r.Set("entrust_id", _entrustId)
	return nil
}

// GetEntrustId EntrustId Getter
func (r AlibabaLegalSuitCourtEntrustGetAPIRequest) GetEntrustId() int64 {
	return r._entrustId
}

var poolAlibabaLegalSuitCourtEntrustGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaLegalSuitCourtEntrustGetRequest()
	},
}

// GetAlibabaLegalSuitCourtEntrustGetRequest 从 sync.Pool 获取 AlibabaLegalSuitCourtEntrustGetAPIRequest
func GetAlibabaLegalSuitCourtEntrustGetAPIRequest() *AlibabaLegalSuitCourtEntrustGetAPIRequest {
	return poolAlibabaLegalSuitCourtEntrustGetAPIRequest.Get().(*AlibabaLegalSuitCourtEntrustGetAPIRequest)
}

// ReleaseAlibabaLegalSuitCourtEntrustGetAPIRequest 将 AlibabaLegalSuitCourtEntrustGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaLegalSuitCourtEntrustGetAPIRequest(v *AlibabaLegalSuitCourtEntrustGetAPIRequest) {
	v.Reset()
	poolAlibabaLegalSuitCourtEntrustGetAPIRequest.Put(v)
}
