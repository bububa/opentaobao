package alihealth2

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTradeDrugOrdersGetAPIRequest 阿里健康获取某一药店全部订单 API请求
// taobao.trade.drug.orders.get
//
// 阿里健康获取某一药店全部订单
type TaobaoTradeDrugOrdersGetAPIRequest struct {
	model.Params
	// 关键字
	_keyword string
	// 外卖分店ID
	_shopId int64
	// （必填字段）订单状态 待确认订单2 , 退款中订单4 , 已发货12 关闭20 交易成功21
	_orderStatus int64
	// 返回记录数，超过20按20条返回数据
	_pageSize int64
	// 页码
	_pageNo int64
	// true-查询仅按商家维度  false-查询按商家下所属店铺维度
	_isAllShop bool
	// true 仅有支付宝订单,false 包括所有类型订单(货到付款,支付券等)
	_isAllOrder bool
}

// NewTaobaoTradeDrugOrdersGetRequest 初始化TaobaoTradeDrugOrdersGetAPIRequest对象
func NewTaobaoTradeDrugOrdersGetRequest() *TaobaoTradeDrugOrdersGetAPIRequest {
	return &TaobaoTradeDrugOrdersGetAPIRequest{
		Params: model.NewParams(7),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTradeDrugOrdersGetAPIRequest) Reset() {
	r._keyword = ""
	r._shopId = 0
	r._orderStatus = 0
	r._pageSize = 0
	r._pageNo = 0
	r._isAllShop = false
	r._isAllOrder = false
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTradeDrugOrdersGetAPIRequest) GetApiMethodName() string {
	return "taobao.trade.drug.orders.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTradeDrugOrdersGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTradeDrugOrdersGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetKeyword is Keyword Setter
// 关键字
func (r *TaobaoTradeDrugOrdersGetAPIRequest) SetKeyword(_keyword string) error {
	r._keyword = _keyword
	r.Set("keyword", _keyword)
	return nil
}

// GetKeyword Keyword Getter
func (r TaobaoTradeDrugOrdersGetAPIRequest) GetKeyword() string {
	return r._keyword
}

// SetShopId is ShopId Setter
// 外卖分店ID
func (r *TaobaoTradeDrugOrdersGetAPIRequest) SetShopId(_shopId int64) error {
	r._shopId = _shopId
	r.Set("shop_id", _shopId)
	return nil
}

// GetShopId ShopId Getter
func (r TaobaoTradeDrugOrdersGetAPIRequest) GetShopId() int64 {
	return r._shopId
}

// SetOrderStatus is OrderStatus Setter
// （必填字段）订单状态 待确认订单2 , 退款中订单4 , 已发货12 关闭20 交易成功21
func (r *TaobaoTradeDrugOrdersGetAPIRequest) SetOrderStatus(_orderStatus int64) error {
	r._orderStatus = _orderStatus
	r.Set("order_status", _orderStatus)
	return nil
}

// GetOrderStatus OrderStatus Getter
func (r TaobaoTradeDrugOrdersGetAPIRequest) GetOrderStatus() int64 {
	return r._orderStatus
}

// SetPageSize is PageSize Setter
// 返回记录数，超过20按20条返回数据
func (r *TaobaoTradeDrugOrdersGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoTradeDrugOrdersGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNo is PageNo Setter
// 页码
func (r *TaobaoTradeDrugOrdersGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoTradeDrugOrdersGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetIsAllShop is IsAllShop Setter
// true-查询仅按商家维度  false-查询按商家下所属店铺维度
func (r *TaobaoTradeDrugOrdersGetAPIRequest) SetIsAllShop(_isAllShop bool) error {
	r._isAllShop = _isAllShop
	r.Set("is_all_shop", _isAllShop)
	return nil
}

// GetIsAllShop IsAllShop Getter
func (r TaobaoTradeDrugOrdersGetAPIRequest) GetIsAllShop() bool {
	return r._isAllShop
}

// SetIsAllOrder is IsAllOrder Setter
// true 仅有支付宝订单,false 包括所有类型订单(货到付款,支付券等)
func (r *TaobaoTradeDrugOrdersGetAPIRequest) SetIsAllOrder(_isAllOrder bool) error {
	r._isAllOrder = _isAllOrder
	r.Set("is_all_order", _isAllOrder)
	return nil
}

// GetIsAllOrder IsAllOrder Getter
func (r TaobaoTradeDrugOrdersGetAPIRequest) GetIsAllOrder() bool {
	return r._isAllOrder
}

var poolTaobaoTradeDrugOrdersGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTradeDrugOrdersGetRequest()
	},
}

// GetTaobaoTradeDrugOrdersGetRequest 从 sync.Pool 获取 TaobaoTradeDrugOrdersGetAPIRequest
func GetTaobaoTradeDrugOrdersGetAPIRequest() *TaobaoTradeDrugOrdersGetAPIRequest {
	return poolTaobaoTradeDrugOrdersGetAPIRequest.Get().(*TaobaoTradeDrugOrdersGetAPIRequest)
}

// ReleaseTaobaoTradeDrugOrdersGetAPIRequest 将 TaobaoTradeDrugOrdersGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoTradeDrugOrdersGetAPIRequest(v *TaobaoTradeDrugOrdersGetAPIRequest) {
	v.Reset()
	poolTaobaoTradeDrugOrdersGetAPIRequest.Put(v)
}
