package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaSerchcrowdStateBatchUpdateAPIResponse 单品搜索人群修改状态 API返回值
// taobao.simba.serchcrowd.state.batch.update
//
// 暂停或启用单品推广搜索人群溢价
type TaobaoSimbaSerchcrowdStateBatchUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaSerchcrowdStateBatchUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaSerchcrowdStateBatchUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaSerchcrowdStateBatchUpdateAPIResponseModel).Reset()
}

// TaobaoSimbaSerchcrowdStateBatchUpdateAPIResponseModel is 单品搜索人群修改状态 成功返回结果
type TaobaoSimbaSerchcrowdStateBatchUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_serchcrowd_state_batch_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Adgrouptargetingtags []AdgroupTargetingTagDto `json:"adgrouptargetingtags,omitempty" xml:"adgrouptargetingtags>adgroup_targeting_tag_dto,omitempty"`
	// 部分失败时返回错误List
	ErrorList []string `json:"error_list,omitempty" xml:"error_list>string,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaSerchcrowdStateBatchUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Adgrouptargetingtags = m.Adgrouptargetingtags[:0]
	m.ErrorList = m.ErrorList[:0]
}

var poolTaobaoSimbaSerchcrowdStateBatchUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaSerchcrowdStateBatchUpdateAPIResponse)
	},
}

// GetTaobaoSimbaSerchcrowdStateBatchUpdateAPIResponse 从 sync.Pool 获取 TaobaoSimbaSerchcrowdStateBatchUpdateAPIResponse
func GetTaobaoSimbaSerchcrowdStateBatchUpdateAPIResponse() *TaobaoSimbaSerchcrowdStateBatchUpdateAPIResponse {
	return poolTaobaoSimbaSerchcrowdStateBatchUpdateAPIResponse.Get().(*TaobaoSimbaSerchcrowdStateBatchUpdateAPIResponse)
}

// ReleaseTaobaoSimbaSerchcrowdStateBatchUpdateAPIResponse 将 TaobaoSimbaSerchcrowdStateBatchUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaSerchcrowdStateBatchUpdateAPIResponse(v *TaobaoSimbaSerchcrowdStateBatchUpdateAPIResponse) {
	v.Reset()
	poolTaobaoSimbaSerchcrowdStateBatchUpdateAPIResponse.Put(v)
}
