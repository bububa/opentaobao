package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
抽奖平台奖池查询接口 APIResponse
alibaba.marketing.lottery.activity.query

抽奖平台奖池查询接口
*/
type AlibabaMarketingLotteryActivityQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_marketing_lottery_activity_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 分页结果
    
    PagingDto   *PagingDto `json:"paging_dto,omitempty" xml:"