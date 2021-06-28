package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
联营商药品柜核销 APIResponse
wdk.wms.pick.medicine.checksell

联营商药品柜核销
*/
type WdkWmsPickMedicineChecksellAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"wdk_wms_pick_medicine_checksell_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   *MedicineResultDto `json:"result,omitempty" xml:"