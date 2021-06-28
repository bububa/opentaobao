package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询仓库覆盖范围 APIResponse
taobao.region.warehouse.query

查询仓库覆盖范围
*/
type TaobaoRegionWarehouseQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"region_warehouse_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *BaseResult `json:"result,omitempty" xml:"