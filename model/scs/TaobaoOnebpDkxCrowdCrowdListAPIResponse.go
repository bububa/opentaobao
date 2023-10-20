package scs

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOnebpDkxCrowdCrowdListAPIResponse 获取人群信息列表 API返回值
// taobao.onebp.dkx.crowd.crowd.list
//
// 获取人群信息列表。场景和bizCode的对应关系为：拉新快adStrategyDkx，上新快adStrategyShangXin ，货品加速adStrategyProductSpeed，入会快adStrategyRuHui，预热蓄水adStrategyYuRe，爆发收割adStrategyBaoFa
type TaobaoOnebpDkxCrowdCrowdListAPIResponse struct {
	model.CommonResponse
	TaobaoOnebpDkxCrowdCrowdListAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOnebpDkxCrowdCrowdListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOnebpDkxCrowdCrowdListAPIResponseModel).Reset()
}

// TaobaoOnebpDkxCrowdCrowdListAPIResponseModel is 获取人群信息列表 成功返回结果
type TaobaoOnebpDkxCrowdCrowdListAPIResponseModel struct {
	XMLName xml.Name `xml:"onebp_dkx_crowd_crowd_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoOnebpDkxCrowdCrowdListResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOnebpDkxCrowdCrowdListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoOnebpDkxCrowdCrowdListAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOnebpDkxCrowdCrowdListAPIResponse)
	},
}

// GetTaobaoOnebpDkxCrowdCrowdListAPIResponse 从 sync.Pool 获取 TaobaoOnebpDkxCrowdCrowdListAPIResponse
func GetTaobaoOnebpDkxCrowdCrowdListAPIResponse() *TaobaoOnebpDkxCrowdCrowdListAPIResponse {
	return poolTaobaoOnebpDkxCrowdCrowdListAPIResponse.Get().(*TaobaoOnebpDkxCrowdCrowdListAPIResponse)
}

// ReleaseTaobaoOnebpDkxCrowdCrowdListAPIResponse 将 TaobaoOnebpDkxCrowdCrowdListAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOnebpDkxCrowdCrowdListAPIResponse(v *TaobaoOnebpDkxCrowdCrowdListAPIResponse) {
	v.Reset()
	poolTaobaoOnebpDkxCrowdCrowdListAPIResponse.Put(v)
}
