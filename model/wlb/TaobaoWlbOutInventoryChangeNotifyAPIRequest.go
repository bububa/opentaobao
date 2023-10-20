package wlb

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoWlbOutInventoryChangeNotifyAPIRequest 外部库存变化通知（企业物流用户使用） API请求
// taobao.wlb.out.inventory.change.notify
//
// 拥有自有仓的企业物流用户通过该接口把自有仓的库存通知到物流宝，由物流宝维护该库存，控制前台显示库存的准确性。
type TaobaoWlbOutInventoryChangeNotifyAPIRequest struct {
	model.Params
	// WLB_ITEM--物流宝商品 IC_ITEM--淘宝商品 IC_SKU--淘宝sku
	_type string
	// OUT--出库 IN--入库
	_opType string
	// （1）OTHER： 其他 （2）TAOBAO_TRADE： 淘宝交易 （3）OTHER_TRADE：其他交易 （4）ALLCOATE： 调拨 （5）CHECK:盘点 （6）PURCHASE:采购
	_source string
	// 订单号，如果source为TAOBAO_TRADE,order_source_code必须为淘宝交易号
	_orderSourceCode string
	// 库存变化唯一标识，用于去重，一个外部唯一编码唯一标识一次库存变化。
	_outBizCode string
	// 目前非必须，系统会选择默认值
	_storeCode string
	// 物流宝商品id或前台宝贝id（由type类型决定）
	_itemId int64
	// 库存变化数量
	_changeCount int64
	// 本次库存变化后库存余额
	_resultCount int64
}

// NewTaobaoWlbOutInventoryChangeNotifyRequest 初始化TaobaoWlbOutInventoryChangeNotifyAPIRequest对象
func NewTaobaoWlbOutInventoryChangeNotifyRequest() *TaobaoWlbOutInventoryChangeNotifyAPIRequest {
	return &TaobaoWlbOutInventoryChangeNotifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoWlbOutInventoryChangeNotifyAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.out.inventory.change.notify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoWlbOutInventoryChangeNotifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoWlbOutInventoryChangeNotifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetType is Type Setter
// WLB_ITEM--物流宝商品 IC_ITEM--淘宝商品 IC_SKU--淘宝sku
func (r *TaobaoWlbOutInventoryChangeNotifyAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TaobaoWlbOutInventoryChangeNotifyAPIRequest) GetType() string {
	return r._type
}

// SetOpType is OpType Setter
// OUT--出库 IN--入库
func (r *TaobaoWlbOutInventoryChangeNotifyAPIRequest) SetOpType(_opType string) error {
	r._opType = _opType
	r.Set("op_type", _opType)
	return nil
}

// GetOpType OpType Getter
func (r TaobaoWlbOutInventoryChangeNotifyAPIRequest) GetOpType() string {
	return r._opType
}

// SetSource is Source Setter
// （1）OTHER： 其他 （2）TAOBAO_TRADE： 淘宝交易 （3）OTHER_TRADE：其他交易 （4）ALLCOATE： 调拨 （5）CHECK:盘点 （6）PURCHASE:采购
func (r *TaobaoWlbOutInventoryChangeNotifyAPIRequest) SetSource(_source string) error {
	r._source = _source
	r.Set("source", _source)
	return nil
}

// GetSource Source Getter
func (r TaobaoWlbOutInventoryChangeNotifyAPIRequest) GetSource() string {
	return r._source
}

// SetOrderSourceCode is OrderSourceCode Setter
// 订单号，如果source为TAOBAO_TRADE,order_source_code必须为淘宝交易号
func (r *TaobaoWlbOutInventoryChangeNotifyAPIRequest) SetOrderSourceCode(_orderSourceCode string) error {
	r._orderSourceCode = _orderSourceCode
	r.Set("order_source_code", _orderSourceCode)
	return nil
}

// GetOrderSourceCode OrderSourceCode Getter
func (r TaobaoWlbOutInventoryChangeNotifyAPIRequest) GetOrderSourceCode() string {
	return r._orderSourceCode
}

// SetOutBizCode is OutBizCode Setter
// 库存变化唯一标识，用于去重，一个外部唯一编码唯一标识一次库存变化。
func (r *TaobaoWlbOutInventoryChangeNotifyAPIRequest) SetOutBizCode(_outBizCode string) error {
	r._outBizCode = _outBizCode
	r.Set("out_biz_code", _outBizCode)
	return nil
}

// GetOutBizCode OutBizCode Getter
func (r TaobaoWlbOutInventoryChangeNotifyAPIRequest) GetOutBizCode() string {
	return r._outBizCode
}

// SetStoreCode is StoreCode Setter
// 目前非必须，系统会选择默认值
func (r *TaobaoWlbOutInventoryChangeNotifyAPIRequest) SetStoreCode(_storeCode string) error {
	r._storeCode = _storeCode
	r.Set("store_code", _storeCode)
	return nil
}

// GetStoreCode StoreCode Getter
func (r TaobaoWlbOutInventoryChangeNotifyAPIRequest) GetStoreCode() string {
	return r._storeCode
}

// SetItemId is ItemId Setter
// 物流宝商品id或前台宝贝id（由type类型决定）
func (r *TaobaoWlbOutInventoryChangeNotifyAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r TaobaoWlbOutInventoryChangeNotifyAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetChangeCount is ChangeCount Setter
// 库存变化数量
func (r *TaobaoWlbOutInventoryChangeNotifyAPIRequest) SetChangeCount(_changeCount int64) error {
	r._changeCount = _changeCount
	r.Set("change_count", _changeCount)
	return nil
}

// GetChangeCount ChangeCount Getter
func (r TaobaoWlbOutInventoryChangeNotifyAPIRequest) GetChangeCount() int64 {
	return r._changeCount
}

// SetResultCount is ResultCount Setter
// 本次库存变化后库存余额
func (r *TaobaoWlbOutInventoryChangeNotifyAPIRequest) SetResultCount(_resultCount int64) error {
	r._resultCount = _resultCount
	r.Set("result_count", _resultCount)
	return nil
}

// GetResultCount ResultCount Getter
func (r TaobaoWlbOutInventoryChangeNotifyAPIRequest) GetResultCount() int64 {
	return r._resultCount
}
