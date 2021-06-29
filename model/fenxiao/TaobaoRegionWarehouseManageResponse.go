package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
编辑仓库覆盖范围 API返回值 
taobao.region.warehouse.manage

编辑仓库覆盖范围
*/
type TaobaoRegionWarehouseManageAPIResponse struct {
    model.CommonResponse
    TaobaoRegionWarehouseManageResponse
}

// 编辑仓库覆盖范围 成功返回结果
type TaobaoRegionWarehouseManageResponse struct {
    XMLName xml.Name `xml:"region_warehouse_manage_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 返回结果
    Result   *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}
