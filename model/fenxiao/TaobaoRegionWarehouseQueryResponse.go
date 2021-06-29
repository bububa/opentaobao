package fenxiao

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询仓库覆盖范围 API返回值 
taobao.region.warehouse.query

查询仓库覆盖范围
*/
type TaobaoRegionWarehouseQueryAPIResponse struct {
    model.CommonResponse
    TaobaoRegionWarehouseQueryResponse
}

// 查询仓库覆盖范围 成功返回结果
type TaobaoRegionWarehouseQueryResponse struct {
    XMLName xml.Name `xml:"region_warehouse_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}
