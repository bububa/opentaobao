package icbu

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商品批量上下架接口 APIResponse
alibaba.icbu.product.batch.update.display

给国际站的三方服务商提供批量上下架接口
*/
type AlibabaIcbuProductBatchUpdateDisplayAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_icbu_product_batch_update_display_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 只有出错才会显示，唯一标识这次请求
    
    TraceId   string `json:"trace_id,omitempty" xml:"