package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpCampaigngroupFindlistAPIResponse 查询计划组列表 API返回值
// taobao.universalbp.campaigngroup.findlist
//
// 查询某个场景内的计划组列表
type TaobaoUniversalbpCampaigngroupFindlistAPIResponse struct {
	model.CommonResponse
	TaobaoUniversalbpCampaigngroupFindlistAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUniversalbpCampaigngroupFindlistAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUniversalbpCampaigngroupFindlistAPIResponseModel).Reset()
}

// TaobaoUniversalbpCampaigngroupFindlistAPIResponseModel is 查询计划组列表 成功返回结果
type TaobaoUniversalbpCampaigngroupFindlistAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_campaigngroup_findlist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoUniversalbpCampaigngroupFindlistTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUniversalbpCampaigngroupFindlistAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoUniversalbpCampaigngroupFindlistAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpCampaigngroupFindlistAPIResponse)
	},
}

// GetTaobaoUniversalbpCampaigngroupFindlistAPIResponse 从 sync.Pool 获取 TaobaoUniversalbpCampaigngroupFindlistAPIResponse
func GetTaobaoUniversalbpCampaigngroupFindlistAPIResponse() *TaobaoUniversalbpCampaigngroupFindlistAPIResponse {
	return poolTaobaoUniversalbpCampaigngroupFindlistAPIResponse.Get().(*TaobaoUniversalbpCampaigngroupFindlistAPIResponse)
}

// ReleaseTaobaoUniversalbpCampaigngroupFindlistAPIResponse 将 TaobaoUniversalbpCampaigngroupFindlistAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUniversalbpCampaigngroupFindlistAPIResponse(v *TaobaoUniversalbpCampaigngroupFindlistAPIResponse) {
	v.Reset()
	poolTaobaoUniversalbpCampaigngroupFindlistAPIResponse.Put(v)
}
