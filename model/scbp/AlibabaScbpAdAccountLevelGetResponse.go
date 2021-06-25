package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询推广账户等级 APIResponse
alibaba.scbp.ad.account.level.get

查询推广账户等级
*/
type AlibabaScbpAdAccountLevelGetAPIResponse struct {
    model.CommonResponse
    Response *AlibabaScbpAdAccountLevelGetResponse `json:"alibaba_scbp_ad_account_level_get_response,omitempty"`
}

type AlibabaScbpAdAccountLevelGetResponse struct {

    // 推广账户等级
    CustLevelDto   *TopCustLevelDto `json:"cust_level_dto,omitempty"`

}
