package idle

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidleconsignmentspustatisticsAPIResponse 闲鱼帮卖同步服务商交易统计信息 API返回值
// alibaba.idle.consignment.spu.statistics
//
// 闲鱼帮卖同步服务商交易统计信息
type AlibabaidleconsignmentspustatisticsAPIResponse struct {
	model.CommonResponse
	AlibabaidleconsignmentspustatisticsAPIResponseModel
}

// AlibabaidleconsignmentspustatisticsAPIResponseModel is 闲鱼帮卖同步服务商交易统计信息 成功返回结果
type AlibabaidleconsignmentspustatisticsAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_consignment_spu_statistics_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}
