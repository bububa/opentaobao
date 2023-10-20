package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUmpDetailGetAPIResponse 查询活动详情 API返回值
// taobao.ump.detail.get
//
// 查询活动详情
type TaobaoUmpDetailGetAPIResponse struct {
	model.CommonResponse
	TaobaoUmpDetailGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUmpDetailGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUmpDetailGetAPIResponseModel).Reset()
}

// TaobaoUmpDetailGetAPIResponseModel is 查询活动详情 成功返回结果
type TaobaoUmpDetailGetAPIResponseModel struct {
	XMLName xml.Name `xml:"ump_detail_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 活动详情信息，可以通过ump sdk中的MarketingTool来进行处理
	Content string `json:"content,omitempty" xml:"content,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUmpDetailGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Content = ""
}

var poolTaobaoUmpDetailGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUmpDetailGetAPIResponse)
	},
}

// GetTaobaoUmpDetailGetAPIResponse 从 sync.Pool 获取 TaobaoUmpDetailGetAPIResponse
func GetTaobaoUmpDetailGetAPIResponse() *TaobaoUmpDetailGetAPIResponse {
	return poolTaobaoUmpDetailGetAPIResponse.Get().(*TaobaoUmpDetailGetAPIResponse)
}

// ReleaseTaobaoUmpDetailGetAPIResponse 将 TaobaoUmpDetailGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUmpDetailGetAPIResponse(v *TaobaoUmpDetailGetAPIResponse) {
	v.Reset()
	poolTaobaoUmpDetailGetAPIResponse.Put(v)
}
