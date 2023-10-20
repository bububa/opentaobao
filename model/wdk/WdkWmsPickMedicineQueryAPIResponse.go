package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// WdkWmsPickMedicineQueryAPIResponse 查询拣货单中的药品信息 API返回值
// wdk.wms.pick.medicine.query
//
// 联营商药机查询拣货单中的药品信息
type WdkWmsPickMedicineQueryAPIResponse struct {
	model.CommonResponse
	WdkWmsPickMedicineQueryAPIResponseModel
}

// Reset 清空结构体
func (m *WdkWmsPickMedicineQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.WdkWmsPickMedicineQueryAPIResponseModel).Reset()
}

// WdkWmsPickMedicineQueryAPIResponseModel is 查询拣货单中的药品信息 成功返回结果
type WdkWmsPickMedicineQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"wdk_wms_pick_medicine_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *WdkWmsPickMedicineQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *WdkWmsPickMedicineQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolWdkWmsPickMedicineQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(WdkWmsPickMedicineQueryAPIResponse)
	},
}

// GetWdkWmsPickMedicineQueryAPIResponse 从 sync.Pool 获取 WdkWmsPickMedicineQueryAPIResponse
func GetWdkWmsPickMedicineQueryAPIResponse() *WdkWmsPickMedicineQueryAPIResponse {
	return poolWdkWmsPickMedicineQueryAPIResponse.Get().(*WdkWmsPickMedicineQueryAPIResponse)
}

// ReleaseWdkWmsPickMedicineQueryAPIResponse 将 WdkWmsPickMedicineQueryAPIResponse 保存到 sync.Pool
func ReleaseWdkWmsPickMedicineQueryAPIResponse(v *WdkWmsPickMedicineQueryAPIResponse) {
	v.Reset()
	poolWdkWmsPickMedicineQueryAPIResponse.Put(v)
}
