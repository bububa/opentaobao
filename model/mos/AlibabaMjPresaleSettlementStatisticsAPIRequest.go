package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMjPresaleSettlementStatisticsAPIRequest 预购结算数据统计 API请求
// alibaba.mj.presale.settlement.statistics
//
// 预购结算数据统计
type AlibabaMjPresaleSettlementStatisticsAPIRequest struct {
	model.Params
	// 外部门店编码
	_storeNo string
	// 活动期号
	_actionNo int64
}

// NewAlibabaMjPresaleSettlementStatisticsRequest 初始化AlibabaMjPresaleSettlementStatisticsAPIRequest对象
func NewAlibabaMjPresaleSettlementStatisticsRequest() *AlibabaMjPresaleSettlementStatisticsAPIRequest {
	return &AlibabaMjPresaleSettlementStatisticsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMjPresaleSettlementStatisticsAPIRequest) GetApiMethodName() string {
	return "alibaba.mj.presale.settlement.statistics"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMjPresaleSettlementStatisticsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMjPresaleSettlementStatisticsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreNo is StoreNo Setter
// 外部门店编码
func (r *AlibabaMjPresaleSettlementStatisticsAPIRequest) SetStoreNo(_storeNo string) error {
	r._storeNo = _storeNo
	r.Set("store_no", _storeNo)
	return nil
}

// GetStoreNo StoreNo Getter
func (r AlibabaMjPresaleSettlementStatisticsAPIRequest) GetStoreNo() string {
	return r._storeNo
}

// SetActionNo is ActionNo Setter
// 活动期号
func (r *AlibabaMjPresaleSettlementStatisticsAPIRequest) SetActionNo(_actionNo int64) error {
	r._actionNo = _actionNo
	r.Set("action_no", _actionNo)
	return nil
}

// GetActionNo ActionNo Getter
func (r AlibabaMjPresaleSettlementStatisticsAPIRequest) GetActionNo() int64 {
	return r._actionNo
}
