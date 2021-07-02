package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoJstAstrolabeOrderstatusSyncAPIRequest 线下门店派单以及单据相关操作接口 API请求
// taobao.jst.astrolabe.orderstatus.sync
//
// 针对ERP系统部署在门店的商家，将派单状态回流到星盘
type TaobaoJstAstrolabeOrderstatusSyncAPIRequest struct {
	model.Params
	// 子订单Id
	_subOrderIds []int64
	// 事件发生时间
	_actionTime string
	// 操作人
	_operator string
	// 业务类型
	_type string
	// 订单状态
	_status string
	// 目标门店的商户中心门店编码
	_storeId int64
	// 交易订单
	_parentOrderCode int64
}

// NewTaobaoJstAstrolabeOrderstatusSyncRequest 初始化TaobaoJstAstrolabeOrderstatusSyncAPIRequest对象
func NewTaobaoJstAstrolabeOrderstatusSyncRequest() *TaobaoJstAstrolabeOrderstatusSyncAPIRequest {
	return &TaobaoJstAstrolabeOrderstatusSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoJstAstrolabeOrderstatusSyncAPIRequest) GetApiMethodName() string {
	return "taobao.jst.astrolabe.orderstatus.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoJstAstrolabeOrderstatusSyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetSubOrderIds is SubOrderIds Setter
// 子订单Id
func (r *TaobaoJstAstrolabeOrderstatusSyncAPIRequest) SetSubOrderIds(_subOrderIds []int64) error {
	r._subOrderIds = _subOrderIds
	r.Set("sub_order_ids", _subOrderIds)
	return nil
}

// GetSubOrderIds SubOrderIds Getter
func (r TaobaoJstAstrolabeOrderstatusSyncAPIRequest) GetSubOrderIds() []int64 {
	return r._subOrderIds
}

// SetActionTime is ActionTime Setter
// 事件发生时间
func (r *TaobaoJstAstrolabeOrderstatusSyncAPIRequest) SetActionTime(_actionTime string) error {
	r._actionTime = _actionTime
	r.Set("action_time", _actionTime)
	return nil
}

// GetActionTime ActionTime Getter
func (r TaobaoJstAstrolabeOrderstatusSyncAPIRequest) GetActionTime() string {
	return r._actionTime
}

// SetOperator is Operator Setter
// 操作人
func (r *TaobaoJstAstrolabeOrderstatusSyncAPIRequest) SetOperator(_operator string) error {
	r._operator = _operator
	r.Set("operator", _operator)
	return nil
}

// GetOperator Operator Getter
func (r TaobaoJstAstrolabeOrderstatusSyncAPIRequest) GetOperator() string {
	return r._operator
}

// SetType is Type Setter
// 业务类型
func (r *TaobaoJstAstrolabeOrderstatusSyncAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TaobaoJstAstrolabeOrderstatusSyncAPIRequest) GetType() string {
	return r._type
}

// SetStatus is Status Setter
// 订单状态
func (r *TaobaoJstAstrolabeOrderstatusSyncAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r TaobaoJstAstrolabeOrderstatusSyncAPIRequest) GetStatus() string {
	return r._status
}

// SetStoreId is StoreId Setter
// 目标门店的商户中心门店编码
func (r *TaobaoJstAstrolabeOrderstatusSyncAPIRequest) SetStoreId(_storeId int64) error {
	r._storeId = _storeId
	r.Set("store_id", _storeId)
	return nil
}

// GetStoreId StoreId Getter
func (r TaobaoJstAstrolabeOrderstatusSyncAPIRequest) GetStoreId() int64 {
	return r._storeId
}

// SetParentOrderCode is ParentOrderCode Setter
// 交易订单
func (r *TaobaoJstAstrolabeOrderstatusSyncAPIRequest) SetParentOrderCode(_parentOrderCode int64) error {
	r._parentOrderCode = _parentOrderCode
	r.Set("parent_order_code", _parentOrderCode)
	return nil
}

// GetParentOrderCode ParentOrderCode Getter
func (r TaobaoJstAstrolabeOrderstatusSyncAPIRequest) GetParentOrderCode() int64 {
	return r._parentOrderCode
}
