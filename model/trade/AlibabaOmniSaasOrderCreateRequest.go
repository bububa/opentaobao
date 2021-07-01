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
type AlibabaOmniSaasOrderCreateAPIRequest struct {
    model.Params
    // 商品列表
    _goodsDetails   []GoodsDetail
    // 买家标识，淘系用户或用户手机号。当支付渠道为支付宝时，此字段为淘宝会员码或支付宝付款码。(当前仅支持淘系用户，手机号下单稍后开放)
    _buyerId   string
    // 门店ID
    _storeId   string
    // 收银设备类型
    _device   string
    // 收银设备号
    _deviceNo   string
    // 收银员标识
    _operatorId   string
    // ALIPAY：支付宝用户；TAOBAO：淘宝会员码；MOBILE：手机号
    _buyerIdType   string
    // ALIPAY：支付宝付款；BANK_CARD：刷卡
    _payChannel   string
    // 商家自有优惠
    _couponInfos   []CouponInfo
    // PLACE：淘宝商户中心门店ID；CUSTOM：商户自有门店编码，需要维护到淘宝商户中心
    _storeIdType   string
    // 请求号，用于标识一次请求
    _requestNo   string
}

// 初始化AlibabaOmniSaasOrderCreateAPIRequest对象
func NewAlibabaOmniSaasOrderCreateRequest() *AlibabaOmniSaasOrderCreateAPIRequest{
    return &AlibabaOmniSaasOrderCreateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaOmniSaasOrderCreateAPIRequest) GetApiMethodName() string {
    return "alibaba.omni.saas.order.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaOmniSaasOrderCreateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// GoodsDetails Setter
// 商品列表
func (r *AlibabaOmniSaasOrderCreateAPIRequest) SetGoodsDetails(_goodsDetails []GoodsDetail) error {
    r._goodsDetails = _goodsDetails
    r.Set("goods_details", _goodsDetails)
    return nil
}

// GoodsDetails Getter
func (r AlibabaOmniSaasOrderCreateAPIRequest) GetGoodsDetails() []GoodsDetail {
    return r._goodsDetails
}
// BuyerId Setter
// 买家标识，淘系用户或用户手机号。当支付渠道为支付宝时，此字段为淘宝会员码或支付宝付款码。(当前仅支持淘系用户，手机号下单稍后开放)
func (r *AlibabaOmniSaasOrderCreateAPIRequest) SetBuyerId(_buyerId string) error {
    r._buyerId = _buyerId
    r.Set("buyer_id", _buyerId)
    return nil
}

// BuyerId Getter
func (r AlibabaOmniSaasOrderCreateAPIRequest) GetBuyerId() string {
    return r._buyerId
}
// StoreId Setter
// 门店ID
func (r *AlibabaOmniSaasOrderCreateAPIRequest) SetStoreId(_storeId string) error {
    r._storeId = _storeId
    r.Set("store_id", _storeId)
    return nil
}

// StoreId Getter
func (r AlibabaOmniSaasOrderCreateAPIRequest) GetStoreId() string {
    return r._storeId
}
// Device Setter
// 收银设备类型
func (r *AlibabaOmniSaasOrderCreateAPIRequest) SetDevice(_device string) error {
    r._device = _device
    r.Set("device", _device)
    return nil
}

// Device Getter
func (r AlibabaOmniSaasOrderCreateAPIRequest) GetDevice() string {
    return r._device
}
// DeviceNo Setter
// 收银设备号
func (r *AlibabaOmniSaasOrderCreateAPIRequest) SetDeviceNo(_deviceNo string) error {
    r._deviceNo = _deviceNo
    r.Set("device_no", _deviceNo)
    return nil
}

// DeviceNo Getter
func (r AlibabaOmniSaasOrderCreateAPIRequest) GetDeviceNo() string {
    return r._deviceNo
}
// OperatorId Setter
// 收银员标识
func (r *AlibabaOmniSaasOrderCreateAPIRequest) SetOperatorId(_operatorId string) error {
    r._operatorId = _operatorId
    r.Set("operator_id", _operatorId)
    return nil
}

// OperatorId Getter
func (r AlibabaOmniSaasOrderCreateAPIRequest) GetOperatorId() string {
    return r._operatorId
}
// BuyerIdType Setter
// ALIPAY：支付宝用户；TAOBAO：淘宝会员码；MOBILE：手机号
func (r *AlibabaOmniSaasOrderCreateAPIRequest) SetBuyerIdType(_buyerIdType string) error {
    r._buyerIdType = _buyerIdType
    r.Set("buyer_id_type", _buyerIdType)
    return nil
}

// BuyerIdType Getter
func (r AlibabaOmniSaasOrderCreateAPIRequest) GetBuyerIdType() string {
    return r._buyerIdType
}
// PayChannel Setter
// ALIPAY：支付宝付款；BANK_CARD：刷卡
func (r *AlibabaOmniSaasOrderCreateAPIRequest) SetPayChannel(_payChannel string) error {
    r._payChannel = _payChannel
    r.Set("pay_channel", _payChannel)
    return nil
}

// PayChannel Getter
func (r AlibabaOmniSaasOrderCreateAPIRequest) GetPayChannel() string {
    return r._payChannel
}
// CouponInfos Setter
// 商家自有优惠
func (r *AlibabaOmniSaasOrderCreateAPIRequest) SetCouponInfos(_couponInfos []CouponInfo) error {
    r._couponInfos = _couponInfos
    r.Set("coupon_infos", _couponInfos)
    return nil
}

// CouponInfos Getter
func (r AlibabaOmniSaasOrderCreateAPIRequest) GetCouponInfos() []CouponInfo {
    return r._couponInfos
}
// StoreIdType Setter
// PLACE：淘宝商户中心门店ID；CUSTOM：商户自有门店编码，需要维护到淘宝商户中心
func (r *AlibabaOmniSaasOrderCreateAPIRequest) SetStoreIdType(_storeIdType string) error {
    r._storeIdType = _storeIdType
    r.Set("store_id_type", _storeIdType)
    return nil
}

// StoreIdType Getter
func (r AlibabaOmniSaasOrderCreateAPIRequest) GetStoreIdType() string {
    return r._storeIdType
}
// RequestNo Setter
// 请求号，用于标识一次请求
func (r *AlibabaOmniSaasOrderCreateAPIRequest) SetRequestNo(_requestNo string) error {
    r._requestNo = _requestNo
    r.Set("request_no", _requestNo)
    return nil
}

// RequestNo Getter
func (r AlibabaOmniSaasOrderCreateAPIRequest) GetRequestNo() string {
    return r._requestNo
}
