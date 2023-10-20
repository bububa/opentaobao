package viapi

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliyunViapiGoodstechRecognizeFurnitureAttributeAPIResponse 家居属性识别 API返回值
// aliyun.viapi.goodstech.recognize.furniture.attribute
//
// 识别输入的家居模型图的风格，目前支持16种风格识别。(参数图片/链接必须通过以下方式获取: https://help.aliyun.com/document_detail/155645.html )
type AliyunViapiGoodstechRecognizeFurnitureAttributeAPIResponse struct {
	model.CommonResponse
	AliyunViapiGoodstechRecognizeFurnitureAttributeAPIResponseModel
}

// Reset 清空结构体
func (m *AliyunViapiGoodstechRecognizeFurnitureAttributeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliyunViapiGoodstechRecognizeFurnitureAttributeAPIResponseModel).Reset()
}

// AliyunViapiGoodstechRecognizeFurnitureAttributeAPIResponseModel is 家居属性识别 成功返回结果
type AliyunViapiGoodstechRecognizeFurnitureAttributeAPIResponseModel struct {
	XMLName xml.Name `xml:"aliyun_viapi_goodstech_recognize_furniture_attribute_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求ID
	TaobaoRequestId string `json:"taobao_request_id,omitempty" xml:"taobao_request_id,omitempty"`
	// 系统自动生成
	Data *AliyunViapiGoodstechRecognizeFurnitureAttributeData `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *AliyunViapiGoodstechRecognizeFurnitureAttributeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TaobaoRequestId = ""
	m.Data = nil
}

var poolAliyunViapiGoodstechRecognizeFurnitureAttributeAPIResponse = sync.Pool{
	New: func() any {
		return new(AliyunViapiGoodstechRecognizeFurnitureAttributeAPIResponse)
	},
}

// GetAliyunViapiGoodstechRecognizeFurnitureAttributeAPIResponse 从 sync.Pool 获取 AliyunViapiGoodstechRecognizeFurnitureAttributeAPIResponse
func GetAliyunViapiGoodstechRecognizeFurnitureAttributeAPIResponse() *AliyunViapiGoodstechRecognizeFurnitureAttributeAPIResponse {
	return poolAliyunViapiGoodstechRecognizeFurnitureAttributeAPIResponse.Get().(*AliyunViapiGoodstechRecognizeFurnitureAttributeAPIResponse)
}

// ReleaseAliyunViapiGoodstechRecognizeFurnitureAttributeAPIResponse 将 AliyunViapiGoodstechRecognizeFurnitureAttributeAPIResponse 保存到 sync.Pool
func ReleaseAliyunViapiGoodstechRecognizeFurnitureAttributeAPIResponse(v *AliyunViapiGoodstechRecognizeFurnitureAttributeAPIResponse) {
	v.Reset()
	poolAliyunViapiGoodstechRecognizeFurnitureAttributeAPIResponse.Put(v)
}
