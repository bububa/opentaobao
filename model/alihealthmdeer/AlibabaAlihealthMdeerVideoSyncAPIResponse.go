package alihealthmdeer

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthmdeervideosyncAPIResponse 合作伙伴视频同步给医知鹿接口 API返回值
// alibaba.alihealth.mdeer.video.sync
//
// 合伙做伴内容同步接口，用来视频同步
type AlibabaalihealthmdeervideosyncAPIResponse struct {
	model.CommonResponse
	AlibabaalihealthmdeervideosyncAPIResponseModel
}

// AlibabaalihealthmdeervideosyncAPIResponseModel is 合作伙伴视频同步给医知鹿接口 成功返回结果
type AlibabaalihealthmdeervideosyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_mdeer_video_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
