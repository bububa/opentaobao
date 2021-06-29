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
type TaobaoRpReturngoodsAgreeRequest struct {
    model.Params
    // 退款编号
    refundId   int64
    // 卖家姓名，淘宝退款为必填项。
    name   string
    // 卖家提供的退货地址，淘宝退款为必填项。
    address   string
    // 卖家提供的退货地址的邮编，淘宝退款为必填项。
    post   string
    // 卖家座机，淘宝退款为必填项。
    tel   string
    // 卖家手机，淘宝退款为必填项。
    mobile   string
    // 卖家退货留言，天猫退款为必填项。
    remark   string
    // 售中：onsale，售后：aftersale，天猫退款为必填项。
    refundPhase   string
    // 退款版本号，天猫退款为必填项。
    refundVersion   int64
    // 卖家收货地址编号，天猫淘宝退款都为必填项。
    sellerAddressId   int64
    // 邮费承担方，买家承担值为1，卖家承担值为0
    postFeeBearRole   int64
    // 是否虚拟退货，可选项
    virtualReturnGoods   bool
}

// 初始化TaobaoRpReturngoodsAgreeRequest对象
func NewTaobaoRpReturngoodsAgreeRequest() *TaobaoRpReturngoodsAgreeRequest{
    return &TaobaoRpReturngoodsAgreeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoRpReturngoodsAgreeRequest) GetApiMethodName() string {
    return "taobao.rp.returngoods.agree"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoRpReturngoodsAgreeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefundId Setter
// 退款编号
func (r *TaobaoRpReturngoodsAgreeRequest) SetRefundId(refundId int64) error {
    r.refundId = refundId
    r.Set("refund_id", refundId)
    return nil
}

// RefundId Getter
func (r TaobaoRpReturngoodsAgreeRequest) GetRefundId() int64 {
    return r.refundId
}
// Name Setter
// 卖家姓名，淘宝退款为必填项。
func (r *TaobaoRpReturngoodsAgreeRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

// Name Getter
func (r TaobaoRpReturngoodsAgreeRequest) GetName() string {
    return r.name
}
// Address Setter
// 卖家提供的退货地址，淘宝退款为必填项。
func (r *TaobaoRpReturngoodsAgreeRequest) SetAddress(address string) error {
    r.address = address
    r.Set("address", address)
    return nil
}

// Address Getter
func (r TaobaoRpReturngoodsAgreeRequest) GetAddress() string {
    return r.address
}
// Post Setter
// 卖家提供的退货地址的邮编，淘宝退款为必填项。
func (r *TaobaoRpReturngoodsAgreeRequest) SetPost(post string) error {
    r.post = post
    r.Set("post", post)
    return nil
}

// Post Getter
func (r TaobaoRpReturngoodsAgreeRequest) GetPost() string {
    return r.post
}
// Tel Setter
// 卖家座机，淘宝退款为必填项。
func (r *TaobaoRpReturngoodsAgreeRequest) SetTel(tel string) error {
    r.tel = tel
    r.Set("tel", tel)
    return nil
}

// Tel Getter
func (r TaobaoRpReturngoodsAgreeRequest) GetTel() string {
    return r.tel
}
// Mobile Setter
// 卖家手机，淘宝退款为必填项。
func (r *TaobaoRpReturngoodsAgreeRequest) SetMobile(mobile string) error {
    r.mobile = mobile
    r.Set("mobile", mobile)
    return nil
}

// Mobile Getter
func (r TaobaoRpReturngoodsAgreeRequest) GetMobile() string {
    return r.mobile
}
// Remark Setter
// 卖家退货留言，天猫退款为必填项。
func (r *TaobaoRpReturngoodsAgreeRequest) SetRemark(remark string) error {
    r.remark = remark
    r.Set("remark", remark)
    return nil
}

// Remark Getter
func (r TaobaoRpReturngoodsAgreeRequest) GetRemark() string {
    return r.remark
}
// RefundPhase Setter
// 售中：onsale，售后：aftersale，天猫退款为必填项。
func (r *TaobaoRpReturngoodsAgreeRequest) SetRefundPhase(refundPhase string) error {
    r.refundPhase = refundPhase
    r.Set("refund_phase", refundPhase)
    return nil
}

// RefundPhase Getter
func (r TaobaoRpReturngoodsAgreeRequest) GetRefundPhase() string {
    return r.refundPhase
}
// RefundVersion Setter
// 退款版本号，天猫退款为必填项。
func (r *TaobaoRpReturngoodsAgreeRequest) SetRefundVersion(refundVersion int64) error {
    r.refundVersion = refundVersion
    r.Set("refund_version", refundVersion)
    return nil
}

// RefundVersion Getter
func (r TaobaoRpReturngoodsAgreeRequest) GetRefundVersion() int64 {
    return r.refundVersion
}
// SellerAddressId Setter
// 卖家收货地址编号，天猫淘宝退款都为必填项。
func (r *TaobaoRpReturngoodsAgreeRequest) SetSellerAddressId(sellerAddressId int64) error {
    r.sellerAddressId = sellerAddressId
    r.Set("seller_address_id", sellerAddressId)
    return nil
}

// SellerAddressId Getter
func (r TaobaoRpReturngoodsAgreeRequest) GetSellerAddressId() int64 {
    return r.sellerAddressId
}
// PostFeeBearRole Setter
// 邮费承担方，买家承担值为1，卖家承担值为0
func (r *TaobaoRpReturngoodsAgreeRequest) SetPostFeeBearRole(postFeeBearRole int64) error {
    r.postFeeBearRole = postFeeBearRole
    r.Set("post_fee_bear_role", postFeeBearRole)
    return nil
}

// PostFeeBearRole Getter
func (r TaobaoRpReturngoodsAgreeRequest) GetPostFeeBearRole() int64 {
    return r.postFeeBearRole
}
// VirtualReturnGoods Setter
// 是否虚拟退货，可选项
func (r *TaobaoRpReturngoodsAgreeRequest) SetVirtualReturnGoods(virtualReturnGoods bool) error {
    r.virtualReturnGoods = virtualReturnGoods
    r.Set("virtual_return_goods", virtualReturnGoods)
    return nil
}

// VirtualReturnGoods Getter
func (r TaobaoRpReturngoodsAgreeRequest) GetVirtualReturnGoods() bool {
    return r.virtualReturnGoods
}
