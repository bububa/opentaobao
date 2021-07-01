package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaSscSupplyplatformServiceInventoryEditAPIResponse
编辑服务库存 API返回值
alibaba.ssc.supplyplatform.service.inventory.edit

实时编辑服务库存。只支持增加或减少库存，不支持设置绝对库存值。
需要自己处理好幂等逻辑。
要先查询当前库存值，并基于返回结果做编辑操作。
参考alibaba.ssc.supplyplatform.service.inventory.query和alibaba.ssc.supplyplatform.servicecapacity.save */
type AlibabaSscSupplyplatformServiceInventoryEditAPIResponse struct {
	model.CommonResponse
	AlibabaSscSupplyplatformServiceInventoryEditAPIResponseModel
}

// AlibabaSscSupplyplatformServiceInventoryEditAPIResponseModel is 编辑服务库存 成功返回结果
type AlibabaSscSupplyplatformServiceInventoryEditAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ssc_supplyplatform_service_inventory_edit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaSscSupplyplatformServiceInventoryEditResult `json:"result,omitempty" xml:"result,omitempty"`
}
