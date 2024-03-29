package viapi

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliyunViapiImagesegSegmentCommonimageAPIResponse 通用分割 API返回值
// aliyun.viapi.imageseg.segment.commonimage
//
// 识别输入图像中的视觉中心物体轮廓，与背景进行分离，返回分割后的前景物体图（4通道png透明图）。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html )
type AliyunViapiImagesegSegmentCommonimageAPIResponse struct {
	model.CommonResponse
	AliyunViapiImagesegSegmentCommonimageAPIResponseModel
}

// Reset 清空结构体
func (m *AliyunViapiImagesegSegmentCommonimageAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliyunViapiImagesegSegmentCommonimageAPIResponseModel).Reset()
}

// AliyunViapiImagesegSegmentCommonimageAPIResponseModel is 通用分割 成功返回结果
type AliyunViapiImagesegSegmentCommonimageAPIResponseModel struct {
	XMLName xml.Name `xml:"aliyun_viapi_imageseg_segment_commonimage_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求ID
	TaobaoRequestId string `json:"taobao_request_id,omitempty" xml:"taobao_request_id,omitempty"`
	// 系统自动生成
	Data *AliyunViapiImagesegSegmentCommonimageData `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *AliyunViapiImagesegSegmentCommonimageAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TaobaoRequestId = ""
	m.Data = nil
}

var poolAliyunViapiImagesegSegmentCommonimageAPIResponse = sync.Pool{
	New: func() any {
		return new(AliyunViapiImagesegSegmentCommonimageAPIResponse)
	},
}

// GetAliyunViapiImagesegSegmentCommonimageAPIResponse 从 sync.Pool 获取 AliyunViapiImagesegSegmentCommonimageAPIResponse
func GetAliyunViapiImagesegSegmentCommonimageAPIResponse() *AliyunViapiImagesegSegmentCommonimageAPIResponse {
	return poolAliyunViapiImagesegSegmentCommonimageAPIResponse.Get().(*AliyunViapiImagesegSegmentCommonimageAPIResponse)
}

// ReleaseAliyunViapiImagesegSegmentCommonimageAPIResponse 将 AliyunViapiImagesegSegmentCommonimageAPIResponse 保存到 sync.Pool
func ReleaseAliyunViapiImagesegSegmentCommonimageAPIResponse(v *AliyunViapiImagesegSegmentCommonimageAPIResponse) {
	v.Reset()
	poolAliyunViapiImagesegSegmentCommonimageAPIResponse.Put(v)
}
