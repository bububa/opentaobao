package scs

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOnebpDkxCreativeCreativeReportOfflineAPIResponse 获取创意离线报表 API返回值
// taobao.onebp.dkx.creative.creative.report.offline
//
// 获取创意离线报表。场景和bizCode的对应关系为：拉新快adStrategyDkx，上新快adStrategyShangXin ，货品加速adStrategyProductSpeed，入会快adStrategyRuHui，预热蓄水adStrategyYuRe，爆发收割adStrategyBaoFa
type TaobaoOnebpDkxCreativeCreativeReportOfflineAPIResponse struct {
	model.CommonResponse
	TaobaoOnebpDkxCreativeCreativeReportOfflineAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOnebpDkxCreativeCreativeReportOfflineAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOnebpDkxCreativeCreativeReportOfflineAPIResponseModel).Reset()
}

// TaobaoOnebpDkxCreativeCreativeReportOfflineAPIResponseModel is 获取创意离线报表 成功返回结果
type TaobaoOnebpDkxCreativeCreativeReportOfflineAPIResponseModel struct {
	XMLName xml.Name `xml:"onebp_dkx_creative_creative_report_offline_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoOnebpDkxCreativeCreativeReportOfflineResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOnebpDkxCreativeCreativeReportOfflineAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoOnebpDkxCreativeCreativeReportOfflineAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOnebpDkxCreativeCreativeReportOfflineAPIResponse)
	},
}

// GetTaobaoOnebpDkxCreativeCreativeReportOfflineAPIResponse 从 sync.Pool 获取 TaobaoOnebpDkxCreativeCreativeReportOfflineAPIResponse
func GetTaobaoOnebpDkxCreativeCreativeReportOfflineAPIResponse() *TaobaoOnebpDkxCreativeCreativeReportOfflineAPIResponse {
	return poolTaobaoOnebpDkxCreativeCreativeReportOfflineAPIResponse.Get().(*TaobaoOnebpDkxCreativeCreativeReportOfflineAPIResponse)
}

// ReleaseTaobaoOnebpDkxCreativeCreativeReportOfflineAPIResponse 将 TaobaoOnebpDkxCreativeCreativeReportOfflineAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOnebpDkxCreativeCreativeReportOfflineAPIResponse(v *TaobaoOnebpDkxCreativeCreativeReportOfflineAPIResponse) {
	v.Reset()
	poolTaobaoOnebpDkxCreativeCreativeReportOfflineAPIResponse.Put(v)
}
