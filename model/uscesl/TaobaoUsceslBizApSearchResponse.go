package uscesl

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
AP列表查询 APIResponse
taobao.uscesl.biz.ap.search

查询当前门店下登记的AP列表
*/
type TaobaoUsceslBizApSearchAPIResponse struct {
    model.CommonResponse
    TaobaoUsceslBizApSearchResponse
}

type TaobaoUsceslBizApSearchResponse struct {
    XMLName xml.Name `xml:"uscesl_biz_ap_search_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 成功与否看result.success，返回true或者false
    
    Result   *TaobaoUsceslBizApSearchResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
