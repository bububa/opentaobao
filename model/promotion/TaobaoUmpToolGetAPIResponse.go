package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUmpToolGetAPIResponse 查询工具 API返回值
// taobao.ump.tool.get
//
// 根据工具id获取一个工具对象
type TaobaoUmpToolGetAPIResponse struct {
	model.CommonResponse
	TaobaoUmpToolGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUmpToolGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUmpToolGetAPIResponseModel).Reset()
}

// TaobaoUmpToolGetAPIResponseModel is 查询工具 成功返回结果
type TaobaoUmpToolGetAPIResponseModel struct {
	XMLName xml.Name `xml:"ump_tool_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 工具信息内容，格式为json，可以通过提供给的sdk里面的MarketingBuilder来处理这个内容
	Content string `json:"content,omitempty" xml:"content,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUmpToolGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Content = ""
}

var poolTaobaoUmpToolGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUmpToolGetAPIResponse)
	},
}

// GetTaobaoUmpToolGetAPIResponse 从 sync.Pool 获取 TaobaoUmpToolGetAPIResponse
func GetTaobaoUmpToolGetAPIResponse() *TaobaoUmpToolGetAPIResponse {
	return poolTaobaoUmpToolGetAPIResponse.Get().(*TaobaoUmpToolGetAPIResponse)
}

// ReleaseTaobaoUmpToolGetAPIResponse 将 TaobaoUmpToolGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUmpToolGetAPIResponse(v *TaobaoUmpToolGetAPIResponse) {
	v.Reset()
	poolTaobaoUmpToolGetAPIResponse.Put(v)
}
