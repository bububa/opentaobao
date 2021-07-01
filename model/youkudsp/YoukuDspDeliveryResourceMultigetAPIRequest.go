package youkudsp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YoukuDspDeliveryResourceMultigetAPIRequest
优酷实时批量获取可投放设备资源 API请求
youku.dsp.delivery.resource.multiget

优酷实时获取可投放设备资源信息,为第三方渠道提供素材获取人群识别的api,支持批量获取 */
type YoukuDspDeliveryResourceMultigetAPIRequest struct {
	model.Params
	// 渠道id
	_channelId int64
	// 子渠道id
	_subChannelId int64
	// 设备id串(md5加密)，多个设备逗号隔开
	_deviceIds string
	// 设备类型imei或者idfa
	_deviceIdType string
	// 投放类型push或者feed
	_deliveryType string
}

// New
