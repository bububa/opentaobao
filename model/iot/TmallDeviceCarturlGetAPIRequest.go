package iot

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallDeviceCarturlGetAPIRequest) GetApiMethodName() string {
	return "tmall.device.carturl.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallDeviceCarturlGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ItemIds Setter
// 商品信息，格式为 商品ID_SKU ID_数量，多条记录以逗号(,)分割
func (r *TmallDeviceCarturlGetAPIRequest) SetItemIds(_itemIds []string) error {
	r._itemIds = _itemIds
	r.Set("item_ids", _itemIds)
	return nil
}

// Get ItemIds Getter
func (r TmallDeviceCarturlGetAPIRequest) GetItemIds() []string {
	return r._itemIds
}

// Set is DeviceCode Setter
// 设备业务编码
func (r *TmallDeviceCarturlGetAPIRequest) SetDeviceCode(_deviceCode string) error {
	r._deviceCode = _deviceCode
	r.Set("device_code", _deviceCode)
	return nil
}

// Get DeviceCode Getter
func (r TmallDeviceCarturlGetAPIRequest) GetDeviceCode() string {
	return r._deviceCode
}

// Set is Longterm Setter
// 是否申请长期链接
func (r *TmallDeviceCarturlGetAPIRequest) SetLongterm(_longterm bool) error {
	r._longterm = _longterm
	r.Set("longterm", _longterm)
	return nil
}

// Get Longterm Getter
func (r TmallDeviceCarturlGetAPIRequest) GetLongterm() bool {
	return r._longterm
}
