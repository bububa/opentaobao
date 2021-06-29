package wenyuvideo

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
根据人物名称查询人物列表 APIResponse
youku.wenyuvideo.persion.search

根据人物名称查询人物列表
*/
type YoukuWenyuvideoPersionSearchAPIResponse struct {
    model.CommonResponse
    YoukuWenyuvideoPersionSearchResponse
}

type YoukuWenyuvideoPersionSearchResponse struct {
    XMLName xml.Name `xml:"youku_wenyuvideo_persion_search_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 接口返回model
    
    Result   *YoukuWenyuvideoPersionSearchResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
