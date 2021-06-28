package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
联营商药品柜核销 APIResponse
wdk.wms.pick.medicine.checksell

联营商药品柜核销
*/
type WdkWmsPickMedicineChecksellAPIResponse struct {
    model.CommonResponse
    // Response *WdkWmsPickMedicineChecksellResponse `json:"wdk_wms_pick_medicine_checksell_response,omitempty"` 
    WdkWmsPickMedicineChecksellResponse
}

/* model for simplify = false
type WdkWmsPickMedicineChecksellResponse struct {

    // result
    
    Result  *struct {
        MedicineResultDto  *MedicineResultDto `json:"medicine_result_dto,omitempty"`
    } `json:"result,omitempty"`
    

}
*/

type WdkWmsPickMedicineChecksellResponse struct {

    // result
    Result   *MedicineResultDto `json:"result,omitempty"`

}
