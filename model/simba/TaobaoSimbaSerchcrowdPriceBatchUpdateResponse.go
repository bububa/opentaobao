package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
单品推广搜索人群修改溢价 APIResponse
taobao.simba.serchcrowd.price.batch.update

单品推广搜索人群修改溢价, 不支持跨推广单元修改
*/
type TaobaoSimbaSerchcrowdPriceBatchUpdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"simba_serchcrowd_price_batch_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Adgrouptargetingtags   []AdgroupTargetingTagDto `json:"adgrouptargetingtags,omitempty" xml:"