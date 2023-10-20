package youkudsp

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YoukuDspDeliveryResourceMultigetAPIRequest) Reset() {
	r._deviceIds = ""
	r._deviceIdType = ""
	r._deliveryType = ""
	r._channelId = 0
	r._subChannelId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YoukuDspDeliveryResourceMultigetAPIRequest) GetApiMethodName() string {
	return "youku.dsp.delivery.resource.multiget"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YoukuDspDeliveryResourceMultigetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YoukuDspDeliveryResourceMultigetAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolYoukuDspDeliveryResourceMultigetAPIRequest = sync.Pool{
	New: func() any {
		return NewYoukuDspDeliveryResourceMultigetRequest()
	},
}

// GetYoukuDspDeliveryResourceMultigetRequest 从 sync.Pool 获取 YoukuDspDeliveryResourceMultigetAPIRequest
func GetYoukuDspDeliveryResourceMultigetAPIRequest() *YoukuDspDeliveryResourceMultigetAPIRequest {
	return poolYoukuDspDeliveryResourceMultigetAPIRequest.Get().(*YoukuDspDeliveryResourceMultigetAPIRequest)
}

// ReleaseYoukuDspDeliveryResourceMultigetAPIRequest 将 YoukuDspDeliveryResourceMultigetAPIRequest 放入 sync.Pool
func ReleaseYoukuDspDeliveryResourceMultigetAPIRequest(v *YoukuDspDeliveryResourceMultigetAPIRequest) {
	v.Reset()
	poolYoukuDspDeliveryResourceMultigetAPIRequest.Put(v)
}
