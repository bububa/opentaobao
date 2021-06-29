package tmallnr

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
区域零售商品打标去标 APIResponse
tmall.nr.item.tag.ops

参加区域零售的商品，需要批量打标或去标，方便后续设置商品库存
*/
type TmallNrItemTagOpsAPIResponse struct {
    model.CommonResponse
    TmallNrItemTagOpsResponse
}

type TmallNrItemTagOpsResponse struct {
    XMLName xml.Name `xml:"tmall_nr_item_tag_ops_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 返回结果
    
    Result   *NewRetailResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
