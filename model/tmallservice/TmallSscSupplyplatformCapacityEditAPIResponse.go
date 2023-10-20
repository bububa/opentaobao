package tmallservice

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *TmallSscSupplyplatformCapacityEditAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallSscSupplyplatformCapacityEditAPIResponseModel).Reset()
}

// TmallSscSupplyplatformCapacityEditAPIResponseModel is 容量编辑 成功返回结果
type TmallSscSupplyplatformCapacityEditAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_ssc_supplyplatform_capacity_edit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回类型
	Result *TmallSscSupplyplatformCapacityEditResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallSscSupplyplatformCapacityEditAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallSscSupplyplatformCapacityEditAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallSscSupplyplatformCapacityEditAPIResponse)
	},
}

// GetTmallSscSupplyplatformCapacityEditAPIResponse 从 sync.Pool 获取 TmallSscSupplyplatformCapacityEditAPIResponse
func GetTmallSscSupplyplatformCapacityEditAPIResponse() *TmallSscSupplyplatformCapacityEditAPIResponse {
	return poolTmallSscSupplyplatformCapacityEditAPIResponse.Get().(*TmallSscSupplyplatformCapacityEditAPIResponse)
}

// ReleaseTmallSscSupplyplatformCapacityEditAPIResponse 将 TmallSscSupplyplatformCapacityEditAPIResponse 保存到 sync.Pool
func ReleaseTmallSscSupplyplatformCapacityEditAPIResponse(v *TmallSscSupplyplatformCapacityEditAPIResponse) {
	v.Reset()
	poolTmallSscSupplyplatformCapacityEditAPIResponse.Put(v)
}
