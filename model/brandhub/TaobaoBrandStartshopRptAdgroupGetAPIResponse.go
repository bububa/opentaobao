package brandhub

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoBrandStartshopRptAdgroupGetAPIResponse 明星店铺推广单元报表数据查询 API返回值
// taobao.brand.startshop.rpt.adgroup.get
//
// 获取明星店铺广告adgroup分日报表数据，只能查询近90天内的数据，包括展现量，点击量等
type TaobaoBrandStartshopRptAdgroupGetAPIResponse struct {
	model.CommonResponse
	TaobaoBrandStartshopRptAdgroupGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoBrandStartshopRptAdgroupGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoBrandStartshopRptAdgroupGetAPIResponseModel).Reset()
}

// TaobaoBrandStartshopRptAdgroupGetAPIResponseModel is 明星店铺推广单元报表数据查询 成功返回结果
type TaobaoBrandStartshopRptAdgroupGetAPIResponseModel struct {
	XMLName xml.Name `xml:"brand_startshop_rpt_adgroup_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	CampaignRptList []TaobaoBrandStartshopRptAdgroupGetResult `json:"campaign_rpt_list,omitempty" xml:"campaign_rpt_list>taobao_brand_startshop_rpt_adgroup_get_result,omitempty"`
	// 错误信息
	ErrorParam string `json:"error_param,omitempty" xml:"error_param,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoBrandStartshopRptAdgroupGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.CampaignRptList = m.CampaignRptList[:0]
	m.ErrorParam = ""
}

var poolTaobaoBrandStartshopRptAdgroupGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoBrandStartshopRptAdgroupGetAPIResponse)
	},
}

// GetTaobaoBrandStartshopRptAdgroupGetAPIResponse 从 sync.Pool 获取 TaobaoBrandStartshopRptAdgroupGetAPIResponse
func GetTaobaoBrandStartshopRptAdgroupGetAPIResponse() *TaobaoBrandStartshopRptAdgroupGetAPIResponse {
	return poolTaobaoBrandStartshopRptAdgroupGetAPIResponse.Get().(*TaobaoBrandStartshopRptAdgroupGetAPIResponse)
}

// ReleaseTaobaoBrandStartshopRptAdgroupGetAPIResponse 将 TaobaoBrandStartshopRptAdgroupGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoBrandStartshopRptAdgroupGetAPIResponse(v *TaobaoBrandStartshopRptAdgroupGetAPIResponse) {
	v.Reset()
	poolTaobaoBrandStartshopRptAdgroupGetAPIResponse.Put(v)
}
