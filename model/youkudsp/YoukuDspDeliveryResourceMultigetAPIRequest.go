package youkudsp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YoukuDspDeliveryResourceMultigetAPIRequest 优酷实时批量获取可投放设备资源 API请求
// youku.dsp.delivery.resource.multiget
//
// 优酷实时获取可投放设备资源信息,为第三方渠道提供素材获取人群识别的api,支持批量获取
type YoukuDspDeliveryResourceMultigetAPIRequest struct {
	model.Params
	// 设备id串(md5加密)，多个设备逗号隔开
	_deviceIds string
	// 设备类型imei或者idfa
	_deviceIdType string
	// 投放类型push或者feed
	_deliveryType string
	// 渠道id
	_channelId int64
	// 子渠道id
	_subChannelId int64
}

// NewYoukuDspDeliveryResourceMultigetRequest 初始化YoukuDspDeliveryResourceMultigetAPIRequest对象
func NewYoukuDspDeliveryResourceMultigetRequest() *YoukuDspDeliveryResourceMultigetAPIRequest {
	return &YoukuDspDeliveryResourceMultigetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YoukuDspDeliveryResourceMultigetAPIRequest) GetApiMethodName() string {
	return "youku.dsp.delivery.resource.multiget"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YoukuDspDeliveryResourceMultigetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetDeviceIds is DeviceIds Setter
// 设备id串(md5加密)，多个设备逗号隔开
func (r *YoukuDspDeliveryResourceMultigetAPIRequest) SetDeviceIds(_deviceIds string) error {
	r._deviceIds = _deviceIds
	r.Set("device_ids", _deviceIds)
	return nil
}

// GetDeviceIds DeviceIds Getter
func (r YoukuDspDeliveryResourceMultigetAPIRequest) GetDeviceIds() string {
	return r._deviceIds
}

// SetDeviceIdType is DeviceIdType Setter
// 设备类型imei或者idfa
func (r *YoukuDspDeliveryResourceMultigetAPIRequest) SetDeviceIdType(_deviceIdType string) error {
	r._deviceIdType = _deviceIdType
	r.Set("device_id_type", _deviceIdType)
	return nil
}

// GetDeviceIdType DeviceIdType Getter
func (r YoukuDspDeliveryResourceMultigetAPIRequest) GetDeviceIdType() string {
	return r._deviceIdType
}

// SetDeliveryType is DeliveryType Setter
// 投放类型push或者feed
func (r *YoukuDspDeliveryResourceMultigetAPIRequest) SetDeliveryType(_deliveryType string) error {
	r._deliveryType = _deliveryType
	r.Set("delivery_type", _deliveryType)
	return nil
}

// GetDeliveryType DeliveryType Getter
func (r YoukuDspDeliveryResourceMultigetAPIRequest) GetDeliveryType() string {
	return r._deliveryType
}

// SetChannelId is ChannelId Setter
// 渠道id
func (r *YoukuDspDeliveryResourceMultigetAPIRequest) SetChannelId(_channelId int64) error {
	r._channelId = _channelId
	r.Set("channel_id", _channelId)
	return nil
}

// GetChannelId ChannelId Getter
func (r YoukuDspDeliveryResourceMultigetAPIRequest) GetChannelId() int64 {
	return r._channelId
}

// SetSubChannelId is SubChannelId Setter
// 子渠道id
func (r *YoukuDspDeliveryResourceMultigetAPIRequest) SetSubChannelId(_subChannelId int64) error {
	r._subChannelId = _subChannelId
	r.Set("sub_channel_id", _subChannelId)
	return nil
}

// GetSubChannelId SubChannelId Getter
func (r YoukuDspDeliveryResourceMultigetAPIRequest) GetSubChannelId() int64 {
	return r._subChannelId
}
