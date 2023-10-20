package antifraud

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAntifraudRiskuserGetAPIResponse 反欺诈用户风险查询 API返回值
// taobao.antifraud.riskuser.get
//
// 根据用户基础信息，核实平台上的用户是否存在欺诈风险
type TaobaoAntifraudRiskuserGetAPIResponse struct {
	model.CommonResponse
	TaobaoAntifraudRiskuserGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAntifraudRiskuserGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAntifraudRiskuserGetAPIResponseModel).Reset()
}

// TaobaoAntifraudRiskuserGetAPIResponseModel is 反欺诈用户风险查询 成功返回结果
type TaobaoAntifraudRiskuserGetAPIResponseModel struct {
	XMLName xml.Name `xml:"antifraud_riskuser_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 风险结果详情列表，包含多个风险结果单项
	DetailList []AccountRiskDetail `json:"detail_list,omitempty" xml:"detail_list>account_risk_detail,omitempty"`
	// 服务调用成功时, 返回的系统流水号
	EventId string `json:"event_id,omitempty" xml:"event_id,omitempty"`
	// 风险分值
	Score string `json:"score,omitempty" xml:"score,omitempty"`
	// 风险结果, 为reject, review, pass三者之一
	RiskLevel string `json:"risk_level,omitempty" xml:"risk_level,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAntifraudRiskuserGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.DetailList = m.DetailList[:0]
	m.EventId = ""
	m.Score = ""
	m.RiskLevel = ""
}

var poolTaobaoAntifraudRiskuserGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAntifraudRiskuserGetAPIResponse)
	},
}

// GetTaobaoAntifraudRiskuserGetAPIResponse 从 sync.Pool 获取 TaobaoAntifraudRiskuserGetAPIResponse
func GetTaobaoAntifraudRiskuserGetAPIResponse() *TaobaoAntifraudRiskuserGetAPIResponse {
	return poolTaobaoAntifraudRiskuserGetAPIResponse.Get().(*TaobaoAntifraudRiskuserGetAPIResponse)
}

// ReleaseTaobaoAntifraudRiskuserGetAPIResponse 将 TaobaoAntifraudRiskuserGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAntifraudRiskuserGetAPIResponse(v *TaobaoAntifraudRiskuserGetAPIResponse) {
	v.Reset()
	poolTaobaoAntifraudRiskuserGetAPIResponse.Put(v)
}
