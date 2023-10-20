package promotion

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUmpMbbsListGetAPIResponse 通过ids列表获取营销积木块列表 API返回值
// taobao.ump.mbbs.list.get
//
// 通过营销积木id列表来获取营销积木块列表。
type TaobaoUmpMbbsListGetAPIResponse struct {
	model.CommonResponse
	TaobaoUmpMbbsListGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUmpMbbsListGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUmpMbbsListGetAPIResponseModel).Reset()
}

// TaobaoUmpMbbsListGetAPIResponseModel is 通过ids列表获取营销积木块列表 成功返回结果
type TaobaoUmpMbbsListGetAPIResponseModel struct {
	XMLName xml.Name `xml:"ump_mbbs_list_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 营销积木块内容列表，内容为json格式的，可以通过ump sdk里面的MBB.fromJson来处理
	Mbbs []string `json:"mbbs,omitempty" xml:"mbbs>string,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUmpMbbsListGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Mbbs = m.Mbbs[:0]
}

var poolTaobaoUmpMbbsListGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUmpMbbsListGetAPIResponse)
	},
}

// GetTaobaoUmpMbbsListGetAPIResponse 从 sync.Pool 获取 TaobaoUmpMbbsListGetAPIResponse
func GetTaobaoUmpMbbsListGetAPIResponse() *TaobaoUmpMbbsListGetAPIResponse {
	return poolTaobaoUmpMbbsListGetAPIResponse.Get().(*TaobaoUmpMbbsListGetAPIResponse)
}

// ReleaseTaobaoUmpMbbsListGetAPIResponse 将 TaobaoUmpMbbsListGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUmpMbbsListGetAPIResponse(v *TaobaoUmpMbbsListGetAPIResponse) {
	v.Reset()
	poolTaobaoUmpMbbsListGetAPIResponse.Put(v)
}
