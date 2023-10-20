package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaSearchcrowdBatchAddAPIResponse 推广单元增加搜索人群 API返回值
// taobao.simba.searchcrowd.batch.add
//
// 推广单元新增搜索人群
type TaobaoSimbaSearchcrowdBatchAddAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaSearchcrowdBatchAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaSearchcrowdBatchAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaSearchcrowdBatchAddAPIResponseModel).Reset()
}

// TaobaoSimbaSearchcrowdBatchAddAPIResponseModel is 推广单元增加搜索人群 成功返回结果
type TaobaoSimbaSearchcrowdBatchAddAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_searchcrowd_batch_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 定向信息
	Adgrouptargetingtags []AdgroupTargetingTagDto `json:"adgrouptargetingtags,omitempty" xml:"adgrouptargetingtags>adgroup_targeting_tag_dto,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaSearchcrowdBatchAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Adgrouptargetingtags = m.Adgrouptargetingtags[:0]
}

var poolTaobaoSimbaSearchcrowdBatchAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaSearchcrowdBatchAddAPIResponse)
	},
}

// GetTaobaoSimbaSearchcrowdBatchAddAPIResponse 从 sync.Pool 获取 TaobaoSimbaSearchcrowdBatchAddAPIResponse
func GetTaobaoSimbaSearchcrowdBatchAddAPIResponse() *TaobaoSimbaSearchcrowdBatchAddAPIResponse {
	return poolTaobaoSimbaSearchcrowdBatchAddAPIResponse.Get().(*TaobaoSimbaSearchcrowdBatchAddAPIResponse)
}

// ReleaseTaobaoSimbaSearchcrowdBatchAddAPIResponse 将 TaobaoSimbaSearchcrowdBatchAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaSearchcrowdBatchAddAPIResponse(v *TaobaoSimbaSearchcrowdBatchAddAPIResponse) {
	v.Reset()
	poolTaobaoSimbaSearchcrowdBatchAddAPIResponse.Put(v)
}
