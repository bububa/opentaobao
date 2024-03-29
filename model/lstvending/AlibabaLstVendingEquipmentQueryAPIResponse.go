package lstvending

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstVendingEquipmentQueryAPIResponse 自动售卖机设备信息查询 API返回值
// alibaba.lst.vending.equipment.query
//
// 为零售通品牌商提供已租赁的设备信息查询。
type AlibabaLstVendingEquipmentQueryAPIResponse struct {
	model.CommonResponse
	AlibabaLstVendingEquipmentQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLstVendingEquipmentQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLstVendingEquipmentQueryAPIResponseModel).Reset()
}

// AlibabaLstVendingEquipmentQueryAPIResponseModel is 自动售卖机设备信息查询 成功返回结果
type AlibabaLstVendingEquipmentQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_vending_equipment_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果集
	Result *AlibabaLstVendingEquipmentQueryResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLstVendingEquipmentQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaLstVendingEquipmentQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLstVendingEquipmentQueryAPIResponse)
	},
}

// GetAlibabaLstVendingEquipmentQueryAPIResponse 从 sync.Pool 获取 AlibabaLstVendingEquipmentQueryAPIResponse
func GetAlibabaLstVendingEquipmentQueryAPIResponse() *AlibabaLstVendingEquipmentQueryAPIResponse {
	return poolAlibabaLstVendingEquipmentQueryAPIResponse.Get().(*AlibabaLstVendingEquipmentQueryAPIResponse)
}

// ReleaseAlibabaLstVendingEquipmentQueryAPIResponse 将 AlibabaLstVendingEquipmentQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLstVendingEquipmentQueryAPIResponse(v *AlibabaLstVendingEquipmentQueryAPIResponse) {
	v.Reset()
	poolAlibabaLstVendingEquipmentQueryAPIResponse.Put(v)
}
