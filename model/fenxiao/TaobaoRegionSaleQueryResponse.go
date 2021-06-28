package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询商品销售区域 APIResponse
taobao.region.sale.query

查询商品销售区域
*/
type TaobaoRegionSaleQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"region_sale_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *BaseResult `json:"result,omitempty" xml:"