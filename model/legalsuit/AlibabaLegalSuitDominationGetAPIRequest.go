package legalsuit

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalSuitDominationGetAPIRequest 查询管辖信息 API请求
// alibaba.legal.suit.domination.get
//
// 查询管辖信息
type AlibabaLegalSuitDominationGetAPIRequest struct {
	model.Params
	// 案件ID
	_suitId int64
	// 委托ID
	_entrustId int64
	// 管辖ID
	_dominationId int64
}

// NewAlibabaLegalSuitDominationGetRequest 初始化AlibabaLegalSuitDominationGetAPIRequest对象
func NewAlibabaLegalSuitDominationGetRequest() *AlibabaLegalSuitDominationGetAPIRequest {
	return &AlibabaLegalSuitDominationGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLegalSuitDominationGetAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.suit.domination.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLegalSuitDominationGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLegalSuitDominationGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSuitId is SuitId Setter
// 案件ID
func (r *AlibabaLegalSuitDominationGetAPIRequest) SetSuitId(_suitId int64) error {
	r._suitId = _suitId
	r.Set("suit_id", _suitId)
	return nil
}

// GetSuitId SuitId Getter
func (r AlibabaLegalSuitDominationGetAPIRequest) GetSuitId() int64 {
	return r._suitId
}

// SetEntrustId is EntrustId Setter
// 委托ID
func (r *AlibabaLegalSuitDominationGetAPIRequest) SetEntrustId(_entrustId int64) error {
	r._entrustId = _entrustId
	r.Set("entrust_id", _entrustId)
	return nil
}

// GetEntrustId EntrustId Getter
func (r AlibabaLegalSuitDominationGetAPIRequest) GetEntrustId() int64 {
	return r._entrustId
}

// SetDominationId is DominationId Setter
// 管辖ID
func (r *AlibabaLegalSuitDominationGetAPIRequest) SetDominationId(_dominationId int64) error {
	r._dominationId = _dominationId
	r.Set("domination_id", _dominationId)
	return nil
}

// GetDominationId DominationId Getter
func (r AlibabaLegalSuitDominationGetAPIRequest) GetDominationId() int64 {
	return r._dominationId
}
