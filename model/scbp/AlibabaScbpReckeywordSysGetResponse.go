package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
系统推荐 APIResponse
alibaba.scbp.reckeyword.sys.get

查询系统推荐词
*/
type AlibabaScbpReckeywordSysGetAPIResponse struct {
    model.CommonResponse
    AlibabaScbpReckeywordSysGetResponse
}

type AlibabaScbpReckeywordSysGetResponse struct {
    XMLName xml.Name `xml:"alibaba_scbp_reckeyword_sys_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 总个数
    
    TotalNum   int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`

    
    // 总页数
    
    TotalPage   int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`

    
    // 系统推荐词结果列表
    
    ResultList   []RecKeywordDto `json:"result_list,omitempty" xml:"result_list>rec_keyword_dto,omitempty"`
    
    
}
