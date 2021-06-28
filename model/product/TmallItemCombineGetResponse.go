package product

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
组合商品获取接口 APIResponse
tmall.item.combine.get

查询组合商品的SKU信息
*/
type TmallItemCombineGetAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"tmall_item_combine_get_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // results
    
    Results   []string `json:"results,omitempty" xml:"