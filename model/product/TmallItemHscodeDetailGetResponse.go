package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
通过hscode获取计量单位 APIResponse
tmall.item.hscode.detail.get

通过hscode获取计量单位和销售单位
*/
type TmallItemHscodeDetailGetAPIResponse struct {
    model.CommonResponse
    TmallItemHscodeDetailGetResponse
}

type TmallItemHscodeDetailGetResponse struct {
    XMLName xml.Name `xml:"tmall_item_hscode_detail_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回的计量单位和销售单位
    
    Results   []string `json:"results,omitempty" xml:"results>string,omitempty"`
    
    
}
