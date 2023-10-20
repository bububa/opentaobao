package iot

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallDeviceCarturlGetAPIRequest 添加商品到购物车 API请求
// tmall.device.carturl.get
//
// 获取二维码，支持添加商品到购物车
type TmallDeviceCarturlGetAPIRequest struct {
	model.Params
	// 商品信息，格式为 商品ID_SKU ID_数量，多条记录以逗号(,)分割
	_itemIds []string
	// 设备业务编码
	_deviceCode string
	// 是否申请长期链接
	_longterm bool
}

// NewTmallDeviceCarturlGetRequest 初始化TmallDeviceCarturlGetAPIRequest对象
func NewTmallDeviceCarturlGetRequest() *TmallDeviceCarturlGetAPIRequest {
	return &TmallDeviceCarturlGetAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallDeviceCarturlGetAPIRequest) Reset() {
	r._itemIds = r._itemIds[:0]
	r._deviceCode = ""
	r._longterm = false
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallDeviceCarturlGetAPIRequest) GetApiMethodName() string {
	return "tmall.device.carturl.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallDeviceCarturlGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallDeviceCarturlGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemIds is ItemIds Setter
// 商品信息，格式为 商品ID_SKU ID_数量，多条记录以逗号(,)分割
func (r *TmallDeviceCarturlGetAPIRequest) SetItemIds(_itemIds []string) error {
	r._itemIds = _itemIds
	r.Set("item_ids", _itemIds)
	return nil
}

// GetItemIds ItemIds Getter
func (r TmallDeviceCarturlGetAPIRequest) GetItemIds() []string {
	return r._itemIds
}

// SetDeviceCode is DeviceCode Setter
// 设备业务编码
func (r *TmallDeviceCarturlGetAPIRequest) SetDeviceCode(_deviceCode string) error {
	r._deviceCode = _deviceCode
	r.Set("device_code", _deviceCode)
	return nil
}

// GetDeviceCode DeviceCode Getter
func (r TmallDeviceCarturlGetAPIRequest) GetDeviceCode() string {
	return r._deviceCode
}

// SetLongterm is Longterm Setter
// 是否申请长期链接
func (r *TmallDeviceCarturlGetAPIRequest) SetLongterm(_longterm bool) error {
	r._longterm = _longterm
	r.Set("longterm", _longterm)
	return nil
}

// GetLongterm Longterm Getter
func (r TmallDeviceCarturlGetAPIRequest) GetLongterm() bool {
	return r._longterm
}

var poolTmallDeviceCarturlGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallDeviceCarturlGetRequest()
	},
}

// GetTmallDeviceCarturlGetRequest 从 sync.Pool 获取 TmallDeviceCarturlGetAPIRequest
func GetTmallDeviceCarturlGetAPIRequest() *TmallDeviceCarturlGetAPIRequest {
	return poolTmallDeviceCarturlGetAPIRequest.Get().(*TmallDeviceCarturlGetAPIRequest)
}

// ReleaseTmallDeviceCarturlGetAPIRequest 将 TmallDeviceCarturlGetAPIRequest 放入 sync.Pool
func ReleaseTmallDeviceCarturlGetAPIRequest(v *TmallDeviceCarturlGetAPIRequest) {
	v.Reset()
	poolTmallDeviceCarturlGetAPIRequest.Put(v)
}
