package legalsuit

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalSuitCourtEntrustGetAPIRequest 委托开庭服务查询 API请求
// alibaba.legal.suit.court.entrust.get
//
// 查询委托开庭信息
type AlibabaLegalSuitCourtEntrustGetAPIRequest struct {
	model.Params
	// 委托ID
	_entrustId int64
	// 案件ID
	_suitId int64
}

// NewAlibabaLegalSuitCourtEntrustGetRequest 初始化AlibabaLegalSuitCourtEntrustGetAPIRequest对象
func NewAlibabaLegalSuitCourtEntrustGetRequest() *AlibabaLegalSuitCourtEntrustGetAPIRequest {
	return &AlibabaLegalSuitCourtEntrustGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLegalSuitCourtEntrustGetAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.suit.court.entrust.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLegalSuitCourtEntrustGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
