package iot

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailDevicePayUrlGetAPIRequest 贩卖机支付二维链接获取 API请求
// alibaba.retail.device.payUrl.get
//
// 贩卖机支付二维链接获取
type AlibabaRetailDevicePayUrlGetAPIRequest struct {
	model.Params
	// 外部订单id
	_isvOrderId string
	// 业务名称
	_bizName string
	// 设备sn
	_deviceId string
	// 商品id
	_itemId int64
	// 1表示商品box，0或者为空表示普通商品
	_itemType int64
}

// NewAlibabaRetailDevicePayUrlGetRequest 初始化AlibabaRetailDevicePayUrlGetAPIRequest对象
func NewAlibabaRetailDevicePayUrlGetRequest() *AlibabaRetailDevicePayUrlGetAPIRequest {
	return &AlibabaRetailDevicePayUrlGetAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaRetailDevicePayUrlGetAPIRequest) Reset() {
	r._isvOrderId = ""
	r._bizName = ""
	r._deviceId = ""
	r._itemId = 0
	r._itemType = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaRetailDevicePayUrlGetAPIRequest) GetApiMethodName() string {
	return "alibaba.retail.device.payUrl.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaRetailDevicePayUrlGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaRetailDevicePayUrlGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIsvOrderId is IsvOrderId Setter
// 外部订单id
func (r *AlibabaRetailDevicePayUrlGetAPIRequest) SetIsvOrderId(_isvOrderId string) error {
	r._isvOrderId = _isvOrderId
	r.Set("isv_order_id", _isvOrderId)
	return nil
}

// GetIsvOrderId IsvOrderId Getter
func (r AlibabaRetailDevicePayUrlGetAPIRequest) GetIsvOrderId() string {
	return r._isvOrderId
}

// SetBizName is BizName Setter
// 业务名称
func (r *AlibabaRetailDevicePayUrlGetAPIRequest) SetBizName(_bizName string) error {
	r._bizName = _bizName
	r.Set("biz_name", _bizName)
	return nil
}

// GetBizName BizName Getter
func (r AlibabaRetailDevicePayUrlGetAPIRequest) GetBizName() string {
	return r._bizName
}

// SetDeviceId is DeviceId Setter
// 设备sn
func (r *AlibabaRetailDevicePayUrlGetAPIRequest) SetDeviceId(_deviceId string) error {
	r._deviceId = _deviceId
	r.Set("device_id", _deviceId)
	return nil
}

// GetDeviceId DeviceId Getter
func (r AlibabaRetailDevicePayUrlGetAPIRequest) GetDeviceId() string {
	return r._deviceId
}

// SetItemId is ItemId Setter
// 商品id
func (r *AlibabaRetailDevicePayUrlGetAPIRequest) SetItemId(_itemId int64) error {
	r._itemId = _itemId
	r.Set("item_id", _itemId)
	return nil
}

// GetItemId ItemId Getter
func (r AlibabaRetailDevicePayUrlGetAPIRequest) GetItemId() int64 {
	return r._itemId
}

// SetItemType is ItemType Setter
// 1表示商品box，0或者为空表示普通商品
func (r *AlibabaRetailDevicePayUrlGetAPIRequest) SetItemType(_itemType int64) error {
	r._itemType = _itemType
	r.Set("item_type", _itemType)
	return nil
}

// GetItemType ItemType Getter
func (r AlibabaRetailDevicePayUrlGetAPIRequest) GetItemType() int64 {
	return r._itemType
}

var poolAlibabaRetailDevicePayUrlGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaRetailDevicePayUrlGetRequest()
	},
}

// GetAlibabaRetailDevicePayUrlGetRequest 从 sync.Pool 获取 AlibabaRetailDevicePayUrlGetAPIRequest
func GetAlibabaRetailDevicePayUrlGetAPIRequest() *AlibabaRetailDevicePayUrlGetAPIRequest {
	return poolAlibabaRetailDevicePayUrlGetAPIRequest.Get().(*AlibabaRetailDevicePayUrlGetAPIRequest)
}

// ReleaseAlibabaRetailDevicePayUrlGetAPIRequest 将 AlibabaRetailDevicePayUrlGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaRetailDevicePayUrlGetAPIRequest(v *AlibabaRetailDevicePayUrlGetAPIRequest) {
	v.Reset()
	poolAlibabaRetailDevicePayUrlGetAPIRequest.Put(v)
}
