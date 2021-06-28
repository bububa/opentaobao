package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
系统推荐 APIResponse
alibaba.scbp.reckeyword.sys.get

查询系统推荐词
*/
type AlibabaScbpReckeywordSysGetAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaScbpReckeywordSysGetResponse `json:"alibaba_scbp_reckeyword_sys_get_response,omitempty"` 
    AlibabaScbpReckeywordSysGetResponse
}

/* model for simplify = false
type AlibabaScbpReckeywordSysGetResponse struct {

    // 总个数
    
    TotalNum   int64 `json:"total_num,omitempty"`
    

    // 总页数
    
    TotalPage   int64 `json:"total_page,omitempty"`
    

    // 系统推荐词结果列表
    
    ResultList  struct {
        RecKeywordDto  []RecKeywordDto `json:"rec_keyword_dto,omitempty"`
    } `json:"result_list,omitempty"`
    

}
*/

type AlibabaScbpReckeywordSysGetResponse struct {

    // 总个数
    TotalNum   int64 `json:"total_num,omitempty"`

    // 总页数
    TotalPage   int64 `json:"total_page,omitempty"`

    // 系统推荐词结果列表
    ResultList   []RecKeywordDto `json:"result_list,omitempty"`

}
