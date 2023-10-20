package scs

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoonebpdkxcreativecreativereportofflineAPIResponse 获取创意离线报表 API返回值
// taobao.onebp.dkx.creative.creative.report.offline
//
// 获取创意离线报表。场景和bizCode的对应关系为：拉新快adStrategyDkx，上新快adStrategyShangXin ，货品加速adStrategyProductSpeed，入会快adStrategyRuHui，预热蓄水adStrategyYuRe，爆发收割adStrategyBaoFa
type TaobaoonebpdkxcreativecreativereportofflineAPIResponse struct {
	model.CommonResponse
	TaobaoonebpdkxcreativecreativereportofflineAPIResponseModel
}

// TaobaoonebpdkxcreativecreativereportofflineAPIResponseModel is 获取创意离线报表 成功返回结果
type TaobaoonebpdkxcreativecreativereportofflineAPIResponseModel struct {
	XMLName xml.Name `xml:"onebp_dkx_creative_creative_report_offline_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoonebpdkxcreativecreativereportofflineResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
