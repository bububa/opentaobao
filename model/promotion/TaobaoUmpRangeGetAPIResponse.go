package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUmpRangeGetAPIResponse 查询活动范围 API返回值
// taobao.ump.range.get
//
// 查询某个卖家所有参加或者不参加某项活动的物品
type TaobaoUmpRangeGetAPIResponse struct {
	model.CommonResponse
	TaobaoUmpRangeGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUmpRangeGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUmpRangeGetAPIResponseModel).Reset()
}

// TaobaoUmpRangeGetAPIResponseModel is 查询活动范围 成功返回结果
type TaobaoUmpRangeGetAPIResponseModel struct {
	XMLName xml.Name `xml:"ump_range_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 营销范围类列表！
	Ranges []Range `json:"ranges,omitempty" xml:"ranges>range,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUmpRangeGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Ranges = m.Ranges[:0]
}

var poolTaobaoUmpRangeGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUmpRangeGetAPIResponse)
	},
}

// GetTaobaoUmpRangeGetAPIResponse 从 sync.Pool 获取 TaobaoUmpRangeGetAPIResponse
func GetTaobaoUmpRangeGetAPIResponse() *TaobaoUmpRangeGetAPIResponse {
	return poolTaobaoUmpRangeGetAPIResponse.Get().(*TaobaoUmpRangeGetAPIResponse)
}

// ReleaseTaobaoUmpRangeGetAPIResponse 将 TaobaoUmpRangeGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUmpRangeGetAPIResponse(v *TaobaoUmpRangeGetAPIResponse) {
	v.Reset()
	poolTaobaoUmpRangeGetAPIResponse.Put(v)
}
