package scs

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOnebpDkxCreativeCreativeReportRealtimeAPIResponse 获取创意实时报表 API返回值
// taobao.onebp.dkx.creative.creative.report.realtime
//
// 获取创意实时报表。场景和bizCode的对应关系为：拉新快adStrategyDkx，上新快adStrategyShangXin ，货品加速adStrategyProductSpeed，入会快adStrategyRuHui，预热蓄水adStrategyYuRe，爆发收割adStrategyBaoFa
type TaobaoOnebpDkxCreativeCreativeReportRealtimeAPIResponse struct {
	model.CommonResponse
	TaobaoOnebpDkxCreativeCreativeReportRealtimeAPIResponseModel
}

// TaobaoOnebpDkxCreativeCreativeReportRealtimeAPIResponseModel is 获取创意实时报表 成功返回结果
type TaobaoOnebpDkxCreativeCreativeReportRealtimeAPIResponseModel struct {
	XMLName xml.Name `xml:"onebp_dkx_creative_creative_report_realtime_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoOnebpDkxCreativeCreativeReportRealtimeResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
