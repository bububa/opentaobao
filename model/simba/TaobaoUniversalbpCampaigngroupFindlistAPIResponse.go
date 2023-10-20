package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaouniversalbpcampaigngroupfindlistAPIResponse 查询计划组列表 API返回值
// taobao.universalbp.campaigngroup.findlist
//
// 查询某个场景内的计划组列表
type TaobaouniversalbpcampaigngroupfindlistAPIResponse struct {
	model.CommonResponse
	TaobaouniversalbpcampaigngroupfindlistAPIResponseModel
}

// TaobaouniversalbpcampaigngroupfindlistAPIResponseModel is 查询计划组列表 成功返回结果
type TaobaouniversalbpcampaigngroupfindlistAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_campaigngroup_findlist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaouniversalbpcampaigngroupfindlistTopResult `json:"result,omitempty" xml:"result,omitempty"`
}
