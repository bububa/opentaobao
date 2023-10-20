package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamjpresalesettlementstatisticsAPIRequest 预购结算数据统计 API请求
// alibaba.mj.presale.settlement.statistics
//
// 预购结算数据统计
type AlibabamjpresalesettlementstatisticsAPIRequest struct {
	model.Params
	// 外部门店编码
	_storeNo string
	// 活动期号
	_actionNo int64
}

// NewAlibabamjpresalesettlementstatisticsRequest 初始化AlibabamjpresalesettlementstatisticsAPIRequest对象
func NewAlibabamjpresalesettlementstatisticsRequest() *AlibabamjpresalesettlementstatisticsAPIRequest {
	return &AlibabamjpresalesettlementstatisticsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamjpresalesettlementstatisticsAPIRequest) GetApiMethodName() string {
	return "alibaba.mj.presale.settlement.statistics"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamjpresalesettlementstatisticsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamjpresalesettlementstatisticsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreNo is StoreNo Setter
// 外部门店编码
func (r *AlibabamjpresalesettlementstatisticsAPIRequest) SetStoreNo(_storeNo string) error {
	r._storeNo = _storeNo
	r.Set("store_no", _storeNo)
	return nil
}

// GetStoreNo StoreNo Getter
func (r AlibabamjpresalesettlementstatisticsAPIRequest) GetStoreNo() string {
	return r._storeNo
}

// SetActionNo is ActionNo Setter
// 活动期号
func (r *AlibabamjpresalesettlementstatisticsAPIRequest) SetActionNo(_actionNo int64) error {
	r._actionNo = _actionNo
	r.Set("action_no", _actionNo)
	return nil
}

// GetActionNo ActionNo Getter
func (r AlibabamjpresalesettlementstatisticsAPIRequest) GetActionNo() int64 {
	return r._actionNo
}
