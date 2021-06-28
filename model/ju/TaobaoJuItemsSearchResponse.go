package ju

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
聚划算商品搜索接口 APIResponse
taobao.ju.items.search

搜索聚划算商品
*/
type TaobaoJuItemsSearchAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"ju_items_search_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果
    
    Result   *PaginationResult `json:"result,omitempty" xml:"