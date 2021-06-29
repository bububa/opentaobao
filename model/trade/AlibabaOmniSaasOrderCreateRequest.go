package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单创建接口 API请求
alibaba.omni.saas.order.create

服务商利用现有的saas系统和阿里完成交易系统的对接
*/
type AlibabaOmniSaasOrderCreateRequest struct {
    model.Params
    // 商品列表
    goodsDetails   []GoodsDetail
    // 买家标识，淘系用户或用户手机号。当支付渠道为支付宝时，此字段为淘宝会员码或支付宝付款码。(当前仅支持淘系用户，手机号下单稍后开放)
    buyerId   string
    // 门店ID
    storeId   string
    // 收银设备类型
    device   string
    // 收银设备号
    deviceNo   string
    // 收银员标识
    operatorId   string
    // ALIPAY：支付宝用户；TAOBAO：淘宝会员码；MOBILE：手机号
    buyerIdType   string
    // ALIPAY：支付宝付款；BANK_CARD：刷卡
    payChannel   string
    // 商家自有优惠
    couponInfos   []CouponInfo
    // PLACE：淘宝商户中心门店ID；CUSTOM：商户自有门店编码，需要维护到淘宝商户中心
    storeIdType   string
    // 请求号，用于标识一次请求
    requestNo   string
}

// 初始化AlibabaOmniSaasOrderCreateRequest对象
func NewAlibabaOmniSaasOrderCreateRequest() *AlibabaOmniSaasOrderCreateRequest{
    return &AlibabaOmniSaasOrderCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaOmniSaasOrderCreateRequest) GetApiMethodName() string {
    return "alibaba.omni.saas.order.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaOmniSaasOrderCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// GoodsDetails Setter
// 商品列表
func (r *AlibabaOmniSaasOrderCreateRequest) SetGoodsDetails(goodsDetails []GoodsDetail) error {
    r.goodsDetails = goodsDetails
    r.Set("goods_details", goodsDetails)
    return nil
}

// GoodsDetails Getter
func (r AlibabaOmniSaasOrderCreateRequest) GetGoodsDetails() []GoodsDetail {
    return r.goodsDetails
}
// BuyerId Setter
// 买家标识，淘系用户或用户手机号。当支付渠道为支付宝时，此字段为淘宝会员码或支付宝付款码。(当前仅支持淘系用户，手机号下单稍后开放)
func (r *AlibabaOmniSaasOrderCreateRequest) SetBuyerId(buyerId string) error {
    r.buyerId = buyerId
    r.Set("buyer_id", buyerId)
    return nil
}

// BuyerId Getter
func (r AlibabaOmniSaasOrderCreateRequest) GetBuyerId() string {
    return r.buyerId
}
// StoreId Setter
// 门店ID
func (r *AlibabaOmniSaasOrderCreateRequest) SetStoreId(storeId string) error {
    r.storeId = storeId
    r.Set("store_id", storeId)
    return nil
}

// StoreId Getter
func (r AlibabaOmniSaasOrderCreateRequest) GetStoreId() string {
    return r.storeId
}
// Device Setter
// 收银设备类型
func (r *AlibabaOmniSaasOrderCreateRequest) SetDevice(device string) error {
    r.device = device
    r.Set("device", device)
    return nil
}

// Device Getter
func (r AlibabaOmniSaasOrderCreateRequest) GetDevice() string {
    return r.device
}
// DeviceNo Setter
// 收银设备号
func (r *AlibabaOmniSaasOrderCreateRequest) SetDeviceNo(deviceNo string) error {
    r.deviceNo = deviceNo
    r.Set("device_no", deviceNo)
    return nil
}

// DeviceNo Getter
func (r AlibabaOmniSaasOrderCreateRequest) GetDeviceNo() string {
    return r.deviceNo
}
// OperatorId Setter
// 收银员标识
func (r *AlibabaOmniSaasOrderCreateRequest) SetOperatorId(operatorId string) error {
    r.operatorId = operatorId
    r.Set("operator_id", operatorId)
    return nil
}

// OperatorId Getter
func (r AlibabaOmniSaasOrderCreateRequest) GetOperatorId() string {
    return r.operatorId
}
// BuyerIdType Setter
// ALIPAY：支付宝用户；TAOBAO：淘宝会员码；MOBILE：手机号
func (r *AlibabaOmniSaasOrderCreateRequest) SetBuyerIdType(buyerIdType string) error {
    r.buyerIdType = buyerIdType
    r.Set("buyer_id_type", buyerIdType)
    return nil
}

// BuyerIdType Getter
func (r AlibabaOmniSaasOrderCreateRequest) GetBuyerIdType() string {
    return r.buyerIdType
}
// PayChannel Setter
// ALIPAY：支付宝付款；BANK_CARD：刷卡
func (r *AlibabaOmniSaasOrderCreateRequest) SetPayChannel(payChannel string) error {
    r.payChannel = payChannel
    r.Set("pay_channel", payChannel)
    return nil
}

// PayChannel Getter
func (r AlibabaOmniSaasOrderCreateRequest) GetPayChannel() string {
    return r.payChannel
}
// CouponInfos Setter
// 商家自有优惠
func (r *AlibabaOmniSaasOrderCreateRequest) SetCouponInfos(couponInfos []CouponInfo) error {
    r.couponInfos = couponInfos
    r.Set("coupon_infos", couponInfos)
    return nil
}

// CouponInfos Getter
func (r AlibabaOmniSaasOrderCreateRequest) GetCouponInfos() []CouponInfo {
    return r.couponInfos
}
// StoreIdType Setter
// PLACE：淘宝商户中心门店ID；CUSTOM：商户自有门店编码，需要维护到淘宝商户中心
func (r *AlibabaOmniSaasOrderCreateRequest) SetStoreIdType(storeIdType string) error {
    r.storeIdType = storeIdType
    r.Set("store_id_type", storeIdType)
    return nil
}

// StoreIdType Getter
func (r AlibabaOmniSaasOrderCreateRequest) GetStoreIdType() string {
    return r.storeIdType
}
// RequestNo Setter
// 请求号，用于标识一次请求
func (r *AlibabaOmniSaasOrderCreateRequest) SetRequestNo(requestNo string) error {
    r.requestNo = requestNo
    r.Set("request_no", requestNo)
    return nil
}

// RequestNo Getter
func (r AlibabaOmniSaasOrderCreateRequest) GetRequestNo() string {
    return r.requestNo
}
