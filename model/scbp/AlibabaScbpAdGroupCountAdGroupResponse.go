package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
统计adgroup数量 APIResponse
alibaba.scbp.ad.group.count.ad.group

统计adgroup数量
*/
type AlibabaScbpAdGroupCountAdGroupAPIResponse struct {
    model.CommonResponse
    Response *AlibabaScbpAdGroupCountAdGroupResponse `json:"alibaba_scbp_ad_group_count_ad_group_response,omitempty"`
}

type AlibabaScbpAdGroupCountAdGroupResponse struct {

    // 返回结果
    Result   int64 `json:"result,omitempty"`

}
