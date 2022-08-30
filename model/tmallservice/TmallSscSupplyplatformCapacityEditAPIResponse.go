package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallSscSupplyplatformCapacityEditAPIResponse 容量编辑 API返回值
// tmall.ssc.supplyplatform.capacity.edit
//
// 容量编辑
type TmallSscSupplyplatformCapacityEditAPIResponse struct {
	model.CommonResponse
	TmallSscSupplyplatformCapacityEditAPIResponseModel
}

// TmallSscSupplyplatformCapacityEditAPIResponseModel is 容量编辑 成功返回结果
type TmallSscSupplyplatformCapacityEditAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_ssc_supplyplatform_capacity_edit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回类型
	Result *TmallSscSupplyplatformCapacityEditResult `json:"result,omitempty" xml:"result,omitempty"`
}
