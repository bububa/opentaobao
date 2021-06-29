package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
取得一个推广组的所有关键词 APIResponse
taobao.simba.keywordsbyadgroupid.get

取得一个推广组的所有关键词
*/
type TaobaoSimbaKeywordsbyadgroupidGetAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaKeywordsbyadgroupidGetResponse
}

type TaobaoSimbaKeywordsbyadgroupidGetResponse struct {
    XMLName xml.Name `xml:"simba_keywordsbyadgroupid_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 取得的关键词列表
    
    Keywords   []Keyword `json:"keywords,omitempty" xml:"keywords>keyword,omitempty"`
    
    
}
