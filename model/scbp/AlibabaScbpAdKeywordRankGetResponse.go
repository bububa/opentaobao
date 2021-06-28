package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取外贸直通车关键词预估排名 APIResponse
alibaba.scbp.ad.keyword.rank.get

获取外贸直通车关键词预估排名
*/
type AlibabaScbpAdKeywordRankGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_scbp_ad_keyword_rank_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 关键词的预估排名
    
    RankLocation   int64 `json:"rank_location,omitempty" xml:"