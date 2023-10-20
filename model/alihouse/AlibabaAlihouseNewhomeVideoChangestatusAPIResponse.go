package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomevideochangestatusAPIResponse 视频草稿状态更新 API返回值
// alibaba.alihouse.newhome.video.changestatus
//
// 视频草稿状态更新
type AlibabaalihousenewhomevideochangestatusAPIResponse struct {
	model.CommonResponse
	AlibabaalihousenewhomevideochangestatusAPIResponseModel
}

// AlibabaalihousenewhomevideochangestatusAPIResponseModel is 视频草稿状态更新 成功返回结果
type AlibabaalihousenewhomevideochangestatusAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_video_changestatus_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaalihousenewhomevideochangestatusResult `json:"result,omitempty" xml:"result,omitempty"`
}
