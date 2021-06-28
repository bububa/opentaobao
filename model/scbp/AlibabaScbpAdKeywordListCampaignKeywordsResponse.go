package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取计划关键词 APIResponse
alibaba.scbp.ad.keyword.list.campaign.keywords

获取计划关键词
*/
type AlibabaScbpAdKeywordListCampaignKeywordsAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_scbp_ad_keyword_list_campaign_keywords_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回数据集合
    
    ResultList   []KeywordDto `json:"result_list,omitempty" xml:"