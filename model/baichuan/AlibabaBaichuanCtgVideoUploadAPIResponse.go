package baichuan

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibababaichuanctgvideouploadAPIResponse 提供优酷的短视频入淘API API返回值
// alibaba.baichuan.ctg.video.upload
//
// 提供优酷的短视频入淘API
type AlibababaichuanctgvideouploadAPIResponse struct {
	model.CommonResponse
	AlibababaichuanctgvideouploadAPIResponseModel
}

// AlibababaichuanctgvideouploadAPIResponseModel is 提供优酷的短视频入淘API 成功返回结果
type AlibababaichuanctgvideouploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_baichuan_ctg_video_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *CtgResponse `json:"result,omitempty" xml:"result,omitempty"`
}
