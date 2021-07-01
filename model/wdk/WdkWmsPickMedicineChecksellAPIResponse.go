package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* WdkWmsPickMedicineChecksellAPIResponse
联营商药品柜核销 API返回值
wdk.wms.pick.medicine.checksell

联营商药品柜核销 */
type WdkWmsPickMedicineChecksellAPIResponse struct {
	model.CommonResponse
	WdkWmsPickMedicineChecksellAPIResponseModel
}

// WdkWmsPickMedicineChecksellAPIResponseModel is 联营商药品柜核销 成功返回结果
type WdkWmsPickMedicineChecksellAPIResponseModel struct {
	XMLName xml.Name `xml:"wdk_wms_pick_medicine_checksell_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *MedicineResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
