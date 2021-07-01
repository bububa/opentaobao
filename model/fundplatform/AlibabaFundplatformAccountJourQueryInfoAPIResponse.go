package fundplatform

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询账户流水信息 API返回值 
alibaba.fundplatform.account.jour.query.info

外部查询账户流水信息
*/
type AlibabaFundplatformAccountJourQueryInfoAPIResponse struct {
    model.CommonResponse
    AlibabaFundplatformAccountJourQueryInfoAPIResponseModel
}

// 查询账户流水信息 成功返回结果
type AlibabaFundplatformAccountJourQueryInfoAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_fundplatform_account_jour_query_info_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *ResultSupport `json:"result,omitempty" xml:"result,omitempty"`
}
