package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
服务库存查询 APIResponse
alibaba.ssc.supplyplatform.service.inventory.query

查询服务库存。需要保存服务容量成功后，才能进行查询，参数中的provider信息（provider_id和provider_type）与alibaba.ssc.supplyplatform.servicecapacity.save接口中保持一致。
*/
type AlibabaSscSupplyplatformServiceInventoryQueryAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaSscSupplyplatformServiceInventoryQueryResponse `json:"alibaba_ssc_supplyplatform_service_inventory_query_response,omitempty"` 
    AlibabaSscSupplyplatformServiceInventoryQueryResponse
}

/* model for simplify = false
type AlibabaSscSupplyplatformServiceInventoryQueryResponse struct {

    // 接口返回model
    
    Result  *struct {
        AlibabaSscSupplyplatformServiceInventoryQueryResult  *AlibabaSscSupplyplatformServiceInventoryQueryResult `json:"alibaba_ssc_supplyplatform_service_inventory_query_result,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type AlibabaSscSupplyplatformServiceInventoryQueryResponse struct {

    // 接口返回model
    Result   *AlibabaSscSupplyplatformServiceInventoryQueryResult `json:"result,omitempty"`

}
