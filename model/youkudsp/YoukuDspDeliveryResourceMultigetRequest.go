package youkudsp

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
优酷实时批量获取可投放设备资源 API请求
youku.dsp.delivery.resource.multiget

优酷实时获取可投放设备资源信息,为第三方渠道提供素材获取人群识别的api,支持批量获取
*/
type YoukuDspDeliveryResourceMultigetRequest struct {
    model.Params
    // 渠道id
    _channelId   int64
    // 子渠道id
    _subChannelId   int64
    // 设备id串(md5加密)，多个设备逗号隔开
    _deviceIds   string
    // 设备类型imei或者idfa
    _deviceIdType   string
    // 投放类型push或者feed
    _deliveryType   string
}

// 初始化YoukuDspDeliveryResourceMultigetRequest对象
func NewYoukuDspDeliveryResourceMultigetRequest() *YoukuDspDeliveryResourceMultigetRequest{
    return &YoukuDspDeliveryResourceMultigetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YoukuDspDeliveryResourceMultigetRequest) GetApiMethodName() string {
    return "youku.dsp.delivery.resource.multiget"
}

// IRequest interface 方法, 获取API参数
func (r YoukuDspDeliveryResourceMultigetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ChannelId Setter
// 渠道id
func (r *YoukuDspDeliveryResourceMultigetRequest) SetChannelId(_channelId int64) error {
    r._channelId = _channelId
    r.Set("channel_id", _channelId)
    return nil
}

// ChannelId Getter
func (r YoukuDspDeliveryResourceMultigetRequest) GetChannelId() int64 {
    return r._channelId
}
// SubChannelId Setter
// 子渠道id
func (r *YoukuDspDeliveryResourceMultigetRequest) SetSubChannelId(_subChannelId int64) error {
    r._subChannelId = _subChannelId
    r.Set("sub_channel_id", _subChannelId)
    return nil
}

// SubChannelId Getter
func (r YoukuDspDeliveryResourceMultigetRequest) GetSubChannelId() int64 {
    return r._subChannelId
}
// DeviceIds Setter
// 设备id串(md5加密)，多个设备逗号隔开
func (r *YoukuDspDeliveryResourceMultigetRequest) SetDeviceIds(_deviceIds string) error {
    r._deviceIds = _deviceIds
    r.Set("device_ids", _deviceIds)
    return nil
}

// DeviceIds Getter
func (r YoukuDspDeliveryResourceMultigetRequest) GetDeviceIds() string {
    return r._deviceIds
}
// DeviceIdType Setter
// 设备类型imei或者idfa
func (r *YoukuDspDeliveryResourceMultigetRequest) SetDeviceIdType(_deviceIdType string) error {
    r._deviceIdType = _deviceIdType
    r.Set("device_id_type", _deviceIdType)
    return nil
}

// DeviceIdType Getter
func (r YoukuDspDeliveryResourceMultigetRequest) GetDeviceIdType() string {
    return r._deviceIdType
}
// DeliveryType Setter
// 投放类型push或者feed
func (r *YoukuDspDeliveryResourceMultigetRequest) SetDeliveryType(_deliveryType string) error {
    r._deliveryType = _deliveryType
    r.Set("delivery_type", _deliveryType)
    return nil
}

// DeliveryType Getter
func (r YoukuDspDeliveryResourceMultigetRequest) GetDeliveryType() string {
    return r._deliveryType
}
