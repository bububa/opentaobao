package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据一个关键词Id列表取得一组关键词 APIResponse
taobao.simba.keywordsbykeywordids.get

根据一个关键词Id列表取得一组关键词
*/
type TaobaoSimbaKeywordsbykeywordidsGetAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaKeywordsbykeywordidsGetResponse
}

type TaobaoSimbaKeywordsbykeywordidsGetResponse struct {
    XMLName xml.Name `xml:"simba_keywordsbykeywordids_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 取得的关键词列表
    
    Keywords   []Keyword `json:"keywords,omitempty" xml:"keywords>keyword,omitempty"`
    
    
}
