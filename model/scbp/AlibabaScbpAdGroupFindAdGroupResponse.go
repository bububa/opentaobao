package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询推广组 APIResponse
alibaba.scbp.ad.group.find.ad.group

查询推广组
*/
type AlibabaScbpAdGroupFindAdGroupAPIResponse struct {
    model.CommonResponse
    AlibabaScbpAdGroupFindAdGroupResponse
}

type AlibabaScbpAdGroupFindAdGroupResponse struct {
    XMLName xml.Name `xml:"alibaba_scbp_ad_group_find_ad_group_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回数据
    
    ResultList   []AdProductDto `json:"result_list,omitempty" xml:"result_list>ad_product_dto,omitempty"`
    
    
}
