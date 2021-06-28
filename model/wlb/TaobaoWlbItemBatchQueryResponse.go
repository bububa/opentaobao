package wlb

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
批次库存查询接口 APIResponse
taobao.wlb.item.batch.query

根据用户id，item id list和store code来查询商品库存信息和批次信息
*/
type TaobaoWlbItemBatchQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"wlb_item_batch_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果记录的数量
    
    TotalCount   int64 `json:"total_count,omitempty" xml:"