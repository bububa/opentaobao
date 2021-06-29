package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
关键词报表 API返回值 
alibaba.scbp.effect.keyword.list

关键词报表
*/
type AlibabaScbpEffectKeywordListAPIResponse struct {
    model.CommonResponse
    AlibabaScbpEffectKeywordListResponse
}

// 关键词报表 成功返回结果
type AlibabaScbpEffectKeywordListResponse struct {
    XMLName xml.Name `xml:"alibaba_scbp_effect_keyword_list_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 关键词效果数据列表
    KeywordReportList   []AdKeywordEffectDTO `json:"keyword_report_list,omitempty" xml:"keyword_report_list>ad_keyword_effect_dto,omitempty"`
    // 总个数
    TotalNum   int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
    // 总页数
    TotalPage   int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`
}
