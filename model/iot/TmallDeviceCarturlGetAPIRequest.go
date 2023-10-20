package iot

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmalldevicecarturlgetAPIRequest 添加商品到购物车 API请求
// tmall.device.carturl.get
//
// 获取二维码，支持添加商品到购物车
type TmalldevicecarturlgetAPIRequest struct {
	model.Params
	// 商品信息，格式为 商品ID_SKU ID_数量，多条记录以逗号(,)分割
	_itemIds []string
	// 设备业务编码
	_deviceCode string
	// 是否申请长期链接
	_longterm bool
}

// NewTmalldevicecarturlgetRequest 初始化TmalldevicecarturlgetAPIRequest对象
func NewTmalldevicecarturlgetRequest() *TmalldevicecarturlgetAPIRequest {
	return &TmalldevicecarturlgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmalldevicecarturlgetAPIRequest) GetApiMethodName() string {
	return "tmall.device.carturl.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmalldevicecarturlgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmalldevicecarturlgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemIds is ItemIds Setter
// 商品信息，格式为 商品ID_SKU ID_数量，多条记录以逗号(,)分割
func (r *TmalldevicecarturlgetAPIRequest) SetItemIds(_itemIds []string) error {
	r._itemIds = _itemIds
	r.Set("item_ids", _itemIds)
	return nil
}

// GetItemIds ItemIds Getter
func (r TmalldevicecarturlgetAPIRequest) GetItemIds() []string {
	return r._itemIds
}

// SetDeviceCode is DeviceCode Setter
// 设备业务编码
func (r *TmalldevicecarturlgetAPIRequest) SetDeviceCode(_deviceCode string) error {
	r._deviceCode = _deviceCode
	r.Set("device_code", _deviceCode)
	return nil
}

// GetDeviceCode DeviceCode Getter
func (r TmalldevicecarturlgetAPIRequest) GetDeviceCode() string {
	return r._deviceCode
}

// SetLongterm is Longterm Setter
// 是否申请长期链接
func (r *TmalldevicecarturlgetAPIRequest) SetLongterm(_longterm bool) error {
	r._longterm = _longterm
	r.Set("longterm", _longterm)
	return nil
}

// GetLongterm Longterm Getter
func (r TmalldevicecarturlgetAPIRequest) GetLongterm() bool {
	return r._longterm
}
