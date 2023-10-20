package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallservicesettlementbillinfoqueryAPIResponse 查询工单结算信息 API返回值
// tmall.service.settlement.billinfo.query
//
// 提供给服务商查询工单结算信息，包含结算的分成金额以及结算的收款明细，平台抽佣比例。用于服务商进行对账
type TmallservicesettlementbillinfoqueryAPIResponse struct {
	model.CommonResponse
	TmallservicesettlementbillinfoqueryAPIResponseModel
}

// TmallservicesettlementbillinfoqueryAPIResponseModel is 查询工单结算信息 成功返回结果
type TmallservicesettlementbillinfoqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_service_settlement_billinfo_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口响应
	Result *TmallservicesettlementbillinfoqueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
