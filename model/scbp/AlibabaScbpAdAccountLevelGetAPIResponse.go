package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询推广账户等级 API返回值 
alibaba.scbp.ad.account.level.get

查询推广账户等级
*/
type AlibabaScbpAdAccountLevelGetAPIResponse struct {
    model.CommonResponse
    AlibabaScbpAdAccountLevelGetAPIResponseModel
}

// 查询推广账户等级 成功返回结果
type AlibabaScbpAdAccountLevelGetAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_scbp_ad_account_level_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 推广账户等级
    CustLevelDto   *TopCustLevelDto `json:"cust_level_dto,omitempty" xml:"cust_level_dto,omitempty"`
}
