package refund

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
卖家同意退货 API请求
taobao.rp.returngoods.agree

卖家同意退货，支持淘宝和天猫的订单。
*/
type TaobaoRpReturngoodsAgreeAPIRequest struct {
    model.Params
    // 退款编号
    _refundId   int64
    // 卖家姓名，淘宝退款为必填项。
    _name   string
    // 卖家提供的退货地址，淘宝退款为必填项。
    _address   string
    // 卖家提供的退货地址的邮编，淘宝退款为必填项。
    _post   string
    // 卖家座机，淘宝退款为必填项。
    _tel   string
    // 卖家手机，淘宝退款为必填项。
    _mobile   string
    // 卖家退货留言，天猫退款为必填项。
    _remark   string
    // 售中：onsale，售后：aftersale，天猫退款为必填项。
    _refundPhase   string
    // 退款版本号，天猫退款为必填项。
    _refundVersion   int64
    // 卖家收货地址编号，天猫淘宝退款都为必填项。
    _sellerAddressId   int64
    // 邮费承担方，买家承担值为1，卖家承担值为0
    _postFeeBearRole   int64
    // 是否虚拟退货，可选项
    _virtualReturnGoods   bool
}

// 初始化TaobaoRpReturngoodsAgreeAPIRequest对象
func NewTaobaoRpReturngoodsAgreeRequest() *TaobaoRpReturngoodsAgreeAPIRequest{
    return &TaobaoRpReturngoodsAgreeAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoRpReturngoodsAgreeAPIRequest) GetApiMethodName() string {
    return "taobao.rp.returngoods.agree"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoRpReturngoodsAgreeAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefundId Setter
// 退款编号
func (r *TaobaoRpReturngoodsAgreeAPIRequest) SetRefundId(_refundId int64) error {
    r._refundId = _refundId
    r.Set("refund_id", _refundId)
    return nil
}

// RefundId Getter
func (r TaobaoRpReturngoodsAgreeAPIRequest) GetRefundId() int64 {
    return r._refundId
}
// Name Setter
// 卖家姓名，淘宝退款为必填项。
func (r *TaobaoRpReturngoodsAgreeAPIRequest) SetName(_name string) error {
    r._name = _name
    r.Set("name", _name)
    return nil
}

// Name Getter
func (r TaobaoRpReturngoodsAgreeAPIRequest) GetName() string {
    return r._name
}
// Address Setter
// 卖家提供的退货地址，淘宝退款为必填项。
func (r *TaobaoRpReturngoodsAgreeAPIRequest) SetAddress(_address string) error {
    r._address = _address
    r.Set("address", _address)
    return nil
}

// Address Getter
func (r TaobaoRpReturngoodsAgreeAPIRequest) GetAddress() string {
    return r._address
}
// Post Setter
// 卖家提供的退货地址的邮编，淘宝退款为必填项。
func (r *TaobaoRpReturngoodsAgreeAPIRequest) SetPost(_post string) error {
    r._post = _post
    r.Set("post", _post)
    return nil
}

// Post Getter
func (r TaobaoRpReturngoodsAgreeAPIRequest) GetPost() string {
    return r._post
}
// Tel Setter
// 卖家座机，淘宝退款为必填项。
func (r *TaobaoRpReturngoodsAgreeAPIRequest) SetTel(_tel string) error {
    r._tel = _tel
    r.Set("tel", _tel)
    return nil
}

// Tel Getter
func (r TaobaoRpReturngoodsAgreeAPIRequest) GetTel() string {
    return r._tel
}
// Mobile Setter
// 卖家手机，淘宝退款为必填项。
func (r *TaobaoRpReturngoodsAgreeAPIRequest) SetMobile(_mobile string) error {
    r._mobile = _mobile
    r.Set("mobile", _mobile)
    return nil
}

// Mobile Getter
func (r TaobaoRpReturngoodsAgreeAPIRequest) GetMobile() string {
    return r._mobile
}
// Remark Setter
// 卖家退货留言，天猫退款为必填项。
func (r *TaobaoRpReturngoodsAgreeAPIRequest) SetRemark(_remark string) error {
    r._remark = _remark
    r.Set("remark", _remark)
    return nil
}

// Remark Getter
func (r TaobaoRpReturngoodsAgreeAPIRequest) GetRemark() string {
    return r._remark
}
// RefundPhase Setter
// 售中：onsale，售后：aftersale，天猫退款为必填项。
func (r *TaobaoRpReturngoodsAgreeAPIRequest) SetRefundPhase(_refundPhase string) error {
    r._refundPhase = _refundPhase
    r.Set("refund_phase", _refundPhase)
    return nil
}

// RefundPhase Getter
func (r TaobaoRpReturngoodsAgreeAPIRequest) GetRefundPhase() string {
    return r._refundPhase
}
// RefundVersion Setter
// 退款版本号，天猫退款为必填项。
func (r *TaobaoRpReturngoodsAgreeAPIRequest) SetRefundVersion(_refundVersion int64) error {
    r._refundVersion = _refundVersion
    r.Set("refund_version", _refundVersion)
    return nil
}

// RefundVersion Getter
func (r TaobaoRpReturngoodsAgreeAPIRequest) GetRefundVersion() int64 {
    return r._refundVersion
}
// SellerAddressId Setter
// 卖家收货地址编号，天猫淘宝退款都为必填项。
func (r *TaobaoRpReturngoodsAgreeAPIRequest) SetSellerAddressId(_sellerAddressId int64) error {
    r._sellerAddressId = _sellerAddressId
    r.Set("seller_address_id", _sellerAddressId)
    return nil
}

// SellerAddressId Getter
func (r TaobaoRpReturngoodsAgreeAPIRequest) GetSellerAddressId() int64 {
    return r._sellerAddressId
}
// PostFeeBearRole Setter
// 邮费承担方，买家承担值为1，卖家承担值为0
func (r *TaobaoRpReturngoodsAgreeAPIRequest) SetPostFeeBearRole(_postFeeBearRole int64) error {
    r._postFeeBearRole = _postFeeBearRole
    r.Set("post_fee_bear_role", _postFeeBearRole)
    return nil
}

// PostFeeBearRole Getter
func (r TaobaoRpReturngoodsAgreeAPIRequest) GetPostFeeBearRole() int64 {
    return r._postFeeBearRole
}
// VirtualReturnGoods Setter
// 是否虚拟退货，可选项
func (r *TaobaoRpReturngoodsAgreeAPIRequest) SetVirtualReturnGoods(_virtualReturnGoods bool) error {
    r._virtualReturnGoods = _virtualReturnGoods
    r.Set("virtual_return_goods", _virtualReturnGoods)
    return nil
}

// VirtualReturnGoods Getter
func (r TaobaoRpReturngoodsAgreeAPIRequest) GetVirtualReturnGoods() bool {
    return r._virtualReturnGoods
}
