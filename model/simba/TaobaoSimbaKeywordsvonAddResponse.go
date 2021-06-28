package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
创建一批关键词 APIResponse
taobao.simba.keywordsvon.add

创建一批关键词
*/
type TaobaoSimbaKeywordsvonAddAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaKeywordsvonAddResponse
}

type TaobaoSimbaKeywordsvonAddResponse struct {
    XMLName xml.Name `xml:"simba_keywordsvon_add_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 关键词列表
    
    Keywords   []Keyword `json:"keywords,omitempty" xml:"keywords>keyword,omitempty"`
    
    
}
