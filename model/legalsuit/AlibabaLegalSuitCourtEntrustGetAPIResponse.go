package legalsuit

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
委托开庭服务查询 API返回值 
alibaba.legal.suit.court.entrust.get

查询委托开庭信息
*/
type AlibabaLegalSuitCourtEntrustGetAPIResponse struct {
    model.CommonResponse
    AlibabaLegalSuitCourtEntrustGetAPIResponseModel
}

// 委托开庭服务查询 成功返回结果
type AlibabaLegalSuitCourtEntrustGetAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_legal_suit_court_entrust_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // alinkappserver系统返回的通用结果类
    Result   *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}
