package baichuan

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaBaichuanCtgVideoUploadAPIResponse 提供优酷的短视频入淘API API返回值
// alibaba.baichuan.ctg.video.upload
//
// 提供优酷的短视频入淘API
type AlibabaBaichuanCtgVideoUploadAPIResponse struct {
	model.CommonResponse
	AlibabaBaichuanCtgVideoUploadAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaBaichuanCtgVideoUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaBaichuanCtgVideoUploadAPIResponseModel).Reset()
}

// AlibabaBaichuanCtgVideoUploadAPIResponseModel is 提供优酷的短视频入淘API 成功返回结果
type AlibabaBaichuanCtgVideoUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_baichuan_ctg_video_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *CtgResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaBaichuanCtgVideoUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaBaichuanCtgVideoUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaBaichuanCtgVideoUploadAPIResponse)
	},
}

// GetAlibabaBaichuanCtgVideoUploadAPIResponse 从 sync.Pool 获取 AlibabaBaichuanCtgVideoUploadAPIResponse
func GetAlibabaBaichuanCtgVideoUploadAPIResponse() *AlibabaBaichuanCtgVideoUploadAPIResponse {
	return poolAlibabaBaichuanCtgVideoUploadAPIResponse.Get().(*AlibabaBaichuanCtgVideoUploadAPIResponse)
}

// ReleaseAlibabaBaichuanCtgVideoUploadAPIResponse 将 AlibabaBaichuanCtgVideoUploadAPIResponse 保存到 sync.Pool
func ReleaseAlibabaBaichuanCtgVideoUploadAPIResponse(v *AlibabaBaichuanCtgVideoUploadAPIResponse) {
	v.Reset()
	poolAlibabaBaichuanCtgVideoUploadAPIResponse.Put(v)
}
