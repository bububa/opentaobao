package opentrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单创建 API请求
taobao.opentrade.create.order

交易开放创建订单
*/
type TaobaoOpentradeCreateOrderRequest struct {
    model.Params
    // 外部订单ID
    outId   string
    // 买家openID
    openUserId   string
    // 收货地址的收件人姓名
    fullName   string
    // 收货地址的手机号码
    mobile   string
    // 收货地址
    address   string
    // 卖家备忘
    sellerMemo   string
    // 卖家备忘
    buyerMemo   string
    // 商品信息，一次不能超过10个
    itemInfos   []ItemInfo
}

// 初始化TaobaoOpentradeCreateOrderRequest对象
func NewTaobaoOpentradeCreateOrderRequest() *TaobaoOpentradeCreateOrderRequest{
    return &TaobaoOpentradeCreateOrderRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpentradeCreateOrderRequest) GetApiMethodName() string {
    return "taobao.opentrade.create.order"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpentradeCreateOrderRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OutId Setter
// 外部订单ID
func (r *TaobaoOpentradeCreateOrderRequest) SetOutId(outId string) error {
    r.outId = outId
    r.Set("out_id", outId)
    return nil
}

// OutId Getter
func (r TaobaoOpentradeCreateOrderRequest) GetOutId() string {
    return r.outId
}
// OpenUserId Setter
// 买家openID
func (r *TaobaoOpentradeCreateOrderRequest) SetOpenUserId(openUserId string) error {
    r.openUserId = openUserId
    r.Set("open_user_id", openUserId)
    return nil
}

// OpenUserId Getter
func (r TaobaoOpentradeCreateOrderRequest) GetOpenUserId() string {
    return r.openUserId
}
// FullName Setter
// 收货地址的收件人姓名
func (r *TaobaoOpentradeCreateOrderRequest) SetFullName(fullName string) error {
    r.fullName = fullName
    r.Set("full_name", fullName)
    return nil
}

// FullName Getter
func (r TaobaoOpentradeCreateOrderRequest) GetFullName() string {
    return r.fullName
}
// Mobile Setter
// 收货地址的手机号码
func (r *TaobaoOpentradeCreateOrderRequest) SetMobile(mobile string) error {
    r.mobile = mobile
    r.Set("mobile", mobile)
    return nil
}

// Mobile Getter
func (r TaobaoOpentradeCreateOrderRequest) GetMobile() string {
    return r.mobile
}
// Address Setter
// 收货地址
func (r *TaobaoOpentradeCreateOrderRequest) SetAddress(address string) error {
    r.address = address
    r.Set("address", address)
    return nil
}

// Address Getter
func (r TaobaoOpentradeCreateOrderRequest) GetAddress() string {
    return r.address
}
// SellerMemo Setter
// 卖家备忘
func (r *TaobaoOpentradeCreateOrderRequest) SetSellerMemo(sellerMemo string) error {
    r.sellerMemo = sellerMemo
    r.Set("seller_memo", sellerMemo)
    return nil
}

// SellerMemo Getter
func (r TaobaoOpentradeCreateOrderRequest) GetSellerMemo() string {
    return r.sellerMemo
}
// BuyerMemo Setter
// 卖家备忘
func (r *TaobaoOpentradeCreateOrderRequest) SetBuyerMemo(buyerMemo string) error {
    r.buyerMemo = buyerMemo
    r.Set("buyer_memo", buyerMemo)
    return nil
}

// BuyerMemo Getter
func (r TaobaoOpentradeCreateOrderRequest) GetBuyerMemo() string {
    return r.buyerMemo
}
// ItemInfos Setter
// 商品信息，一次不能超过10个
func (r *TaobaoOpentradeCreateOrderRequest) SetItemInfos(itemInfos []ItemInfo) error {
    r.itemInfos = itemInfos
    r.Set("item_infos", itemInfos)
    return nil
}

// ItemInfos Getter
func (r TaobaoOpentradeCreateOrderRequest) GetItemInfos() []ItemInfo {
    return r.itemInfos
}
