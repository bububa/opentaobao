package viapi

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliyunViapiImagesegSegmentheadAPIResponse 头像分割 API返回值
// aliyun.viapi.imageseg.segmenthead
//
// 输入一张图片，对图中人头区域进行抠图解析，输出人头png透明图。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html )
type AliyunViapiImagesegSegmentheadAPIResponse struct {
	model.CommonResponse
	AliyunViapiImagesegSegmentheadAPIResponseModel
}

// Reset 清空结构体
func (m *AliyunViapiImagesegSegmentheadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliyunViapiImagesegSegmentheadAPIResponseModel).Reset()
}

// AliyunViapiImagesegSegmentheadAPIResponseModel is 头像分割 成功返回结果
type AliyunViapiImagesegSegmentheadAPIResponseModel struct {
	XMLName xml.Name `xml:"aliyun_viapi_imageseg_segmenthead_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求ID
	TaobaoRequestId string `json:"taobao_request_id,omitempty" xml:"taobao_request_id,omitempty"`
	// 系统自动生成
	Data *AliyunViapiImagesegSegmentheadData `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *AliyunViapiImagesegSegmentheadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TaobaoRequestId = ""
	m.Data = nil
}

var poolAliyunViapiImagesegSegmentheadAPIResponse = sync.Pool{
	New: func() any {
		return new(AliyunViapiImagesegSegmentheadAPIResponse)
	},
}

// GetAliyunViapiImagesegSegmentheadAPIResponse 从 sync.Pool 获取 AliyunViapiImagesegSegmentheadAPIResponse
func GetAliyunViapiImagesegSegmentheadAPIResponse() *AliyunViapiImagesegSegmentheadAPIResponse {
	return poolAliyunViapiImagesegSegmentheadAPIResponse.Get().(*AliyunViapiImagesegSegmentheadAPIResponse)
}

// ReleaseAliyunViapiImagesegSegmentheadAPIResponse 将 AliyunViapiImagesegSegmentheadAPIResponse 保存到 sync.Pool
func ReleaseAliyunViapiImagesegSegmentheadAPIResponse(v *AliyunViapiImagesegSegmentheadAPIResponse) {
	v.Reset()
	poolAliyunViapiImagesegSegmentheadAPIResponse.Put(v)
}
