package shenjing

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
IB智慧园区-查询TMI流水 API返回值 
alibaba.ib.campus.tmi.query

获取特定银行账户的银行流水
*/
type AlibabaIbCampusTmiQueryAPIResponse struct {
    model.CommonResponse
    AlibabaIbCampusTmiQueryResponse
}

// IB智慧园区-查询TMI流水 成功返回结果
type AlibabaIbCampusTmiQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_ib_campus_tmi_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // TMI流水
    Results   []TradeRecordDTO `json:"results,omitempty" xml:"results>trade_record_dto,omitempty"`
}
