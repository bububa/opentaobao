package trade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTradeShippingaddressUpdateAPIRequest 更改交易的收货地址 API请求
// taobao.trade.shippingaddress.update
//
// 只能更新一笔交易里面的买家收货地址 <br/>只能更新发货前（即买家已付款，等待卖家发货状态）的交易的买家收货地址 <br/>更新后的发货地址可以通过taobao.trade.fullinfo.get查到 <br/>参数中所说的字节为GBK编码的（英文和数字占1字节，中文占2字节）
type TaobaoTradeShippingaddressUpdateAPIRequest struct {
	model.Params
	// 交易编号。
	_tid int64
	// 收货人全名。最大长度为50个字节。
	_receiverName string
	// 固定电话。最大长度为30个字节。
	_receiverPhone string
	// 移动电话。最大长度为11个字节。
	_receiverMobile string
	// 省份。最大长度为32个字节。如：浙江
	_receiverState string
	// 城市。最大长度为32个字节。如：杭州
	_receiverCity string
	// 区/县。最大长度为32个字节。如：西湖区
	_receiverDistrict string
	// 收货地址。最大长度为228个字节。
	_receiverAddress string
	// 邮政编码。必须由6个数字组成。
	_receiverZip string
	// 四级地址。最大长度为32个字节。如：五常街道
	_receiverTown string
}

// NewTaobaoTradeShippingaddressUpdateRequest 初始化TaobaoTradeShippingaddressUpdateAPIRequest对象
func NewTaobaoTradeShippingaddressUpdateRequest() *TaobaoTradeShippingaddressUpdateAPIRequest {
	return &TaobaoTradeShippingaddressUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTradeShippingaddressUpdateAPIRequest) GetApiMethodName() string {
	return "taobao.trade.shippingaddress.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTradeShippingaddressUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTid is Tid Setter
// 交易编号。
func (r *TaobaoTradeShippingaddressUpdateAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaoTradeShippingaddressUpdateAPIRequest) GetTid() int64 {
	return r._tid
}

// SetReceiverName is ReceiverName Setter
// 收货人全名。最大长度为50个字节。
func (r *TaobaoTradeShippingaddressUpdateAPIRequest) SetReceiverName(_receiverName string) error {
	r._receiverName = _receiverName
	r.Set("receiver_name", _receiverName)
	return nil
}

// GetReceiverName ReceiverName Getter
func (r TaobaoTradeShippingaddressUpdateAPIRequest) GetReceiverName() string {
	return r._receiverName
}

// SetReceiverPhone is ReceiverPhone Setter
// 固定电话。最大长度为30个字节。
func (r *TaobaoTradeShippingaddressUpdateAPIRequest) SetReceiverPhone(_receiverPhone string) error {
	r._receiverPhone = _receiverPhone
	r.Set("receiver_phone", _receiverPhone)
	return nil
}

// GetReceiverPhone ReceiverPhone Getter
func (r TaobaoTradeShippingaddressUpdateAPIRequest) GetReceiverPhone() string {
	return r._receiverPhone
}

// SetReceiverMobile is ReceiverMobile Setter
// 移动电话。最大长度为11个字节。
func (r *TaobaoTradeShippingaddressUpdateAPIRequest) SetReceiverMobile(_receiverMobile string) error {
	r._receiverMobile = _receiverMobile
	r.Set("receiver_mobile", _receiverMobile)
	return nil
}

// GetReceiverMobile ReceiverMobile Getter
func (r TaobaoTradeShippingaddressUpdateAPIRequest) GetReceiverMobile() string {
	return r._receiverMobile
}

// SetReceiverState is ReceiverState Setter
// 省份。最大长度为32个字节。如：浙江
func (r *TaobaoTradeShippingaddressUpdateAPIRequest) SetReceiverState(_receiverState string) error {
	r._receiverState = _receiverState
	r.Set("receiver_state", _receiverState)
	return nil
}

// GetReceiverState ReceiverState Getter
func (r TaobaoTradeShippingaddressUpdateAPIRequest) GetReceiverState() string {
	return r._receiverState
}

// SetReceiverCity is ReceiverCity Setter
// 城市。最大长度为32个字节。如：杭州
func (r *TaobaoTradeShippingaddressUpdateAPIRequest) SetReceiverCity(_receiverCity string) error {
	r._receiverCity = _receiverCity
	r.Set("receiver_city", _receiverCity)
	return nil
}

// GetReceiverCity ReceiverCity Getter
func (r TaobaoTradeShippingaddressUpdateAPIRequest) GetReceiverCity() string {
	return r._receiverCity
}

// SetReceiverDistrict is ReceiverDistrict Setter
// 区/县。最大长度为32个字节。如：西湖区
func (r *TaobaoTradeShippingaddressUpdateAPIRequest) SetReceiverDistrict(_receiverDistrict string) error {
	r._receiverDistrict = _receiverDistrict
	r.Set("receiver_district", _receiverDistrict)
	return nil
}

// GetReceiverDistrict ReceiverDistrict Getter
func (r TaobaoTradeShippingaddressUpdateAPIRequest) GetReceiverDistrict() string {
	return r._receiverDistrict
}

// SetReceiverAddress is ReceiverAddress Setter
// 收货地址。最大长度为228个字节。
func (r *TaobaoTradeShippingaddressUpdateAPIRequest) SetReceiverAddress(_receiverAddress string) error {
	r._receiverAddress = _receiverAddress
	r.Set("receiver_address", _receiverAddress)
	return nil
}

// GetReceiverAddress ReceiverAddress Getter
func (r TaobaoTradeShippingaddressUpdateAPIRequest) GetReceiverAddress() string {
	return r._receiverAddress
}

// SetReceiverZip is ReceiverZip Setter
// 邮政编码。必须由6个数字组成。
func (r *TaobaoTradeShippingaddressUpdateAPIRequest) SetReceiverZip(_receiverZip string) error {
	r._receiverZip = _receiverZip
	r.Set("receiver_zip", _receiverZip)
	return nil
}

// GetReceiverZip ReceiverZip Getter
func (r TaobaoTradeShippingaddressUpdateAPIRequest) GetReceiverZip() string {
	return r._receiverZip
}

// SetReceiverTown is ReceiverTown Setter
// 四级地址。最大长度为32个字节。如：五常街道
func (r *TaobaoTradeShippingaddressUpdateAPIRequest) SetReceiverTown(_receiverTown string) error {
	r._receiverTown = _receiverTown
	r.Set("receiver_town", _receiverTown)
	return nil
}

// GetReceiverTown ReceiverTown Getter
func (r TaobaoTradeShippingaddressUpdateAPIRequest) GetReceiverTown() string {
	return r._receiverTown
}
