package brandhub

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBrandStarshopRptTargetGetAPIResponse 明星店铺定向维度报表 API返回值
// taobao.brand.starshop.rpt.target.get
//
// 获取明星店铺定向维度分日报表数据，只能查询近90天内的数据，包括展现量，点击量等
type TaobaoBrandStarshopRptTargetGetAPIResponse struct {
	model.CommonResponse
	TaobaoBrandStarshopRptTargetGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBrandStarshopRptTargetGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBrandStarshopRptTargetGetAPIResponseModel).Reset()
}

// TaobaoBrandStarshopRptTargetGetAPIResponseModel is 明星店铺定向维度报表 成功返回结果
type TaobaoBrandStarshopRptTargetGetAPIResponseModel struct {
	XMLName xml.Name `xml:"brand_starshop_rpt_target_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	CampaignRptList []TaobaoBrandStarshopRptTargetGetResult `json:"campaign_rpt_list,omitempty" xml:"campaign_rpt_list>taobao_brand_starshop_rpt_target_get_result,omitempty"`
	// 错误信息
	ErrorParam string `json:"error_param,omitempty" xml:"error_param,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBrandStarshopRptTargetGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.CampaignRptList = m.CampaignRptList[:0]
	m.ErrorParam = ""
}

var poolTaobaoBrandStarshopRptTargetGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBrandStarshopRptTargetGetAPIResponse)
	},
}

// GetTaobaoBrandStarshopRptTargetGetAPIResponse 从 sync.Pool 获取 TaobaoBrandStarshopRptTargetGetAPIResponse
func GetTaobaoBrandStarshopRptTargetGetAPIResponse() *TaobaoBrandStarshopRptTargetGetAPIResponse {
	return poolTaobaoBrandStarshopRptTargetGetAPIResponse.Get().(*TaobaoBrandStarshopRptTargetGetAPIResponse)
}

// ReleaseTaobaoBrandStarshopRptTargetGetAPIResponse 将 TaobaoBrandStarshopRptTargetGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBrandStarshopRptTargetGetAPIResponse(v *TaobaoBrandStarshopRptTargetGetAPIResponse) {
	v.Reset()
	poolTaobaoBrandStarshopRptTargetGetAPIResponse.Put(v)
}
