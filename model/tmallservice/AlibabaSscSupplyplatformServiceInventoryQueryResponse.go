package tmallservice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
服务库存查询 API返回值 
alibaba.ssc.supplyplatform.service.inventory.query

查询服务库存。需要保存服务容量成功后，才能进行查询，参数中的provider信息（provider_id和provider_type）与alibaba.ssc.supplyplatform.servicecapacity.save接口中保持一致。
*/
type AlibabaSscSupplyplatformServiceInventoryQueryAPIResponse struct {
    model.CommonResponse
    AlibabaSscSupplyplatformServiceInventoryQueryResponse
}

// 服务库存查询 成功返回结果
type AlibabaSscSupplyplatformServiceInventoryQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_ssc_supplyplatform_service_inventory_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 接口返回model
    Result   *AlibabaSscSupplyplatformServiceInventoryQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
