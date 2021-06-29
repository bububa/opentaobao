package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
阿里健康获取某一药店全部订单 API请求
taobao.trade.drug.orders.get

阿里健康获取某一药店全部订单
*/
type TaobaoTradeDrugOrdersGetRequest struct {
    model.Params
    // 外卖分店ID
    shopId   int64
    // 关键字
    keyword   string
    // true-查询仅按商家维度  false-查询按商家下所属店铺维度
    isAllShop   bool
    // true 仅有支付宝订单,false 包括所有类型订单(货到付款,支付券等)
    isAllOrder   bool
    // （必填字段）订单状态 待确认订单2 , 退款中订单4 , 已发货12 关闭20 交易成功21
    orderStatus   int64
    // 返回记录数，超过20按20条返回数据
    pageSize   int64
    // 页码
    pageNo   int64
}

// 初始化TaobaoTradeDrugOrdersGetRequest对象
func NewTaobaoTradeDrugOrdersGetRequest() *TaobaoTradeDrugOrdersGetRequest{
    return &TaobaoTradeDrugOrdersGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTradeDrugOrdersGetRequest) GetApiMethodName() string {
    return "taobao.trade.drug.orders.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTradeDrugOrdersGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ShopId Setter
// 外卖分店ID
func (r *TaobaoTradeDrugOrdersGetRequest) SetShopId(shopId int64) error {
    r.shopId = shopId
    r.Set("shop_id", shopId)
    return nil
}

// ShopId Getter
func (r TaobaoTradeDrugOrdersGetRequest) GetShopId() int64 {
    return r.shopId
}
// Keyword Setter
// 关键字
func (r *TaobaoTradeDrugOrdersGetRequest) SetKeyword(keyword string) error {
    r.keyword = keyword
    r.Set("keyword", keyword)
    return nil
}

// Keyword Getter
func (r TaobaoTradeDrugOrdersGetRequest) GetKeyword() string {
    return r.keyword
}
// IsAllShop Setter
// true-查询仅按商家维度  false-查询按商家下所属店铺维度
func (r *TaobaoTradeDrugOrdersGetRequest) SetIsAllShop(isAllShop bool) error {
    r.isAllShop = isAllShop
    r.Set("is_all_shop", isAllShop)
    return nil
}

// IsAllShop Getter
func (r TaobaoTradeDrugOrdersGetRequest) GetIsAllShop() bool {
    return r.isAllShop
}
// IsAllOrder Setter
// true 仅有支付宝订单,false 包括所有类型订单(货到付款,支付券等)
func (r *TaobaoTradeDrugOrdersGetRequest) SetIsAllOrder(isAllOrder bool) error {
    r.isAllOrder = isAllOrder
    r.Set("is_all_order", isAllOrder)
    return nil
}

// IsAllOrder Getter
func (r TaobaoTradeDrugOrdersGetRequest) GetIsAllOrder() bool {
    return r.isAllOrder
}
// OrderStatus Setter
// （必填字段）订单状态 待确认订单2 , 退款中订单4 , 已发货12 关闭20 交易成功21
func (r *TaobaoTradeDrugOrdersGetRequest) SetOrderStatus(orderStatus int64) error {
    r.orderStatus = orderStatus
    r.Set("order_status", orderStatus)
    return nil
}

// OrderStatus Getter
func (r TaobaoTradeDrugOrdersGetRequest) GetOrderStatus() int64 {
    return r.orderStatus
}
// PageSize Setter
// 返回记录数，超过20按20条返回数据
func (r *TaobaoTradeDrugOrdersGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r TaobaoTradeDrugOrdersGetRequest) GetPageSize() int64 {
    return r.pageSize
}
// PageNo Setter
// 页码
func (r *TaobaoTradeDrugOrdersGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r TaobaoTradeDrugOrdersGetRequest) GetPageNo() int64 {
    return r.pageNo
}
