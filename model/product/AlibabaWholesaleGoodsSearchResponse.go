package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批发市场产品搜索 APIResponse
alibaba.wholesale.goods.search

批发市场产品搜索
*/
type AlibabaWholesaleGoodsSearchAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wholesale_goods_search_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 在线批发商品搜索结果
    
    WholesaleGoodsSearchResult   *WholesaleSearchOpenResult `json:"wholesale_goods_search_result,omitempty" xml:"