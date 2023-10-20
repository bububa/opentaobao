package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpCrowdFindrecommendcrowdAPIResponse 查询推荐人群 API返回值
// taobao.universalbp.crowd.findrecommendcrowd
//
// 入参推广信息，查询推荐人群，查出的推荐人群列表，不需转换，直接可用于入参其他计划创编类接口
type TaobaoUniversalbpCrowdFindrecommendcrowdAPIResponse struct {
	model.CommonResponse
	TaobaoUniversalbpCrowdFindrecommendcrowdAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUniversalbpCrowdFindrecommendcrowdAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUniversalbpCrowdFindrecommendcrowdAPIResponseModel).Reset()
}

// TaobaoUniversalbpCrowdFindrecommendcrowdAPIResponseModel is 查询推荐人群 成功返回结果
type TaobaoUniversalbpCrowdFindrecommendcrowdAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_crowd_findrecommendcrowd_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoUniversalbpCrowdFindrecommendcrowdTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUniversalbpCrowdFindrecommendcrowdAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoUniversalbpCrowdFindrecommendcrowdAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpCrowdFindrecommendcrowdAPIResponse)
	},
}

// GetTaobaoUniversalbpCrowdFindrecommendcrowdAPIResponse 从 sync.Pool 获取 TaobaoUniversalbpCrowdFindrecommendcrowdAPIResponse
func GetTaobaoUniversalbpCrowdFindrecommendcrowdAPIResponse() *TaobaoUniversalbpCrowdFindrecommendcrowdAPIResponse {
	return poolTaobaoUniversalbpCrowdFindrecommendcrowdAPIResponse.Get().(*TaobaoUniversalbpCrowdFindrecommendcrowdAPIResponse)
}

// ReleaseTaobaoUniversalbpCrowdFindrecommendcrowdAPIResponse 将 TaobaoUniversalbpCrowdFindrecommendcrowdAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUniversalbpCrowdFindrecommendcrowdAPIResponse(v *TaobaoUniversalbpCrowdFindrecommendcrowdAPIResponse) {
	v.Reset()
	poolTaobaoUniversalbpCrowdFindrecommendcrowdAPIResponse.Put(v)
}
