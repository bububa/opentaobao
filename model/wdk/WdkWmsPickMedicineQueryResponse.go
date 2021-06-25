package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询拣货单中的药品信息 APIResponse
wdk.wms.pick.medicine.query

联营商药机查询拣货单中的药品信息
*/
type WdkWmsPickMedicineQueryAPIResponse struct {
    model.CommonResponse
    Response *WdkWmsPickMedicineQueryResponse `json:"wdk_wms_pick_medicine_query_response,omitempty"`
}

type WdkWmsPickMedicineQueryResponse struct {

    // 接口返回model
    Result   *WdkWmsPickMedicineQueryResult `json:"result,omitempty"`

}
