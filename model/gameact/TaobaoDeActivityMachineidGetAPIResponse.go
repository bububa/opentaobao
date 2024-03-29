package gameact

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoDeActivityMachineidGetAPIResponse 获取设备号 API返回值
// taobao.de.activity.machineid.get
//
// 获取机器设备id
type TaobaoDeActivityMachineidGetAPIResponse struct {
	model.CommonResponse
	TaobaoDeActivityMachineidGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoDeActivityMachineidGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoDeActivityMachineidGetAPIResponseModel).Reset()
}

// TaobaoDeActivityMachineidGetAPIResponseModel is 获取设备号 成功返回结果
type TaobaoDeActivityMachineidGetAPIResponseModel struct {
	XMLName xml.Name `xml:"de_activity_machineid_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 机器号
	MachineId string `json:"machine_id,omitempty" xml:"machine_id,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoDeActivityMachineidGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MachineId = ""
}

var poolTaobaoDeActivityMachineidGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoDeActivityMachineidGetAPIResponse)
	},
}

// GetTaobaoDeActivityMachineidGetAPIResponse 从 sync.Pool 获取 TaobaoDeActivityMachineidGetAPIResponse
func GetTaobaoDeActivityMachineidGetAPIResponse() *TaobaoDeActivityMachineidGetAPIResponse {
	return poolTaobaoDeActivityMachineidGetAPIResponse.Get().(*TaobaoDeActivityMachineidGetAPIResponse)
}

// ReleaseTaobaoDeActivityMachineidGetAPIResponse 将 TaobaoDeActivityMachineidGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoDeActivityMachineidGetAPIResponse(v *TaobaoDeActivityMachineidGetAPIResponse) {
	v.Reset()
	poolTaobaoDeActivityMachineidGetAPIResponse.Put(v)
}
