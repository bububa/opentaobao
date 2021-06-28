package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
编辑仓库覆盖范围 APIResponse
taobao.region.warehouse.manage

编辑仓库覆盖范围
*/
type TaobaoRegionWarehouseManageAPIResponse struct {
    model.CommonResponse
    TaobaoRegionWarehouseManageResponse
}

type TaobaoRegionWarehouseManageResponse struct {
    XMLName xml.Name `xml:"region_warehouse_manage_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果
    
    Result   *BaseResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
