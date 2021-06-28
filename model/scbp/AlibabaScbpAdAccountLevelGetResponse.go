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
    // Response *AlibabaScbpAdAccountLevelGetResponse `json:"alibaba_scbp_ad_account_level_get_response,omitempty"` 
    AlibabaScbpAdAccountLevelGetResponse
}

/* model for simplify = false
type AlibabaScbpAdAccountLevelGetResponse struct {

    // 推广账户等级
    
    CustLevelDto  *struct {
        TopCustLevelDto  *TopCustLevelDto `json:"top_cust_level_dto,omitempty"`
    } `json:"cust_level_dto,omitempty"`
    

}
*/

type AlibabaScbpAdAccountLevelGetResponse struct {

    // 推广账户等级
    CustLevelDto   *TopCustLevelDto `json:"cust_level_dto,omitempty"`

}
