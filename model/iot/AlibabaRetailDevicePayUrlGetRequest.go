package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
贩卖机支付二维链接获取 API请求
alibaba.retail.device.payUrl.get

贩卖机支付二维链接获取
*/
type AlibabaRetailDevicePayUrlGetRequest struct {
    model.Params
    // 外部订单id
    isvOrderId   string
    // 业务名称
    bizName   string
    // 商品id
    itemId   int64
    // 设备sn
    deviceId   string
    // 1表示商品box，0或者为空表示普通商品
    itemType   int64
}

// 初始化AlibabaRetailDevicePayUrlGetRequest对象
func NewAlibabaRetailDevicePayUrlGetRequest() *AlibabaRetailDevicePayUrlGetRequest{
    return &AlibabaRetailDevicePayUrlGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaRetailDevicePayUrlGetRequest) GetApiMethodName() string {
    return "alibaba.retail.device.payUrl.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaRetailDevicePayUrlGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// IsvOrderId Setter
// 外部订单id
func (r *AlibabaRetailDevicePayUrlGetRequest) SetIsvOrderId(isvOrderId string) error {
    r.isvOrderId = isvOrderId
    r.Set("isv_order_id", isvOrderId)
    return nil
}

// IsvOrderId Getter
func (r AlibabaRetailDevicePayUrlGetRequest) GetIsvOrderId() string {
    return r.isvOrderId
}
// BizName Setter
// 业务名称
func (r *AlibabaRetailDevicePayUrlGetRequest) SetBizName(bizName string) error {
    r.bizName = bizName
    r.Set("biz_name", bizName)
    return nil
}

// BizName Getter
func (r AlibabaRetailDevicePayUrlGetRequest) GetBizName() string {
    return r.bizName
}
// ItemId Setter
// 商品id
func (r *AlibabaRetailDevicePayUrlGetRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

// ItemId Getter
func (r AlibabaRetailDevicePayUrlGetRequest) GetItemId() int64 {
    return r.itemId
}
// DeviceId Setter
// 设备sn
func (r *AlibabaRetailDevicePayUrlGetRequest) SetDeviceId(deviceId string) error {
    r.deviceId = deviceId
    r.Set("device_id", deviceId)
    return nil
}

// DeviceId Getter
func (r AlibabaRetailDevicePayUrlGetRequest) GetDeviceId() string {
    return r.deviceId
}
// ItemType Setter
// 1表示商品box，0或者为空表示普通商品
func (r *AlibabaRetailDevicePayUrlGetRequest) SetItemType(itemType int64) error {
    r.itemType = itemType
    r.Set("item_type", itemType)
    return nil
}

// ItemType Getter
func (r AlibabaRetailDevicePayUrlGetRequest) GetItemType() int64 {
    return r.itemType
}
