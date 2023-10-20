package wlbimports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaowlbimportsgeneralconsignAPIRequest 一般进口发货 API请求
// taobao.wlb.imports.general.consign
//
// 将订单信息发送到菜鸟海外转运仓；
// 业务规则：
// 1）交易订单为待发货状态。
// 2）单笔订单多个商品，交易金额不能大于1000人民币。
type TaobaowlbimportsgeneralconsignAPIRequest struct {
	model.Params
	// 仓库编码
	_storeCode string
	// 第1段物流承运商
	_firstLogistics string
	// 第1段物流运单号
	_firstWaybillno string
	// 增值服务编码.多个以逗号分隔
	_vasCode string
	// 交易订单id
	_tradeOrderId int64
	// 物流资源ID
	_resourceId int64
	// 卖家发货地址库ID；不填的话取默认发货地址；如果填写的senderId对应多个地址，取第一个
	_senderId int64
	// 卖家退货地址库ID；不填写的话取默认退货地址；如果填写的cancelId对应多个地址，取第一个
	_cancelId int64
}

// NewTaobaowlbimportsgeneralconsignRequest 初始化TaobaowlbimportsgeneralconsignAPIRequest对象
func NewTaobaowlbimportsgeneralconsignRequest() *TaobaowlbimportsgeneralconsignAPIRequest {
	return &TaobaowlbimportsgeneralconsignAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaowlbimportsgeneralconsignAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.imports.general.consign"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaowlbimportsgeneralconsignAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaowlbimportsgeneralconsignAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetStoreCode is StoreCode Setter
// 仓库编码
func (r *TaobaowlbimportsgeneralconsignAPIRequest) SetStoreCode(_storeCode string) error {
	r._storeCode = _storeCode
	r.Set("store_code", _storeCode)
	return nil
}

// GetStoreCode StoreCode Getter
func (r TaobaowlbimportsgeneralconsignAPIRequest) GetStoreCode() string {
	return r._storeCode
}

// SetFirstLogistics is FirstLogistics Setter
// 第1段物流承运商
func (r *TaobaowlbimportsgeneralconsignAPIRequest) SetFirstLogistics(_firstLogistics string) error {
	r._firstLogistics = _firstLogistics
	r.Set("first_logistics", _firstLogistics)
	return nil
}

// GetFirstLogistics FirstLogistics Getter
func (r TaobaowlbimportsgeneralconsignAPIRequest) GetFirstLogistics() string {
	return r._firstLogistics
}

// SetFirstWaybillno is FirstWaybillno Setter
// 第1段物流运单号
func (r *TaobaowlbimportsgeneralconsignAPIRequest) SetFirstWaybillno(_firstWaybillno string) error {
	r._firstWaybillno = _firstWaybillno
	r.Set("first_waybillno", _firstWaybillno)
	return nil
}

// GetFirstWaybillno FirstWaybillno Getter
func (r TaobaowlbimportsgeneralconsignAPIRequest) GetFirstWaybillno() string {
	return r._firstWaybillno
}

// SetVasCode is VasCode Setter
// 增值服务编码.多个以逗号分隔
func (r *TaobaowlbimportsgeneralconsignAPIRequest) SetVasCode(_vasCode string) error {
	r._vasCode = _vasCode
	r.Set("vas_code", _vasCode)
	return nil
}

// GetVasCode VasCode Getter
func (r TaobaowlbimportsgeneralconsignAPIRequest) GetVasCode() string {
	return r._vasCode
}

// SetTradeOrderId is TradeOrderId Setter
// 交易订单id
func (r *TaobaowlbimportsgeneralconsignAPIRequest) SetTradeOrderId(_tradeOrderId int64) error {
	r._tradeOrderId = _tradeOrderId
	r.Set("trade_order_id", _tradeOrderId)
	return nil
}

// GetTradeOrderId TradeOrderId Getter
func (r TaobaowlbimportsgeneralconsignAPIRequest) GetTradeOrderId() int64 {
	return r._tradeOrderId
}

// SetResourceId is ResourceId Setter
// 物流资源ID
func (r *TaobaowlbimportsgeneralconsignAPIRequest) SetResourceId(_resourceId int64) error {
	r._resourceId = _resourceId
	r.Set("resource_id", _resourceId)
	return nil
}

// GetResourceId ResourceId Getter
func (r TaobaowlbimportsgeneralconsignAPIRequest) GetResourceId() int64 {
	return r._resourceId
}

// SetSenderId is SenderId Setter
// 卖家发货地址库ID；不填的话取默认发货地址；如果填写的senderId对应多个地址，取第一个
func (r *TaobaowlbimportsgeneralconsignAPIRequest) SetSenderId(_senderId int64) error {
	r._senderId = _senderId
	r.Set("sender_id", _senderId)
	return nil
}

// GetSenderId SenderId Getter
func (r TaobaowlbimportsgeneralconsignAPIRequest) GetSenderId() int64 {
	return r._senderId
}

// SetCancelId is CancelId Setter
// 卖家退货地址库ID；不填写的话取默认退货地址；如果填写的cancelId对应多个地址，取第一个
func (r *TaobaowlbimportsgeneralconsignAPIRequest) SetCancelId(_cancelId int64) error {
	r._cancelId = _cancelId
	r.Set("cancel_id", _cancelId)
	return nil
}

// GetCancelId CancelId Getter
func (r TaobaowlbimportsgeneralconsignAPIRequest) GetCancelId() int64 {
	return r._cancelId
}
