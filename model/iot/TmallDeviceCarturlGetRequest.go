package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
添加商品到购物车 API请求
tmall.device.carturl.get

获取二维码，支持添加商品到购物车
*/
type TmallDeviceCarturlGetRequest struct {
    model.Params
    // 商品信息，格式为 商品ID_SKU ID_数量，多条记录以逗号(,)分割
    itemIds   []string
    // 设备业务编码
    deviceCode   string
    // 是否申请长期链接
    longterm   bool
}

// 初始化TmallDeviceCarturlGetRequest对象
func NewTmallDeviceCarturlGetRequest() *TmallDeviceCarturlGetRequest{
    return &TmallDeviceCarturlGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallDeviceCarturlGetRequest) GetApiMethodName() string {
    return "tmall.device.carturl.get"
}

// IRequest interface 方法, 获取API参数
func (r TmallDeviceCarturlGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ItemIds Setter
// 商品信息，格式为 商品ID_SKU ID_数量，多条记录以逗号(,)分割
func (r *TmallDeviceCarturlGetRequest) SetItemIds(itemIds []string) error {
    r.itemIds = itemIds
    r.Set("item_ids", itemIds)
    return nil
}

// ItemIds Getter
func (r TmallDeviceCarturlGetRequest) GetItemIds() []string {
    return r.itemIds
}
// DeviceCode Setter
// 设备业务编码
func (r *TmallDeviceCarturlGetRequest) SetDeviceCode(deviceCode string) error {
    r.deviceCode = deviceCode
    r.Set("device_code", deviceCode)
    return nil
}

// DeviceCode Getter
func (r TmallDeviceCarturlGetRequest) GetDeviceCode() string {
    return r.deviceCode
}
// Longterm Setter
// 是否申请长期链接
func (r *TmallDeviceCarturlGetRequest) SetLongterm(longterm bool) error {
    r.longterm = longterm
    r.Set("longterm", longterm)
    return nil
}

// Longterm Getter
func (r TmallDeviceCarturlGetRequest) GetLongterm() bool {
    return r.longterm
}
