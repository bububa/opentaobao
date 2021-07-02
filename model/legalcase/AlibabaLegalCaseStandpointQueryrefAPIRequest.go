package legalcase

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLegalCaseStandpointQueryrefAPIRequest 查询推送口径信息 API请求
// alibaba.legal.case.standpoint.queryref
//
// 查询推送口径信息
type AlibabaLegalCaseStandpointQueryrefAPIRequest struct {
	model.Params
	// 案件id
	_suitId int64
	// 委托id
	_entrustId int64
	// 查询的口径id列表
	_standpointIds []int64
}

// NewAlibabaLegalCaseStandpointQueryrefRequest 初始化AlibabaLegalCaseStandpointQueryrefAPIRequest对象
func NewAlibabaLegalCaseStandpointQueryrefRequest() *AlibabaLegalCaseStandpointQueryrefAPIRequest {
	return &AlibabaLegalCaseStandpointQueryrefAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLegalCaseStandpointQueryrefAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.case.standpoint.queryref"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLegalCaseStandpointQueryrefAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is SuitId Setter
// 案件id
func (r *AlibabaLegalCaseStandpointQueryrefAPIRequest) SetSuitId(_suitId int64) error {
	r._suitId = _suitId
	r.Set("suit_id", _suitId)
	return nil
}

// Get SuitId Getter
func (r AlibabaLegalCaseStandpointQueryrefAPIRequest) GetSuitId() int64 {
	return r._suitId
}

// Set is EntrustId Setter
// 委托id
func (r *AlibabaLegalCaseStandpointQueryrefAPIRequest) SetEntrustId(_entrustId int64) error {
	r._entrustId = _entrustId
	r.Set("entrust_id", _entrustId)
	return nil
}

// Get EntrustId Getter
func (r AlibabaLegalCaseStandpointQueryrefAPIRequest) GetEntrustId() int64 {
	return r._entrustId
}

// Set is StandpointIds Setter
// 查询的口径id列表
func (r *AlibabaLegalCaseStandpointQueryrefAPIRequest) SetStandpointIds(_standpointIds []int64) error {
	r._standpointIds = _standpointIds
	r.Set("standpoint_ids", _standpointIds)
	return nil
}

// Get StandpointIds Getter
func (r AlibabaLegalCaseStandpointQueryrefAPIRequest) GetStandpointIds() []int64 {
	return r._standpointIds
}
