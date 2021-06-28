package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询推广账户等级 APIResponse
alibaba.scbp.ad.account.level.get

查询推广账户等级
*/
type AlibabaScbpAdAccountLevelGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_scbp_ad_account_level_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 推广账户等级
    
    CustLevelDto   *TopCustLevelDto `json:"cust_level_dto,omitempty" xml:"