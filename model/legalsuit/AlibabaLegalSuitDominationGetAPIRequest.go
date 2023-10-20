package legalsuit

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalegalsuitdominationgetAPIRequest 查询管辖信息 API请求
// alibaba.legal.suit.domination.get
//
// 查询管辖信息
type AlibabalegalsuitdominationgetAPIRequest struct {
	model.Params
	// 案件ID
	_suitId int64
	// 委托ID
	_entrustId int64
	// 管辖ID
	_dominationId int64
}

// NewAlibabalegalsuitdominationgetRequest 初始化AlibabalegalsuitdominationgetAPIRequest对象
func NewAlibabalegalsuitdominationgetRequest() *AlibabalegalsuitdominationgetAPIRequest {
	return &AlibabalegalsuitdominationgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalegalsuitdominationgetAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.suit.domination.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalegalsuitdominationgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalegalsuitdominationgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSuitId is SuitId Setter
// 案件ID
func (r *AlibabalegalsuitdominationgetAPIRequest) SetSuitId(_suitId int64) error {
	r._suitId = _suitId
	r.Set("suit_id", _suitId)
	return nil
}

// GetSuitId SuitId Getter
func (r AlibabalegalsuitdominationgetAPIRequest) GetSuitId() int64 {
	return r._suitId
}

// SetEntrustId is EntrustId Setter
// 委托ID
func (r *AlibabalegalsuitdominationgetAPIRequest) SetEntrustId(_entrustId int64) error {
	r._entrustId = _entrustId
	r.Set("entrust_id", _entrustId)
	return nil
}

// GetEntrustId EntrustId Getter
func (r AlibabalegalsuitdominationgetAPIRequest) GetEntrustId() int64 {
	return r._entrustId
}

// SetDominationId is DominationId Setter
// 管辖ID
func (r *AlibabalegalsuitdominationgetAPIRequest) SetDominationId(_dominationId int64) error {
	r._dominationId = _dominationId
	r.Set("domination_id", _dominationId)
	return nil
}

// GetDominationId DominationId Getter
func (r AlibabalegalsuitdominationgetAPIRequest) GetDominationId() int64 {
	return r._dominationId
}
