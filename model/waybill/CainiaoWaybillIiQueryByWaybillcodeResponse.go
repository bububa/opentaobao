package waybill

import (
    "github.com/bububa/opentaobao/model"
)

/* 
通过面单号查询面单信息 APIResponse
cainiao.waybill.ii.query.by.waybillcode

通过面单号查看面单号的当前状态，如签收、发货、失效等。
*/
type CainiaoWaybillIiQueryByWaybillcodeAPIResponse struct {
    model.CommonResponse
    // Response *CainiaoWaybillIiQueryByWaybillcodeResponse `json:"cainiao_waybill_ii_query_by_waybillcode_response,omitempty"` 
    CainiaoWaybillIiQueryByWaybillcodeResponse
}

/* model for simplify = false
type CainiaoWaybillIiQueryByWaybillcodeResponse struct {

    // 查询返回值
    
    Modules  struct {
        WaybillCloudPrintWithResultDescResponse  []WaybillCloudPrintWithResultDescResponse `json:"waybill_cloud_print_with_result_desc_response,omitempty"`
    } `json:"modules,omitempty"`
    

}
*/

type CainiaoWaybillIiQueryByWaybillcodeResponse struct {

    // 查询返回值
    Modules   []WaybillCloudPrintWithResultDescResponse `json:"modules,omitempty"`

}
