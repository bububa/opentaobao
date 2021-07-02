package mos

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMjPresaleSettlementStatisticsAPIResponse 预购结算数据统计 API返回值
// alibaba.mj.presale.settlement.statistics
//
// 预购结算数据统计
type AlibabaMjPresaleSettlementStatisticsAPIResponse struct {
	model.CommonResponse
	AlibabaMjPresaleSettlementStatisticsAPIResponseModel
}

// AlibabaMjPresaleSettlementStatisticsAPIResponseModel is 预购结算数据统计 成功返回结果
type AlibabaMjPresaleSettlementStatisticsAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mj_presale_settlement_statistics_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// data
	Data *AlibabaMjPresaleSettlementStatisticsData `json:"data,omitempty" xml:"data,omitempty"`
}
