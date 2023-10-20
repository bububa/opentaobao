package fenxiao

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoInventoryAdjustTradeAPIRequest) Reset() {
	r._tbOrderType = ""
	r._operateTime = ""
	r._bizUniqueCode = ""
	r._items = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoInventoryAdjustTradeAPIRequest) GetApiMethodName() string {
	return "taobao.inventory.adjust.trade"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoInventoryAdjustTradeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoInventoryAdjustTradeAPIRequest) GetRawParams() model.Params {
	return r.Params
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
// 商品初始库存信息： [{ &#34;TBOrderCode”:”淘宝交易号”,&#34;TBSubOrderCode &#34;:&#34;淘宝子交易单号,赠品可以不填&#34;,&#34;”isGift”:”TRUE或者FALSE,是否赠品”,storeCode&#34;:&#34;商家仓库编码&#34;,&#34; scItemId &#34;:&#34;商品后端ID&#34;,&#34;scItemCode&#34;:&#34;商品商家编码&#34;,&#34; originScItemId &#34;:&#34;原商品ID&#34;,&#34;inventoryType&#34;:&#34;&#34;,&#34;quantity&#34;:&#34;111&#34;,&#34;isComplete&#34;:&#34;TRUE或者FALSE，是否全部确认出库&#34;}]
func (r *TaobaoInventoryAdjustTradeAPIRequest) SetItems(_items string) error {
	r._items = _items
	r.Set("items", _items)
	return nil
}

// GetItems Items Getter
func (r TaobaoInventoryAdjustTradeAPIRequest) GetItems() string {
	return r._items
}

var poolTaobaoInventoryAdjustTradeAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoInventoryAdjustTradeRequest()
	},
}

// GetTaobaoInventoryAdjustTradeRequest 从 sync.Pool 获取 TaobaoInventoryAdjustTradeAPIRequest
func GetTaobaoInventoryAdjustTradeAPIRequest() *TaobaoInventoryAdjustTradeAPIRequest {
	return poolTaobaoInventoryAdjustTradeAPIRequest.Get().(*TaobaoInventoryAdjustTradeAPIRequest)
}

// ReleaseTaobaoInventoryAdjustTradeAPIRequest 将 TaobaoInventoryAdjustTradeAPIRequest 放入 sync.Pool
func ReleaseTaobaoInventoryAdjustTradeAPIRequest(v *TaobaoInventoryAdjustTradeAPIRequest) {
	v.Reset()
	poolTaobaoInventoryAdjustTradeAPIRequest.Put(v)
}
