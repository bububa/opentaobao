package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaAdgroupsbycampaignidGetAPIResponse 批量得到推广计划下的推广单元 API返回值
// taobao.simba.adgroupsbycampaignid.get
//
// 根据推广计划ID分页获取推广计划下的推广单元信息
type TaobaoSimbaAdgroupsbycampaignidGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaAdgroupsbycampaignidGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaAdgroupsbycampaignidGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaAdgroupsbycampaignidGetAPIResponseModel).Reset()
}

// TaobaoSimbaAdgroupsbycampaignidGetAPIResponseModel is 批量得到推广计划下的推广单元 成功返回结果
type TaobaoSimbaAdgroupsbycampaignidGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_adgroupsbycampaignid_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回的推广组分页对象
	Adgroups *ADGroupPage `json:"adgroups,omitempty" xml:"adgroups,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaAdgroupsbycampaignidGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Adgroups = nil
}

var poolTaobaoSimbaAdgroupsbycampaignidGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaAdgroupsbycampaignidGetAPIResponse)
	},
}

// GetTaobaoSimbaAdgroupsbycampaignidGetAPIResponse 从 sync.Pool 获取 TaobaoSimbaAdgroupsbycampaignidGetAPIResponse
func GetTaobaoSimbaAdgroupsbycampaignidGetAPIResponse() *TaobaoSimbaAdgroupsbycampaignidGetAPIResponse {
	return poolTaobaoSimbaAdgroupsbycampaignidGetAPIResponse.Get().(*TaobaoSimbaAdgroupsbycampaignidGetAPIResponse)
}

// ReleaseTaobaoSimbaAdgroupsbycampaignidGetAPIResponse 将 TaobaoSimbaAdgroupsbycampaignidGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaAdgroupsbycampaignidGetAPIResponse(v *TaobaoSimbaAdgroupsbycampaignidGetAPIResponse) {
	v.Reset()
	poolTaobaoSimbaAdgroupsbycampaignidGetAPIResponse.Put(v)
}
