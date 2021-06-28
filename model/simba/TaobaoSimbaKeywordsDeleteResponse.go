package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
删除一批关键词 APIResponse
taobao.simba.keywords.delete

删除一批关键词
*/
type TaobaoSimbaKeywordsDeleteAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaKeywordsDeleteResponse
}

type TaobaoSimbaKeywordsDeleteResponse struct {
    XMLName xml.Name `xml:"simba_keywords_delete_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 成功删除的关键词列表
    
    Keywords   []Keyword `json:"keywords,omitempty" xml:"keywords>keyword,omitempty"`
    
    
}
