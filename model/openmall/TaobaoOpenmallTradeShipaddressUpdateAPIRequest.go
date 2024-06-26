package openmall

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenmallTradeShipaddressUpdateAPIRequest Openmall订单收货地址修改 API请求
// taobao.openmall.trade.shipaddress.update
//
// Openmall订单收货地址修改
type TaobaoOpenmallTradeShipaddressUpdateAPIRequest struct {
	model.Params
	// 媒体渠道，代表分销者的身份，签约支付宝代扣的渠道商淘宝账号nick
	_distributor string
	// 收货地址。最大长度为228个字节。
	_receiverAddress string
	// 城市。最大长度为32个字节。如：杭州
	_receiverCity string
	// 区/县。最大长度为32个字节。如：西湖区
	_receiverDistrict string
	// 移动电话。最大长度为11个字节。
	_receiverMobile string
	// 收货人全名。最大长度为50个字节。
	_receiverName string
	// 固定电话。最大长度为30个字节。
	_receiverPhone string
	// 省份。最大长度为32个字节。如：浙江
	_receiverState string
	// 6位数字的邮编
	_receiverZip string
	// 淘宝订单号
	_tid int64
}

// NewTaobaoOpenmallTradeShipaddressUpdateRequest 初始化TaobaoOpenmallTradeShipaddressUpdateAPIRequest对象
func NewTaobaoOpenmallTradeShipaddressUpdateRequest() *TaobaoOpenmallTradeShipaddressUpdateAPIRequest {
	return &TaobaoOpenmallTradeShipaddressUpdateAPIRequest{
		Params: model.NewParams(10),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOpenmallTradeShipaddressUpdateAPIRequest) Reset() {
	r._distributor = ""
	r._receiverAddress = ""
	r._receiverCity = ""
	r._receiverDistrict = ""
	r._receiverMobile = ""
	r._receiverName = ""
	r._receiverPhone = ""
	r._receiverState = ""
	r._receiverZip = ""
	r._tid = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenmallTradeShipaddressUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.openmall.trade.shipaddress.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenmallTradeShipaddressUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOpenmallTradeShipaddressUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDistributor is Distributor Setter
// 媒体渠道，代表分销者的身份，签约支付宝代扣的渠道商淘宝账号nick
func (r *TaobaoOpenmallTradeShipaddressUpdateAPIRequest) SetDistributor(_distributor string) error {
	r._distributor = _distributor
	r.Set("distributor", _distributor)
	return nil
}

// GetDistributor Distributor Getter
func (r TaobaoOpenmallTradeShipaddressUpdateAPIRequest) GetDistributor() string {
	return r._distributor
}

// SetReceiverAddress is ReceiverAddress Setter
// 收货地址。最大长度为228个字节。
func (r *TaobaoOpenmallTradeShipaddressUpdateAPIRequest) SetReceiverAddress(_receiverAddress string) error {
	r._receiverAddress = _receiverAddress
	r.Set("receiver_address", _receiverAddress)
	return nil
}

// GetReceiverAddress ReceiverAddress Getter
func (r TaobaoOpenmallTradeShipaddressUpdateAPIRequest) GetReceiverAddress() string {
	return r._receiverAddress
}

// SetReceiverCity is ReceiverCity Setter
// 城市。最大长度为32个字节。如：杭州
func (r *TaobaoOpenmallTradeShipaddressUpdateAPIRequest) SetReceiverCity(_receiverCity string) error {
	r._receiverCity = _receiverCity
	r.Set("receiver_city", _receiverCity)
	return nil
}

// GetReceiverCity ReceiverCity Getter
func (r TaobaoOpenmallTradeShipaddressUpdateAPIRequest) GetReceiverCity() string {
	return r._receiverCity
}

// SetReceiverDistrict is ReceiverDistrict Setter
// 区/县。最大长度为32个字节。如：西湖区
func (r *TaobaoOpenmallTradeShipaddressUpdateAPIRequest) SetReceiverDistrict(_receiverDistrict string) error {
	r._receiverDistrict = _receiverDistrict
	r.Set("receiver_district", _receiverDistrict)
	return nil
}

// GetReceiverDistrict ReceiverDistrict Getter
func (r TaobaoOpenmallTradeShipaddressUpdateAPIRequest) GetReceiverDistrict() string {
	return r._receiverDistrict
}

// SetReceiverMobile is ReceiverMobile Setter
// 移动电话。最大长度为11个字节。
func (r *TaobaoOpenmallTradeShipaddressUpdateAPIRequest) SetReceiverMobile(_receiverMobile string) error {
	r._receiverMobile = _receiverMobile
	r.Set("receiver_mobile", _receiverMobile)
	return nil
}

// GetReceiverMobile ReceiverMobile Getter
func (r TaobaoOpenmallTradeShipaddressUpdateAPIRequest) GetReceiverMobile() string {
	return r._receiverMobile
}

// SetReceiverName is ReceiverName Setter
// 收货人全名。最大长度为50个字节。
func (r *TaobaoOpenmallTradeShipaddressUpdateAPIRequest) SetReceiverName(_receiverName string) error {
	r._receiverName = _receiverName
	r.Set("receiver_name", _receiverName)
	return nil
}

// GetReceiverName ReceiverName Getter
func (r TaobaoOpenmallTradeShipaddressUpdateAPIRequest) GetReceiverName() string {
	return r._receiverName
}

// SetReceiverPhone is ReceiverPhone Setter
// 固定电话。最大长度为30个字节。
func (r *TaobaoOpenmallTradeShipaddressUpdateAPIRequest) SetReceiverPhone(_receiverPhone string) error {
	r._receiverPhone = _receiverPhone
	r.Set("receiver_phone", _receiverPhone)
	return nil
}

// GetReceiverPhone ReceiverPhone Getter
func (r TaobaoOpenmallTradeShipaddressUpdateAPIRequest) GetReceiverPhone() string {
	return r._receiverPhone
}

// SetReceiverState is ReceiverState Setter
// 省份。最大长度为32个字节。如：浙江
func (r *TaobaoOpenmallTradeShipaddressUpdateAPIRequest) SetReceiverState(_receiverState string) error {
	r._receiverState = _receiverState
	r.Set("receiver_state", _receiverState)
	return nil
}

// GetReceiverState ReceiverState Getter
func (r TaobaoOpenmallTradeShipaddressUpdateAPIRequest) GetReceiverState() string {
	return r._receiverState
}

// SetReceiverZip is ReceiverZip Setter
// 6位数字的邮编
func (r *TaobaoOpenmallTradeShipaddressUpdateAPIRequest) SetReceiverZip(_receiverZip string) error {
	r._receiverZip = _receiverZip
	r.Set("receiver_zip", _receiverZip)
	return nil
}

// GetReceiverZip ReceiverZip Getter
func (r TaobaoOpenmallTradeShipaddressUpdateAPIRequest) GetReceiverZip() string {
	return r._receiverZip
}

// SetTid is Tid Setter
// 淘宝订单号
func (r *TaobaoOpenmallTradeShipaddressUpdateAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoOpenmallTradeShipaddressUpdateAPIRequest) GetTid() int64 {
	return r._tid
}

var poolTaobaoOpenmallTradeShipaddressUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOpenmallTradeShipaddressUpdateRequest()
	},
}

// GetTaobaoOpenmallTradeShipaddressUpdateRequest 从 sync.Pool 获取 TaobaoOpenmallTradeShipaddressUpdateAPIRequest
func GetTaobaoOpenmallTradeShipaddressUpdateAPIRequest() *TaobaoOpenmallTradeShipaddressUpdateAPIRequest {
	return poolTaobaoOpenmallTradeShipaddressUpdateAPIRequest.Get().(*TaobaoOpenmallTradeShipaddressUpdateAPIRequest)
}

// ReleaseTaobaoOpenmallTradeShipaddressUpdateAPIRequest 将 TaobaoOpenmallTradeShipaddressUpdateAPIRequest 放入 sync.Pool
func ReleaseTaobaoOpenmallTradeShipaddressUpdateAPIRequest(v *TaobaoOpenmallTradeShipaddressUpdateAPIRequest) {
	v.Reset()
	poolTaobaoOpenmallTradeShipaddressUpdateAPIRequest.Put(v)
}
