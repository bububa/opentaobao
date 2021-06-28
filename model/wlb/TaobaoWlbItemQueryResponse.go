package wlb

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
分页查询商品 APIResponse
taobao.wlb.item.query

根据状态、卖家、SKU等信息查询商品列表
*/
type TaobaoWlbItemQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"wlb_item_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果总数
    
    TotalCount   int64 `json:"total_count,omitempty" xml:"