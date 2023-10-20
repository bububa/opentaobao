package tbtrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotradeshippingaddressupdateAPIRequest 更改交易的收货地址 API请求
// taobao.trade.shippingaddress.update
//
// 只能更新一笔交易里面的买家收货地址 &lt;br/&gt;只能更新发货前（即买家已付款，等待卖家发货状态）的交易的买家收货地址 &lt;br/&gt;更新后的发货地址可以通过taobao.trade.fullinfo.get查到 &lt;br/&gt;参数中所说的字节为GBK编码的（英文和数字占1字节，中文占2字节）
type TaobaotradeshippingaddressupdateAPIRequest struct {
	model.Params
	// 收货人全名。最大长度为50个字节。
	_receiverName string
	// 座机号码。最大长度为30个字节。传-1表示删除
	_receiverPhone string
	// 移动电话。最大长度为11个字节。传-1表示删除
	_receiverMobile string
	// 省份。最大长度为32个字节。如：浙江
	_receiverState string
	// 城市。最大长度为32个字节。如：杭州
	_receiverCity string
	// 区/县。最大长度为32个字节。如：西湖区
	_receiverDistrict string
	// 收货地址。最大长度为228个字节。
	_receiverAddress string
	// 邮政编码。必须由6个数字组成。注：邮政编码根据地址信息自动填入，不可单独修改
	_receiverZip string
	// 四级地址。最大长度为32个字节。如：五常街道
	_receiverTown string
	// 交易编号。
	_tid int64
}

// NewTaobaotradeshippingaddressupdateRequest 初始化TaobaotradeshippingaddressupdateAPIRequest对象
func NewTaobaotradeshippingaddressupdateRequest() *TaobaotradeshippingaddressupdateAPIRequest {
	return &TaobaotradeshippingaddressupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotradeshippingaddressupdateAPIRequest) GetApiMethodName() string {
	return "taobao.trade.shippingaddress.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotradeshippingaddressupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotradeshippingaddressupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReceiverName is ReceiverName Setter
// 收货人全名。最大长度为50个字节。
func (r *TaobaotradeshippingaddressupdateAPIRequest) SetReceiverName(_receiverName string) error {
	r._receiverName = _receiverName
	r.Set("receiver_name", _receiverName)
	return nil
}

// GetReceiverName ReceiverName Getter
func (r TaobaotradeshippingaddressupdateAPIRequest) GetReceiverName() string {
	return r._receiverName
}

// SetReceiverPhone is ReceiverPhone Setter
// 座机号码。最大长度为30个字节。传-1表示删除
func (r *TaobaotradeshippingaddressupdateAPIRequest) SetReceiverPhone(_receiverPhone string) error {
	r._receiverPhone = _receiverPhone
	r.Set("receiver_phone", _receiverPhone)
	return nil
}

// GetReceiverPhone ReceiverPhone Getter
func (r TaobaotradeshippingaddressupdateAPIRequest) GetReceiverPhone() string {
	return r._receiverPhone
}

// SetReceiverMobile is ReceiverMobile Setter
// 移动电话。最大长度为11个字节。传-1表示删除
func (r *TaobaotradeshippingaddressupdateAPIRequest) SetReceiverMobile(_receiverMobile string) error {
	r._receiverMobile = _receiverMobile
	r.Set("receiver_mobile", _receiverMobile)
	return nil
}

// GetReceiverMobile ReceiverMobile Getter
func (r TaobaotradeshippingaddressupdateAPIRequest) GetReceiverMobile() string {
	return r._receiverMobile
}

// SetReceiverState is ReceiverState Setter
// 省份。最大长度为32个字节。如：浙江
func (r *TaobaotradeshippingaddressupdateAPIRequest) SetReceiverState(_receiverState string) error {
	r._receiverState = _receiverState
	r.Set("receiver_state", _receiverState)
	return nil
}

// GetReceiverState ReceiverState Getter
func (r TaobaotradeshippingaddressupdateAPIRequest) GetReceiverState() string {
	return r._receiverState
}

// SetReceiverCity is ReceiverCity Setter
// 城市。最大长度为32个字节。如：杭州
func (r *TaobaotradeshippingaddressupdateAPIRequest) SetReceiverCity(_receiverCity string) error {
	r._receiverCity = _receiverCity
	r.Set("receiver_city", _receiverCity)
	return nil
}

// GetReceiverCity ReceiverCity Getter
func (r TaobaotradeshippingaddressupdateAPIRequest) GetReceiverCity() string {
	return r._receiverCity
}

// SetReceiverDistrict is ReceiverDistrict Setter
// 区/县。最大长度为32个字节。如：西湖区
func (r *TaobaotradeshippingaddressupdateAPIRequest) SetReceiverDistrict(_receiverDistrict string) error {
	r._receiverDistrict = _receiverDistrict
	r.Set("receiver_district", _receiverDistrict)
	return nil
}

// GetReceiverDistrict ReceiverDistrict Getter
func (r TaobaotradeshippingaddressupdateAPIRequest) GetReceiverDistrict() string {
	return r._receiverDistrict
}

// SetReceiverAddress is ReceiverAddress Setter
// 收货地址。最大长度为228个字节。
func (r *TaobaotradeshippingaddressupdateAPIRequest) SetReceiverAddress(_receiverAddress string) error {
	r._receiverAddress = _receiverAddress
	r.Set("receiver_address", _receiverAddress)
	return nil
}

// GetReceiverAddress ReceiverAddress Getter
func (r TaobaotradeshippingaddressupdateAPIRequest) GetReceiverAddress() string {
	return r._receiverAddress
}

// SetReceiverZip is ReceiverZip Setter
// 邮政编码。必须由6个数字组成。注：邮政编码根据地址信息自动填入，不可单独修改
func (r *TaobaotradeshippingaddressupdateAPIRequest) SetReceiverZip(_receiverZip string) error {
	r._receiverZip = _receiverZip
	r.Set("receiver_zip", _receiverZip)
	return nil
}

// GetReceiverZip ReceiverZip Getter
func (r TaobaotradeshippingaddressupdateAPIRequest) GetReceiverZip() string {
	return r._receiverZip
}

// SetReceiverTown is ReceiverTown Setter
// 四级地址。最大长度为32个字节。如：五常街道
func (r *TaobaotradeshippingaddressupdateAPIRequest) SetReceiverTown(_receiverTown string) error {
	r._receiverTown = _receiverTown
	r.Set("receiver_town", _receiverTown)
	return nil
}

// GetReceiverTown ReceiverTown Getter
func (r TaobaotradeshippingaddressupdateAPIRequest) GetReceiverTown() string {
	return r._receiverTown
}

// SetTid is Tid Setter
// 交易编号。
func (r *TaobaotradeshippingaddressupdateAPIRequest) SetTid(_tid int64) error {
	r._tid = _tid
	r.Set("tid", _tid)
	return nil
}

// GetTid Tid Getter
func (r TaobaotradeshippingaddressupdateAPIRequest) GetTid() int64 {
	return r._tid
}
