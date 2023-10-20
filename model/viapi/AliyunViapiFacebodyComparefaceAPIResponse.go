package viapi

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliyunViapiFacebodyComparefaceAPIResponse 人脸比对1：1 API返回值
// aliyun.viapi.facebody.compareface
//
// 基于输入的两张图片，人脸比对服务可检测两张图片中的人脸，并挑选两张图片的最大人脸进行比较，判断是否是同一人；人脸比对服务还返回了这两个人脸的矩形框、比对的置信度，以及不同误识率的置信度阈值。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html )
type AliyunViapiFacebodyComparefaceAPIResponse struct {
	model.CommonResponse
	AliyunViapiFacebodyComparefaceAPIResponseModel
}

// Reset 清空结构体
func (m *AliyunViapiFacebodyComparefaceAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliyunViapiFacebodyComparefaceAPIResponseModel).Reset()
}

// AliyunViapiFacebodyComparefaceAPIResponseModel is 人脸比对1：1 成功返回结果
type AliyunViapiFacebodyComparefaceAPIResponseModel struct {
	XMLName xml.Name `xml:"aliyun_viapi_facebody_compareface_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求ID
	TaobaoRequestId string `json:"taobao_request_id,omitempty" xml:"taobao_request_id,omitempty"`
	// 系统自动生成
	Data *AliyunViapiFacebodyComparefaceData `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *AliyunViapiFacebodyComparefaceAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TaobaoRequestId = ""
	m.Data = nil
}

var poolAliyunViapiFacebodyComparefaceAPIResponse = sync.Pool{
	New: func() any {
		return new(AliyunViapiFacebodyComparefaceAPIResponse)
	},
}

// GetAliyunViapiFacebodyComparefaceAPIResponse 从 sync.Pool 获取 AliyunViapiFacebodyComparefaceAPIResponse
func GetAliyunViapiFacebodyComparefaceAPIResponse() *AliyunViapiFacebodyComparefaceAPIResponse {
	return poolAliyunViapiFacebodyComparefaceAPIResponse.Get().(*AliyunViapiFacebodyComparefaceAPIResponse)
}

// ReleaseAliyunViapiFacebodyComparefaceAPIResponse 将 AliyunViapiFacebodyComparefaceAPIResponse 保存到 sync.Pool
func ReleaseAliyunViapiFacebodyComparefaceAPIResponse(v *AliyunViapiFacebodyComparefaceAPIResponse) {
	v.Reset()
	poolAliyunViapiFacebodyComparefaceAPIResponse.Put(v)
}
