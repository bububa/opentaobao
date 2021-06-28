package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询推广组 APIResponse
alibaba.scbp.ad.group.find.ad.group

查询推广组
*/
type AlibabaScbpAdGroupFindAdGroupAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaScbpAdGroupFindAdGroupResponse `json:"alibaba_scbp_ad_group_find_ad_group_response,omitempty"` 
    AlibabaScbpAdGroupFindAdGroupResponse
}

/* model for simplify = false
type AlibabaScbpAdGroupFindAdGroupResponse struct {

    // 返回数据
    
    ResultList  struct {
        AdProductDto  []AdProductDto `json:"ad_product_dto,omitempty"`
    } `json:"result_list,omitempty"`
    

}
*/

type AlibabaScbpAdGroupFindAdGroupResponse struct {

    // 返回数据
    ResultList   []AdProductDto `json:"result_list,omitempty"`

}
