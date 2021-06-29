package openmall

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
Openmall订单收货地址修改 API请求
taobao.openmall.trade.shipaddress.update

Openmall订单收货地址修改
*/
type TaobaoOpenmallTradeShipaddressUpdateRequest struct {
    model.Params
    // 媒体渠道，代表分销者的身份，签约支付宝代扣的渠道商淘宝账号nick
    distributor   string
    // 收货地址。最大长度为228个字节。
    receiverAddress   string
    // 城市。最大长度为32个字节。如：杭州
    receiverCity   string
    // 区/县。最大长度为32个字节。如：西湖区
    receiverDistrict   string
    // 移动电话。最大长度为11个字节。
    receiverMobile   string
    // 收货人全名。最大长度为50个字节。
    receiverName   string
    // 固定电话。最大长度为30个字节。
    receiverPhone   string
    // 省份。最大长度为32个字节。如：浙江
    receiverState   string
    // 6位数字的邮编
    receiverZip   string
    // 淘宝订单号
    tid   int64
}

// 初始化TaobaoOpenmallTradeShipaddressUpdateRequest对象
func NewTaobaoOpenmallTradeShipaddressUpdateRequest() *TaobaoOpenmallTradeShipaddressUpdateRequest{
    return &TaobaoOpenmallTradeShipaddressUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenmallTradeShipaddressUpdateRequest) GetApiMethodName() string {
    return "taobao.openmall.trade.shipaddress.update"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenmallTradeShipaddressUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Distributor Setter
// 媒体渠道，代表分销者的身份，签约支付宝代扣的渠道商淘宝账号nick
func (r *TaobaoOpenmallTradeShipaddressUpdateRequest) SetDistributor(distributor string) error {
    r.distributor = distributor
    r.Set("distributor", distributor)
    return nil
}

// Distributor Getter
func (r TaobaoOpenmallTradeShipaddressUpdateRequest) GetDistributor() string {
    return r.distributor
}
// ReceiverAddress Setter
// 收货地址。最大长度为228个字节。
func (r *TaobaoOpenmallTradeShipaddressUpdateRequest) SetReceiverAddress(receiverAddress string) error {
    r.receiverAddress = receiverAddress
    r.Set("receiver_address", receiverAddress)
    return nil
}

// ReceiverAddress Getter
func (r TaobaoOpenmallTradeShipaddressUpdateRequest) GetReceiverAddress() string {
    return r.receiverAddress
}
// ReceiverCity Setter
// 城市。最大长度为32个字节。如：杭州
func (r *TaobaoOpenmallTradeShipaddressUpdateRequest) SetReceiverCity(receiverCity string) error {
    r.receiverCity = receiverCity
    r.Set("receiver_city", receiverCity)
    return nil
}

// ReceiverCity Getter
func (r TaobaoOpenmallTradeShipaddressUpdateRequest) GetReceiverCity() string {
    return r.receiverCity
}
// ReceiverDistrict Setter
// 区/县。最大长度为32个字节。如：西湖区
func (r *TaobaoOpenmallTradeShipaddressUpdateRequest) SetReceiverDistrict(receiverDistrict string) error {
    r.receiverDistrict = receiverDistrict
    r.Set("receiver_district", receiverDistrict)
    return nil
}

// ReceiverDistrict Getter
func (r TaobaoOpenmallTradeShipaddressUpdateRequest) GetReceiverDistrict() string {
    return r.receiverDistrict
}
// ReceiverMobile Setter
// 移动电话。最大长度为11个字节。
func (r *TaobaoOpenmallTradeShipaddressUpdateRequest) SetReceiverMobile(receiverMobile string) error {
    r.receiverMobile = receiverMobile
    r.Set("receiver_mobile", receiverMobile)
    return nil
}

// ReceiverMobile Getter
func (r TaobaoOpenmallTradeShipaddressUpdateRequest) GetReceiverMobile() string {
    return r.receiverMobile
}
// ReceiverName Setter
// 收货人全名。最大长度为50个字节。
func (r *TaobaoOpenmallTradeShipaddressUpdateRequest) SetReceiverName(receiverName string) error {
    r.receiverName = receiverName
    r.Set("receiver_name", receiverName)
    return nil
}

// ReceiverName Getter
func (r TaobaoOpenmallTradeShipaddressUpdateRequest) GetReceiverName() string {
    return r.receiverName
}
// ReceiverPhone Setter
// 固定电话。最大长度为30个字节。
func (r *TaobaoOpenmallTradeShipaddressUpdateRequest) SetReceiverPhone(receiverPhone string) error {
    r.receiverPhone = receiverPhone
    r.Set("receiver_phone", receiverPhone)
    return nil
}

// ReceiverPhone Getter
func (r TaobaoOpenmallTradeShipaddressUpdateRequest) GetReceiverPhone() string {
    return r.receiverPhone
}
// ReceiverState Setter
// 省份。最大长度为32个字节。如：浙江
func (r *TaobaoOpenmallTradeShipaddressUpdateRequest) SetReceiverState(receiverState string) error {
    r.receiverState = receiverState
    r.Set("receiver_state", receiverState)
    return nil
}

// ReceiverState Getter
func (r TaobaoOpenmallTradeShipaddressUpdateRequest) GetReceiverState() string {
    return r.receiverState
}
// ReceiverZip Setter
// 6位数字的邮编
func (r *TaobaoOpenmallTradeShipaddressUpdateRequest) SetReceiverZip(receiverZip string) error {
    r.receiverZip = receiverZip
    r.Set("receiver_zip", receiverZip)
    return nil
}

// ReceiverZip Getter
func (r TaobaoOpenmallTradeShipaddressUpdateRequest) GetReceiverZip() string {
    return r.receiverZip
}
// Tid Setter
// 淘宝订单号
func (r *TaobaoOpenmallTradeShipaddressUpdateRequest) SetTid(tid int64) error {
    r.tid = tid
    r.Set("tid", tid)
    return nil
}

// Tid Getter
func (r TaobaoOpenmallTradeShipaddressUpdateRequest) GetTid() int64 {
    return r.tid
}
