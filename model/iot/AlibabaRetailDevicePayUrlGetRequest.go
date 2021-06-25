package iot

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
贩卖机支付二维链接获取 APIRequest
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

func NewAlibabaRetailDevicePayUrlGetRequest() *AlibabaRetailDevicePayUrlGetRequest{
    return &AlibabaRetailDevicePayUrlGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaRetailDevicePayUrlGetRequest) GetApiMethodName() string {
    return "alibaba.retail.device.payUrl.get"
}

func (r AlibabaRetailDevicePayUrlGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaRetailDevicePayUrlGetRequest) SetIsvOrderId(isvOrderId string) error {
    r.isvOrderId = isvOrderId
    r.Set("isv_order_id", isvOrderId)
    return nil
}

func (r AlibabaRetailDevicePayUrlGetRequest) GetIsvOrderId() string {
    return r.isvOrderId
}

func (r *AlibabaRetailDevicePayUrlGetRequest) SetBizName(bizName string) error {
    r.bizName = bizName
    r.Set("biz_name", bizName)
    return nil
}

func (r AlibabaRetailDevicePayUrlGetRequest) GetBizName() string {
    return r.bizName
}

func (r *AlibabaRetailDevicePayUrlGetRequest) SetItemId(itemId int64) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r AlibabaRetailDevicePayUrlGetRequest) GetItemId() int64 {
    return r.itemId
}

func (r *AlibabaRetailDevicePayUrlGetRequest) SetDeviceId(deviceId string) error {
    r.deviceId = deviceId
    r.Set("device_id", deviceId)
    return nil
}

func (r AlibabaRetailDevicePayUrlGetRequest) GetDeviceId() string {
    return r.deviceId
}

func (r *AlibabaRetailDevicePayUrlGetRequest) SetItemType(itemType int64) error {
    r.itemType = itemType
    r.Set("item_type", itemType)
    return nil
}

func (r AlibabaRetailDevicePayUrlGetRequest) GetItemType() int64 {
    return r.itemType
}

