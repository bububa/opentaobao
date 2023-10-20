package mos

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabamjpresalesettlementstatisticsAPIResponse 预购结算数据统计 API返回值
// alibaba.mj.presale.settlement.statistics
//
// 预购结算数据统计
type AlibabamjpresalesettlementstatisticsAPIResponse struct {
	model.CommonResponse
	AlibabamjpresalesettlementstatisticsAPIResponseModel
}

// AlibabamjpresalesettlementstatisticsAPIResponseModel is 预购结算数据统计 成功返回结果
type AlibabamjpresalesettlementstatisticsAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mj_presale_settlement_statistics_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// data
	Data *AlibabamjpresalesettlementstatisticsData `json:"data,omitempty" xml:"data,omitempty"`
}
