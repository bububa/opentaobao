package waybill

import (
    "github.com/bububa/opentaobao/model"
)

/* 
通过订单号查询电子面单通接口 APIResponse
cainiao.waybill.ii.query.by.tradecode

通过订单号查看面单的信息
*/
type CainiaoWaybillIiQueryByTradecodeAPIResponse struct {
    model.CommonResponse
    // Response *CainiaoWaybillIiQueryByTradecodeResponse `json:"cainiao_waybill_ii_query_by_tradecode_response,omitempty"` 
    CainiaoWaybillIiQueryByTradecodeResponse
}

/* model for simplify = false
type CainiaoWaybillIiQueryByTradecodeResponse struct {

    // 查询返回值
    
    Modules  struct {
        WaybillCloudPrintWithResultDescResponse  []WaybillCloudPrintWithResultDescResponse `json:"waybill_cloud_print_with_result_desc_response,omitempty"`
    } `json:"modules,omitempty"`
    

}
*/

type CainiaoWaybillIiQueryByTradecodeResponse struct {

    // 查询返回值
    Modules   []WaybillCloudPrintWithResultDescResponse `json:"modules,omitempty"`

}
