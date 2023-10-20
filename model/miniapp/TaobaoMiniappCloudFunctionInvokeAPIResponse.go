package miniapp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappCloudFunctionInvokeAPIResponse 外部触发云函数 API返回值
// taobao.miniapp.cloud.function.invoke
//
// 用户isv从外部触发聚石塔云函数的执行。
type TaobaoMiniappCloudFunctionInvokeAPIResponse struct {
	model.CommonResponse
	TaobaoMiniappCloudFunctionInvokeAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoMiniappCloudFunctionInvokeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoMiniappCloudFunctionInvokeAPIResponseModel).Reset()
}

// TaobaoMiniappCloudFunctionInvokeAPIResponseModel is 外部触发云函数 成功返回结果
type TaobaoMiniappCloudFunctionInvokeAPIResponseModel struct {
	XMLName xml.Name `xml:"miniapp_cloud_function_invoke_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回参数，长度限制8个字符，超出部分会被截断
	Data string `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoMiniappCloudFunctionInvokeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Data = ""
}

var poolTaobaoMiniappCloudFunctionInvokeAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoMiniappCloudFunctionInvokeAPIResponse)
	},
}

// GetTaobaoMiniappCloudFunctionInvokeAPIResponse 从 sync.Pool 获取 TaobaoMiniappCloudFunctionInvokeAPIResponse
func GetTaobaoMiniappCloudFunctionInvokeAPIResponse() *TaobaoMiniappCloudFunctionInvokeAPIResponse {
	return poolTaobaoMiniappCloudFunctionInvokeAPIResponse.Get().(*TaobaoMiniappCloudFunctionInvokeAPIResponse)
}

// ReleaseTaobaoMiniappCloudFunctionInvokeAPIResponse 将 TaobaoMiniappCloudFunctionInvokeAPIResponse 保存到 sync.Pool
func ReleaseTaobaoMiniappCloudFunctionInvokeAPIResponse(v *TaobaoMiniappCloudFunctionInvokeAPIResponse) {
	v.Reset()
	poolTaobaoMiniappCloudFunctionInvokeAPIResponse.Put(v)
}
