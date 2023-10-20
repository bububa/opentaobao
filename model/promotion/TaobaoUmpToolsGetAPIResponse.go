package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUmpToolsGetAPIResponse 查询工具列表 API返回值
// taobao.ump.tools.get
//
// 查询工具列表
type TaobaoUmpToolsGetAPIResponse struct {
	model.CommonResponse
	TaobaoUmpToolsGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUmpToolsGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUmpToolsGetAPIResponseModel).Reset()
}

// TaobaoUmpToolsGetAPIResponseModel is 查询工具列表 成功返回结果
type TaobaoUmpToolsGetAPIResponseModel struct {
	XMLName xml.Name `xml:"ump_tools_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 工具列表，单个内容为json格式，需要通过ump的sdk提供的MarketingBuilder来进行处理
	Tools []string `json:"tools,omitempty" xml:"tools>string,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUmpToolsGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Tools = m.Tools[:0]
}

var poolTaobaoUmpToolsGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUmpToolsGetAPIResponse)
	},
}

// GetTaobaoUmpToolsGetAPIResponse 从 sync.Pool 获取 TaobaoUmpToolsGetAPIResponse
func GetTaobaoUmpToolsGetAPIResponse() *TaobaoUmpToolsGetAPIResponse {
	return poolTaobaoUmpToolsGetAPIResponse.Get().(*TaobaoUmpToolsGetAPIResponse)
}

// ReleaseTaobaoUmpToolsGetAPIResponse 将 TaobaoUmpToolsGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUmpToolsGetAPIResponse(v *TaobaoUmpToolsGetAPIResponse) {
	v.Reset()
	poolTaobaoUmpToolsGetAPIResponse.Put(v)
}
