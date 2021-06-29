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
    _distributor   string
    // 收货地址。最大长度为228个字节。
    _receiverAddress   string
    // 城市。最大长度为32个字节。如：杭州
    _receiverCity   string
    // 区/县。最大长度为32个字节。如：西湖区
    _receiverDistrict   string
    // 移动电话。最大长度为11个字节。
    _receiverMobile   string
    // 收货人全名。最大长度为50个字节。
    _receiverName   string
    // 固定电话。最大长度为30个字节。
    _receiverPhone   string
    // 省份。最大长度为32个字节。如：浙江
    _receiverState   string
    // 6位数字的邮编
    _receiverZip   string
    // 淘宝订单号
    _tid   int64
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
func (r *TaobaoOpenmallTradeShipaddressUpdateRequest) SetDistributor(_distributor string) error {
    r._distributor = _distributor
    r.Set("distributor", _distributor)
    return nil
}

// Distributor Getter
func (r TaobaoOpenmallTradeShipaddressUpdateRequest) GetDistributor() string {
    return r._distributor
}
// ReceiverAddress Setter
// 收货地址。最大长度为228个字节。
func (r *TaobaoOpenmallTradeShipaddressUpdateRequest) SetReceiverAddress(_receiverAddress string) error {
    r._receiverAddress = _receiverAddress
    r.Set("receiver_address", _receiverAddress)
    return nil
}

// ReceiverAddress Getter
func (r TaobaoOpenmallTradeShipaddressUpdateRequest) GetReceiverAddress() string {
    return r._receiverAddress
}
// ReceiverCity Setter
// 城市。最大长度为32个字节。如：杭州
func (r *TaobaoOpenmallTradeShipaddressUpdateRequest) SetReceiverCity(_receiverCity string) error {
    r._receiverCity = _receiverCity
    r.Set("receiver_city", _receiverCity)
    return nil
}

// ReceiverCity Getter
func (r TaobaoOpenmallTradeShipaddressUpdateRequest) GetReceiverCity() string {
    return r._receiverCity
}
// ReceiverDistrict Setter
// 区/县。最大长度为32个字节。如：西湖区
func (r *TaobaoOpenmallTradeShipaddressUpdateRequest) SetReceiverDistrict(_receiverDistrict string) error {
    r._receiverDistrict = _receiverDistrict
    r.Set("receiver_district", _receiverDistrict)
    return nil
}

// ReceiverDistrict Getter
func (r TaobaoOpenmallTradeShipaddressUpdateRequest) GetReceiverDistrict() string {
    return r._receiverDistrict
}
// ReceiverMobile Setter
// 移动电话。最大长度为11个字节。
func (r *TaobaoOpenmallTradeShipaddressUpdateRequest) SetReceiverMobile(_receiverMobile string) error {
    r._receiverMobile = _receiverMobile
    r.Set("receiver_mobile", _receiverMobile)
    return nil
}

// ReceiverMobile Getter
func (r TaobaoOpenmallTradeShipaddressUpdateRequest) GetReceiverMobile() string {
    return r._receiverMobile
}
// ReceiverName Setter
// 收货人全名。最大长度为50个字节。
func (r *TaobaoOpenmallTradeShipaddressUpdateRequest) SetReceiverName(_receiverName string) error {
    r._receiverName = _receiverName
    r.Set("receiver_name", _receiverName)
    return nil
}

// ReceiverName Getter
func (r TaobaoOpenmallTradeShipaddressUpdateRequest) GetReceiverName() string {
    return r._receiverName
}
// ReceiverPhone Setter
// 固定电话。最大长度为30个字节。
func (r *TaobaoOpenmallTradeShipaddressUpdateRequest) SetReceiverPhone(_receiverPhone string) error {
    r._receiverPhone = _receiverPhone
    r.Set("receiver_phone", _receiverPhone)
    return nil
}

// ReceiverPhone Getter
func (r TaobaoOpenmallTradeShipaddressUpdateRequest) GetReceiverPhone() string {
    return r._receiverPhone
}
// ReceiverState Setter
// 省份。最大长度为32个字节。如：浙江
func (r *TaobaoOpenmallTradeShipaddressUpdateRequest) SetReceiverState(_receiverState string) error {
    r._receiverState = _receiverState
    r.Set("receiver_state", _receiverState)
    return nil
}

// ReceiverState Getter
func (r TaobaoOpenmallTradeShipaddressUpdateRequest) GetReceiverState() string {
    return r._receiverState
}
// ReceiverZip Setter
// 6位数字的邮编
func (r *TaobaoOpenmallTradeShipaddressUpdateRequest) SetReceiverZip(_receiverZip string) error {
    r._receiverZip = _receiverZip
    r.Set("receiver_zip", _receiverZip)
    return nil
}

// ReceiverZip Getter
func (r TaobaoOpenmallTradeShipaddressUpdateRequest) GetReceiverZip() string {
    return r._receiverZip
}
// Tid Setter
// 淘宝订单号
func (r *TaobaoOpenmallTradeShipaddressUpdateRequest) SetTid(_tid int64) error {
    r._tid = _tid
    r.Set("tid", _tid)
    return nil
}

// Tid Getter
func (r TaobaoOpenmallTradeShipaddressUpdateRequest) GetTid() int64 {
    return r._tid
}
