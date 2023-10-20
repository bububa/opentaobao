package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUmpActivityGetAPIResponse 查询营销活动 API返回值
// taobao.ump.activity.get
//
// 查询营销活动
type TaobaoUmpActivityGetAPIResponse struct {
	model.CommonResponse
	TaobaoUmpActivityGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUmpActivityGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUmpActivityGetAPIResponseModel).Reset()
}

// TaobaoUmpActivityGetAPIResponseModel is 查询营销活动 成功返回结果
type TaobaoUmpActivityGetAPIResponseModel struct {
	XMLName xml.Name `xml:"ump_activity_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 营销活动的内容，可以通过ump sdk中的marketingTool来完成对该内容的处理
	Content string `json:"content,omitempty" xml:"content,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUmpActivityGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Content = ""
}

var poolTaobaoUmpActivityGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUmpActivityGetAPIResponse)
	},
}

// GetTaobaoUmpActivityGetAPIResponse 从 sync.Pool 获取 TaobaoUmpActivityGetAPIResponse
func GetTaobaoUmpActivityGetAPIResponse() *TaobaoUmpActivityGetAPIResponse {
	return poolTaobaoUmpActivityGetAPIResponse.Get().(*TaobaoUmpActivityGetAPIResponse)
}

// ReleaseTaobaoUmpActivityGetAPIResponse 将 TaobaoUmpActivityGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUmpActivityGetAPIResponse(v *TaobaoUmpActivityGetAPIResponse) {
	v.Reset()
	poolTaobaoUmpActivityGetAPIResponse.Put(v)
}
