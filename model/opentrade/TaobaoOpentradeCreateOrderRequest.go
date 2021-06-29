package opentrade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单创建 APIRequest
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

func NewTaobaoOpentradeCreateOrderRequest() *TaobaoOpentradeCreateOrderRequest{
    return &TaobaoOpentradeCreateOrderRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpentradeCreateOrderRequest) GetApiMethodName() string {
    return "taobao.opentrade.create.order"
}

func (r TaobaoOpentradeCreateOrderRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpentradeCreateOrderRequest) SetOutId(outId string) error {
    r.outId = outId
    r.Set("out_id", outId)
    return nil
}

func (r TaobaoOpentradeCreateOrderRequest) GetOutId() string {
    return r.outId
}

func (r *TaobaoOpentradeCreateOrderRequest) SetOpenUserId(openUserId string) error {
    r.openUserId = openUserId
    r.Set("open_user_id", openUserId)
    return nil
}

func (r TaobaoOpentradeCreateOrderRequest) GetOpenUserId() string {
    return r.openUserId
}

func (r *TaobaoOpentradeCreateOrderRequest) SetFullName(fullName string) error {
    r.fullName = fullName
    r.Set("full_name", fullName)
    return nil
}

func (r TaobaoOpentradeCreateOrderRequest) GetFullName() string {
    return r.fullName
}

func (r *TaobaoOpentradeCreateOrderRequest) SetMobile(mobile string) error {
    r.mobile = mobile
    r.Set("mobile", mobile)
    return nil
}

func (r TaobaoOpentradeCreateOrderRequest) GetMobile() string {
    return r.mobile
}

func (r *TaobaoOpentradeCreateOrderRequest) SetAddress(address string) error {
    r.address = address
    r.Set("address", address)
    return nil
}

func (r TaobaoOpentradeCreateOrderRequest) GetAddress() string {
    return r.address
}

func (r *TaobaoOpentradeCreateOrderRequest) SetSellerMemo(sellerMemo string) error {
    r.sellerMemo = sellerMemo
    r.Set("seller_memo", sellerMemo)
    return nil
}

func (r TaobaoOpentradeCreateOrderRequest) GetSellerMemo() string {
    return r.sellerMemo
}

func (r *TaobaoOpentradeCreateOrderRequest) SetBuyerMemo(buyerMemo string) error {
    r.buyerMemo = buyerMemo
    r.Set("buyer_memo", buyerMemo)
    return nil
}

func (r TaobaoOpentradeCreateOrderRequest) GetBuyerMemo() string {
    return r.buyerMemo
}

func (r *TaobaoOpentradeCreateOrderRequest) SetItemInfos(itemInfos []ItemInfo) error {
    r.itemInfos = itemInfos
    r.Set("item_infos", itemInfos)
    return nil
}

func (r TaobaoOpentradeCreateOrderRequest) GetItemInfos() []ItemInfo {
    return r.itemInfos
}

