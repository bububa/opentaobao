package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoInventoryAdjustTradeAPIRequest 交易库存调整单 API请求
// taobao.inventory.adjust.trade
//
// 商家交易调整库存，淘宝交易、B2B经销等
type TaobaoInventoryAdjustTradeAPIRequest struct {
	model.Params
	// 订单类型：B2C、B2B
	_tbOrderType string
	// 业务操作时间
	_operateTime string
	// 商家外部定单号
	_bizUniqueCode string
	// 商品初始库存信息： [{ "TBOrderCode”:”淘宝交易号”,"TBSubOrderCode ":"淘宝子交易单号,赠品可以不填","”isGift”:”TRUE或者FALSE,是否赠品”,storeCode":"商家仓库编码"," scItemId ":"商品后端ID","scItemCode":"商品商家编码"," originScItemId ":"原商品ID","inventoryType":"","quantity":"111","isComplete":"TRUE或者FALSE，是否全部确认出库"}]
	_items string
}

// NewTaobaoInventoryAdjustTradeRequest 初始化TaobaoInventoryAdjustTradeAPIRequest对象
func NewTaobaoInventoryAdjustTradeRequest() *TaobaoInventoryAdjustTradeAPIRequest {
	return &TaobaoInventoryAdjustTradeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoInventoryAdjustTradeAPIRequest) GetApiMethodName() string {
	return "taobao.inventory.adjust.trade"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoInventoryAdjustTradeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTbOrderType is TbOrderType Setter
// 订单类型：B2C、B2B
func (r *TaobaoInventoryAdjustTradeAPIRequest) SetTbOrderType(_tbOrderType string) error {
	r._tbOrderType = _tbOrderType
	r.Set("tb_order_type", _tbOrderType)
	return nil
}

// GetTbOrderType TbOrderType Getter
func (r TaobaoInventoryAdjustTradeAPIRequest) GetTbOrderType() string {
	return r._tbOrderType
}

// SetOperateTime is OperateTime Setter
// 业务操作时间
func (r *TaobaoInventoryAdjustTradeAPIRequest) SetOperateTime(_operateTime string) error {
	r._operateTime = _operateTime
	r.Set("operate_time", _operateTime)
	return nil
}

// GetOperateTime OperateTime Getter
func (r TaobaoInventoryAdjustTradeAPIRequest) GetOperateTime() string {
	return r._operateTime
}

// SetBizUniqueCode is BizUniqueCode Setter
// 商家外部定单号
func (r *TaobaoInventoryAdjustTradeAPIRequest) SetBizUniqueCode(_bizUniqueCode string) error {
	r._bizUniqueCode = _bizUniqueCode
	r.Set("biz_unique_code", _bizUniqueCode)
	return nil
}

// GetBizUniqueCode BizUniqueCode Getter
func (r TaobaoInventoryAdjustTradeAPIRequest) GetBizUniqueCode() string {
	return r._bizUniqueCode
}

// SetItems is Items Setter
// 商品初始库存信息： [{ "TBOrderCode”:”淘宝交易号”,"TBSubOrderCode ":"淘宝子交易单号,赠品可以不填","”isGift”:”TRUE或者FALSE,是否赠品”,storeCode":"商家仓库编码"," scItemId ":"商品后端ID","scItemCode":"商品商家编码"," originScItemId ":"原商品ID","inventoryType":"","quantity":"111","isComplete":"TRUE或者FALSE，是否全部确认出库"}]
func (r *TaobaoInventoryAdjustTradeAPIRequest) SetItems(_items string) error {
	r._items = _items
	r.Set("items", _items)
	return nil
}

// GetItems Items Getter
func (r TaobaoInventoryAdjustTradeAPIRequest) GetItems() string {
	return r._items
}
