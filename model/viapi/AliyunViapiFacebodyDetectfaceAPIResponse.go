package viapi

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliyunViapiFacebodyDetectfaceAPIResponse 人脸检测定位 API返回值
// aliyun.viapi.facebody.detectface
//
// 输入图片之后，在人脸检测定位返回结果的基础上，识别各个检测人脸的四种属性，返回性别（男/女）、年龄、表情（笑/不笑）、眼镜（戴/不戴）；并可返回人脸的1024维深度学习特征、(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html )
type AliyunViapiFacebodyDetectfaceAPIResponse struct {
	model.CommonResponse
	AliyunViapiFacebodyDetectfaceAPIResponseModel
}

// Reset 清空结构体
func (m *AliyunViapiFacebodyDetectfaceAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliyunViapiFacebodyDetectfaceAPIResponseModel).Reset()
}

// AliyunViapiFacebodyDetectfaceAPIResponseModel is 人脸检测定位 成功返回结果
type AliyunViapiFacebodyDetectfaceAPIResponseModel struct {
	XMLName xml.Name `xml:"aliyun_viapi_facebody_detectface_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求ID
	TaobaoRequestId string `json:"taobao_request_id,omitempty" xml:"taobao_request_id,omitempty"`
	// 系统自动生成
	Data *AliyunViapiFacebodyDetectfaceData `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *AliyunViapiFacebodyDetectfaceAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TaobaoRequestId = ""
	m.Data = nil
}

var poolAliyunViapiFacebodyDetectfaceAPIResponse = sync.Pool{
	New: func() any {
		return new(AliyunViapiFacebodyDetectfaceAPIResponse)
	},
}

// GetAliyunViapiFacebodyDetectfaceAPIResponse 从 sync.Pool 获取 AliyunViapiFacebodyDetectfaceAPIResponse
func GetAliyunViapiFacebodyDetectfaceAPIResponse() *AliyunViapiFacebodyDetectfaceAPIResponse {
	return poolAliyunViapiFacebodyDetectfaceAPIResponse.Get().(*AliyunViapiFacebodyDetectfaceAPIResponse)
}

// ReleaseAliyunViapiFacebodyDetectfaceAPIResponse 将 AliyunViapiFacebodyDetectfaceAPIResponse 保存到 sync.Pool
func ReleaseAliyunViapiFacebodyDetectfaceAPIResponse(v *AliyunViapiFacebodyDetectfaceAPIResponse) {
	v.Reset()
	poolAliyunViapiFacebodyDetectfaceAPIResponse.Put(v)
}
