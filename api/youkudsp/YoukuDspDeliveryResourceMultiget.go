package youkudsp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/youkudsp"
)

// YoukuDspDeliveryResourceMultiget 优酷实时批量获取可投放设备资源
// youku.dsp.delivery.resource.multiget
//
// 优酷实时获取可投放设备资源信息,为第三方渠道提供素材获取人群识别的api,支持批量获取
func YoukuDspDeliveryResourceMultiget(clt *core.SDKClient, req *youkudsp.YoukuDspDeliveryResourceMultigetAPIRequest, resp *youkudsp.YoukuDspDeliveryResourceMultigetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
