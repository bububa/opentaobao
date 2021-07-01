package brandhub

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBrandStartshopRptWordpackageGetAPIResponse
明星店铺品牌流量包报表数据查询 API返回值
taobao.brand.startshop.rpt.wordpackage.get

获取明星店铺广告词包分日报表数据，只能查询近90天内的数据，包括展现量，点击量等 */
type TaobaoBrandStartshopRptWordpackageGetAPIResponse struct {
	model.CommonResponse
	TaobaoBrandStartshopRptWordpackageGetAPIResponseModel
}

// TaobaoBrandStartshopRptWordpackageGetAPIResponseModel is 明星店铺品牌流量包报表数据查询 成功返回结果
type TaobaoBrandStartshopRptWordpackageGetAPIResponseModel struct {
	XMLName xml.Name `xml:"brand_startshop_rpt_wordpackage_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	ErrorParam string `json:"error_param,omitempty" xml:"error_param,omitempty"`
	// 返回结果
	CampaignRptList []TaobaoBrandStartshopRptWordpackageGetResult `json:"campaign_rpt_list,omitempty" xml:"campaign_rpt_list>taobao_brand_startshop_rpt_wordpackage_get_result,omitempty"`
}
