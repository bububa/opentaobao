package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
分页获取修改过的关键词ID、宝贝id、修改时间 APIResponse
taobao.simba.keywords.changed.get

分页获取修改过的关键词ID、宝贝id、修改时间
*/
type TaobaoSimbaKeywordsChangedGetAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaKeywordsChangedGetResponse
}

type TaobaoSimbaKeywordsChangedGetResponse struct {
    XMLName xml.Name `xml:"simba_keywords_changed_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 关键词分页对象
    
    Keywords   *KeywordPage `json:"keywords,omitempty" xml:"keywords,omitempty"`

    
}
