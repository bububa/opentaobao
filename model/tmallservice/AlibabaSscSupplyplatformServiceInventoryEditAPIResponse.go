package tmallservice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaSscSupplyplatformServiceInventoryEditAPIResponse 编辑服务库存 API返回值
// alibaba.ssc.supplyplatform.service.inventory.edit
//
// 实时编辑服务库存。只支持增加或减少库存，不支持设置绝对库存值。
// 需要自己处理好幂等逻辑。
// 要先查询当前库存值，并基于返回结果做编辑操作。
// 参考alibaba.ssc.supplyplatform.service.inventory.query和alibaba.ssc.supplyplatform.servicecapacity.save
type AlibabaSscSupplyplatformServiceInventoryEditAPIResponse struct {
	model.CommonResponse
	AlibabaSscSupplyplatformServiceInventoryEditAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaSscSupplyplatformServiceInventoryEditAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaSscSupplyplatformServiceInventoryEditAPIResponseModel).Reset()
}

// AlibabaSscSupplyplatformServiceInventoryEditAPIResponseModel is 编辑服务库存 成功返回结果
type AlibabaSscSupplyplatformServiceInventoryEditAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ssc_supplyplatform_service_inventory_edit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaSscSupplyplatformServiceInventoryEditResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaSscSupplyplatformServiceInventoryEditAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaSscSupplyplatformServiceInventoryEditAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaSscSupplyplatformServiceInventoryEditAPIResponse)
	},
}

// GetAlibabaSscSupplyplatformServiceInventoryEditAPIResponse 从 sync.Pool 获取 AlibabaSscSupplyplatformServiceInventoryEditAPIResponse
func GetAlibabaSscSupplyplatformServiceInventoryEditAPIResponse() *AlibabaSscSupplyplatformServiceInventoryEditAPIResponse {
	return poolAlibabaSscSupplyplatformServiceInventoryEditAPIResponse.Get().(*AlibabaSscSupplyplatformServiceInventoryEditAPIResponse)
}

// ReleaseAlibabaSscSupplyplatformServiceInventoryEditAPIResponse 将 AlibabaSscSupplyplatformServiceInventoryEditAPIResponse 保存到 sync.Pool
func ReleaseAlibabaSscSupplyplatformServiceInventoryEditAPIResponse(v *AlibabaSscSupplyplatformServiceInventoryEditAPIResponse) {
	v.Reset()
	poolAlibabaSscSupplyplatformServiceInventoryEditAPIResponse.Put(v)
}
