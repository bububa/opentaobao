package legalcase

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalegalcasestandpointqueryrefAPIRequest 查询推送口径信息 API请求
// alibaba.legal.case.standpoint.queryref
//
// 查询推送口径信息
type AlibabalegalcasestandpointqueryrefAPIRequest struct {
	model.Params
	// 查询的口径id列表
	_standpointIds []string
	// 案件id
	_suitId int64
	// 委托id
	_entrustId int64
}

// NewAlibabalegalcasestandpointqueryrefRequest 初始化AlibabalegalcasestandpointqueryrefAPIRequest对象
func NewAlibabalegalcasestandpointqueryrefRequest() *AlibabalegalcasestandpointqueryrefAPIRequest {
	return &AlibabalegalcasestandpointqueryrefAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalegalcasestandpointqueryrefAPIRequest) GetApiMethodName() string {
	return "alibaba.legal.case.standpoint.queryref"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalegalcasestandpointqueryrefAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalegalcasestandpointqueryrefAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStandpointIds is StandpointIds Setter
// 查询的口径id列表
func (r *AlibabalegalcasestandpointqueryrefAPIRequest) SetStandpointIds(_standpointIds []string) error {
	r._standpointIds = _standpointIds
	r.Set("standpoint_ids", _standpointIds)
	return nil
}

// GetStandpointIds StandpointIds Getter
func (r AlibabalegalcasestandpointqueryrefAPIRequest) GetStandpointIds() []string {
	return r._standpointIds
}

// SetSuitId is SuitId Setter
// 案件id
func (r *AlibabalegalcasestandpointqueryrefAPIRequest) SetSuitId(_suitId int64) error {
	r._suitId = _suitId
	r.Set("suit_id", _suitId)
	return nil
}

// GetSuitId SuitId Getter
func (r AlibabalegalcasestandpointqueryrefAPIRequest) GetSuitId() int64 {
	return r._suitId
}

// SetEntrustId is EntrustId Setter
// 委托id
func (r *AlibabalegalcasestandpointqueryrefAPIRequest) SetEntrustId(_entrustId int64) error {
	r._entrustId = _entrustId
	r.Set("entrust_id", _entrustId)
	return nil
}

// GetEntrustId EntrustId Getter
func (r AlibabalegalcasestandpointqueryrefAPIRequest) GetEntrustId() int64 {
	return r._entrustId
}
