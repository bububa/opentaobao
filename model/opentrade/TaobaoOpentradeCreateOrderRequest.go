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
type TaobaoOpentradeCreateOrderAPIRequest struct {
    model.Params
    // 外部订单ID
    _outId   string
    // 买家openID
    _openUserId   string
    // 收货地址的收件人姓名
    _fullName   string
    // 收货地址的手机号码
    _mobile   string
    // 收货地址
    _address   string
    // 卖家备忘
    _sellerMemo   string
    // 卖家备忘
    _buyerMemo   string
    // 商品信息，一次不能超过10个
    _itemInfos   []ItemInfo
}

// 初始化TaobaoOpentradeCreateOrderAPIRequest对象
func NewTaobaoOpentradeCreateOrderRequest() *TaobaoOpentradeCreateOrderAPIRequest{
    return &TaobaoOpentradeCreateOrderAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpentradeCreateOrderAPIRequest) GetApiMethodName() string {
    return "taobao.opentrade.create.order"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpentradeCreateOrderAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OutId Setter
// 外部订单ID
func (r *TaobaoOpentradeCreateOrderAPIRequest) SetOutId(_outId string) error {
    r._outId = _outId
    r.Set("out_id", _outId)
    return nil
}

// OutId Getter
func (r TaobaoOpentradeCreateOrderAPIRequest) GetOutId() string {
    return r._outId
}
// OpenUserId Setter
// 买家openID
func (r *TaobaoOpentradeCreateOrderAPIRequest) SetOpenUserId(_openUserId string) error {
    r._openUserId = _openUserId
    r.Set("open_user_id", _openUserId)
    return nil
}

// OpenUserId Getter
func (r TaobaoOpentradeCreateOrderAPIRequest) GetOpenUserId() string {
    return r._openUserId
}
// FullName Setter
// 收货地址的收件人姓名
func (r *TaobaoOpentradeCreateOrderAPIRequest) SetFullName(_fullName string) error {
    r._fullName = _fullName
    r.Set("full_name", _fullName)
    return nil
}

// FullName Getter
func (r TaobaoOpentradeCreateOrderAPIRequest) GetFullName() string {
    return r._fullName
}
// Mobile Setter
// 收货地址的手机号码
func (r *TaobaoOpentradeCreateOrderAPIRequest) SetMobile(_mobile string) error {
    r._mobile = _mobile
    r.Set("mobile", _mobile)
    return nil
}

// Mobile Getter
func (r TaobaoOpentradeCreateOrderAPIRequest) GetMobile() string {
    return r._mobile
}
// Address Setter
// 收货地址
func (r *TaobaoOpentradeCreateOrderAPIRequest) SetAddress(_address string) error {
    r._address = _address
    r.Set("address", _address)
    return nil
}

// Address Getter
func (r TaobaoOpentradeCreateOrderAPIRequest) GetAddress() string {
    return r._address
}
// SellerMemo Setter
// 卖家备忘
func (r *TaobaoOpentradeCreateOrderAPIRequest) SetSellerMemo(_sellerMemo string) error {
    r._sellerMemo = _sellerMemo
    r.Set("seller_memo", _sellerMemo)
    return nil
}

// SellerMemo Getter
func (r TaobaoOpentradeCreateOrderAPIRequest) GetSellerMemo() string {
    return r._sellerMemo
}
// BuyerMemo Setter
// 卖家备忘
func (r *TaobaoOpentradeCreateOrderAPIRequest) SetBuyerMemo(_buyerMemo string) error {
    r._buyerMemo = _buyerMemo
    r.Set("buyer_memo", _buyerMemo)
    return nil
}

// BuyerMemo Getter
func (r TaobaoOpentradeCreateOrderAPIRequest) GetBuyerMemo() string {
    return r._buyerMemo
}
// ItemInfos Setter
// 商品信息，一次不能超过10个
func (r *TaobaoOpentradeCreateOrderAPIRequest) SetItemInfos(_itemInfos []ItemInfo) error {
    r._itemInfos = _itemInfos
    r.Set("item_infos", _itemInfos)
    return nil
}

// ItemInfos Getter
func (r TaobaoOpentradeCreateOrderAPIRequest) GetItemInfos() []ItemInfo {
    return r._itemInfos
}
