package viapi

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliyunViapiImagesegSegmenthdbodyAPIResponse 高清人体分割 API返回值
// aliyun.viapi.imageseg.segmenthdbody
//
// 对输入图像中包含的人进行分割，输出结果透明图。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html )
type AliyunViapiImagesegSegmenthdbodyAPIResponse struct {
	model.CommonResponse
	AliyunViapiImagesegSegmenthdbodyAPIResponseModel
}

// Reset 清空结构体
func (m *AliyunViapiImagesegSegmenthdbodyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliyunViapiImagesegSegmenthdbodyAPIResponseModel).Reset()
}

// AliyunViapiImagesegSegmenthdbodyAPIResponseModel is 高清人体分割 成功返回结果
type AliyunViapiImagesegSegmenthdbodyAPIResponseModel struct {
	XMLName xml.Name `xml:"aliyun_viapi_imageseg_segmenthdbody_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求ID
	TaobaoRequestId string `json:"taobao_request_id,omitempty" xml:"taobao_request_id,omitempty"`
	// 系统自动生成
	Data *AliyunViapiImagesegSegmenthdbodyData `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *AliyunViapiImagesegSegmenthdbodyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TaobaoRequestId = ""
	m.Data = nil
}

var poolAliyunViapiImagesegSegmenthdbodyAPIResponse = sync.Pool{
	New: func() any {
		return new(AliyunViapiImagesegSegmenthdbodyAPIResponse)
	},
}

// GetAliyunViapiImagesegSegmenthdbodyAPIResponse 从 sync.Pool 获取 AliyunViapiImagesegSegmenthdbodyAPIResponse
func GetAliyunViapiImagesegSegmenthdbodyAPIResponse() *AliyunViapiImagesegSegmenthdbodyAPIResponse {
	return poolAliyunViapiImagesegSegmenthdbodyAPIResponse.Get().(*AliyunViapiImagesegSegmenthdbodyAPIResponse)
}

// ReleaseAliyunViapiImagesegSegmenthdbodyAPIResponse 将 AliyunViapiImagesegSegmenthdbodyAPIResponse 保存到 sync.Pool
func ReleaseAliyunViapiImagesegSegmenthdbodyAPIResponse(v *AliyunViapiImagesegSegmenthdbodyAPIResponse) {
	v.Reset()
	poolAliyunViapiImagesegSegmenthdbodyAPIResponse.Put(v)
}
