package mc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallmcrecordordersyncAPIRequest 订单信息同步 API请求
// tmall.mc.record.order.sync
//
// 订单信息同步(零售云接口)
type TmallmcrecordordersyncAPIRequest struct {
	model.Params
	// 设备编码
	_deviceCode string
	// 用户openId
	_openId string
	// 核销结果
	_result string
	// 云码版本号
	_version string
	// 原价
	_originPrice int64
	// 实付价
	_payPrice int64
}

// NewTmallmcrecordordersyncRequest 初始化TmallmcrecordordersyncAPIRequest对象
func NewTmallmcrecordordersyncRequest() *TmallmcrecordordersyncAPIRequest {
	return &TmallmcrecordordersyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallmcrecordordersyncAPIRequest) GetApiMethodName() string {
	return "tmall.mc.record.order.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallmcrecordordersyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallmcrecordordersyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeviceCode is DeviceCode Setter
// 设备编码
func (r *TmallmcrecordordersyncAPIRequest) SetDeviceCode(_deviceCode string) error {
	r._deviceCode = _deviceCode
	r.Set("device_code", _deviceCode)
	return nil
}

// GetDeviceCode DeviceCode Getter
func (r TmallmcrecordordersyncAPIRequest) GetDeviceCode() string {
	return r._deviceCode
}

// SetOpenId is OpenId Setter
// 用户openId
func (r *TmallmcrecordordersyncAPIRequest) SetOpenId(_openId string) error {
	r._openId = _openId
	r.Set("open_id", _openId)
	return nil
}

// GetOpenId OpenId Getter
func (r TmallmcrecordordersyncAPIRequest) GetOpenId() string {
	return r._openId
}

// SetResult is Result Setter
// 核销结果
func (r *TmallmcrecordordersyncAPIRequest) SetResult(_result string) error {
	r._result = _result
	r.Set("result", _result)
	return nil
}

// GetResult Result Getter
func (r TmallmcrecordordersyncAPIRequest) GetResult() string {
	return r._result
}

// SetVersion is Version Setter
// 云码版本号
func (r *TmallmcrecordordersyncAPIRequest) SetVersion(_version string) error {
	r._version = _version
	r.Set("version", _version)
	return nil
}

// GetVersion Version Getter
func (r TmallmcrecordordersyncAPIRequest) GetVersion() string {
	return r._version
}

// SetOriginPrice is OriginPrice Setter
// 原价
func (r *TmallmcrecordordersyncAPIRequest) SetOriginPrice(_originPrice int64) error {
	r._originPrice = _originPrice
	r.Set("origin_price", _originPrice)
	return nil
}

// GetOriginPrice OriginPrice Getter
func (r TmallmcrecordordersyncAPIRequest) GetOriginPrice() int64 {
	return r._originPrice
}

// SetPayPrice is PayPrice Setter
// 实付价
func (r *TmallmcrecordordersyncAPIRequest) SetPayPrice(_payPrice int64) error {
	r._payPrice = _payPrice
	r.Set("pay_price", _payPrice)
	return nil
}

// GetPayPrice PayPrice Getter
func (r TmallmcrecordordersyncAPIRequest) GetPayPrice() int64 {
	return r._payPrice
}
