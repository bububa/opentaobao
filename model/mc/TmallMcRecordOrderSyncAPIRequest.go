package mc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallMcRecordOrderSyncAPIRequest
订单信息同步 API请求
tmall.mc.record.order.sync

订单信息同步(零售云接口) */
type TmallMcRecordOrderSyncAPIRequest struct {
	model.Params
	// 设备编码
	_deviceCode string
	// 原价
	_originPrice int64
	// 实付价
	_payPrice int64
	// 用户openId
	_openId string
	// 核销结果
	_result string
	// 云码版本号
	_version string
}

// NewTmallMcRecordOrderSyncRequest 初始化TmallMcRecordOrderSyncAPIRequest对象
func NewTmallMcRecordOrderSyncRequest() *TmallMcRecordOrderSyncAPIRequest {
	return &TmallMcRecordOrderSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallMcRecordOrderSyncAPIRequest) GetApiMethodName() string {
	return "tmall.mc.record.order.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallMcRecordOrderSyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is DeviceCode Setter
// 设备编码
func (r *TmallMcRecordOrderSyncAPIRequest) SetDeviceCode(_deviceCode string) error {
	r._deviceCode = _deviceCode
	r.Set("device_code", _deviceCode)
	return nil
}

// Get DeviceCode Getter
func (r TmallMcRecordOrderSyncAPIRequest) GetDeviceCode() string {
	return r._deviceCode
}

// Set is OriginPrice Setter
// 原价
func (r *TmallMcRecordOrderSyncAPIRequest) SetOriginPrice(_originPrice int64) error {
	r._originPrice = _originPrice
	r.Set("origin_price", _originPrice)
	return nil
}

// Get OriginPrice Getter
func (r TmallMcRecordOrderSyncAPIRequest) GetOriginPrice() int64 {
	return r._originPrice
}

// Set is PayPrice Setter
// 实付价
func (r *TmallMcRecordOrderSyncAPIRequest) SetPayPrice(_payPrice int64) error {
	r._payPrice = _payPrice
	r.Set("pay_price", _payPrice)
	return nil
}

// Get PayPrice Getter
func (r TmallMcRecordOrderSyncAPIRequest) GetPayPrice() int64 {
	return r._payPrice
}

// Set is OpenId Setter
// 用户openId
func (r *TmallMcRecordOrderSyncAPIRequest) SetOpenId(_openId string) error {
	r._openId = _openId
	r.Set("open_id", _openId)
	return nil
}

// Get OpenId Getter
func (r TmallMcRecordOrderSyncAPIRequest) GetOpenId() string {
	return r._openId
}

// Set is Result Setter
// 核销结果
func (r *TmallMcRecordOrderSyncAPIRequest) SetResult(_result string) error {
	r._result = _result
	r.Set("result", _result)
	return nil
}

// Get Result Getter
func (r TmallMcRecordOrderSyncAPIRequest) GetResult() string {
	return r._result
}

// Set is Version Setter
// 云码版本号
func (r *TmallMcRecordOrderSyncAPIRequest) SetVersion(_version string) error {
	r._version = _version
	r.Set("version", _version)
	return nil
}

// Get Version Getter
func (r TmallMcRecordOrderSyncAPIRequest) GetVersion() string {
	return r._version
}
