package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpCrowdFindlistAPIResponse 查询人群绑定列表 API返回值
// taobao.universalbp.crowd.findlist
//
// 查询计划和单元上绑定的人群列表
type TaobaoUniversalbpCrowdFindlistAPIResponse struct {
	model.CommonResponse
	TaobaoUniversalbpCrowdFindlistAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUniversalbpCrowdFindlistAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUniversalbpCrowdFindlistAPIResponseModel).Reset()
}

// TaobaoUniversalbpCrowdFindlistAPIResponseModel is 查询人群绑定列表 成功返回结果
type TaobaoUniversalbpCrowdFindlistAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_crowd_findlist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoUniversalbpCrowdFindlistTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUniversalbpCrowdFindlistAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoUniversalbpCrowdFindlistAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpCrowdFindlistAPIResponse)
	},
}

// GetTaobaoUniversalbpCrowdFindlistAPIResponse 从 sync.Pool 获取 TaobaoUniversalbpCrowdFindlistAPIResponse
func GetTaobaoUniversalbpCrowdFindlistAPIResponse() *TaobaoUniversalbpCrowdFindlistAPIResponse {
	return poolTaobaoUniversalbpCrowdFindlistAPIResponse.Get().(*TaobaoUniversalbpCrowdFindlistAPIResponse)
}

// ReleaseTaobaoUniversalbpCrowdFindlistAPIResponse 将 TaobaoUniversalbpCrowdFindlistAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUniversalbpCrowdFindlistAPIResponse(v *TaobaoUniversalbpCrowdFindlistAPIResponse) {
	v.Reset()
	poolTaobaoUniversalbpCrowdFindlistAPIResponse.Put(v)
}
