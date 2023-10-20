package tmallcar

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCarVehicleinfoRegisterAPIResponse 全量车型导入 API返回值
// taobao.car.vehicleinfo.register
//
// 全量车型导入
type TaobaoCarVehicleinfoRegisterAPIResponse struct {
	model.CommonResponse
	TaobaoCarVehicleinfoRegisterAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoCarVehicleinfoRegisterAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoCarVehicleinfoRegisterAPIResponseModel).Reset()
}

// TaobaoCarVehicleinfoRegisterAPIResponseModel is 全量车型导入 成功返回结果
type TaobaoCarVehicleinfoRegisterAPIResponseModel struct {
	XMLName xml.Name `xml:"car_vehicleinfo_register_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 错误码
	ResultCode int64 `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 返回素材id
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 是否成功
	ResultSuccess bool `json:"result_success,omitempty" xml:"result_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoCarVehicleinfoRegisterAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ResultMsg = ""
	m.ResultCode = 0
	m.Data = false
	m.ResultSuccess = false
}

var poolTaobaoCarVehicleinfoRegisterAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoCarVehicleinfoRegisterAPIResponse)
	},
}

// GetTaobaoCarVehicleinfoRegisterAPIResponse 从 sync.Pool 获取 TaobaoCarVehicleinfoRegisterAPIResponse
func GetTaobaoCarVehicleinfoRegisterAPIResponse() *TaobaoCarVehicleinfoRegisterAPIResponse {
	return poolTaobaoCarVehicleinfoRegisterAPIResponse.Get().(*TaobaoCarVehicleinfoRegisterAPIResponse)
}

// ReleaseTaobaoCarVehicleinfoRegisterAPIResponse 将 TaobaoCarVehicleinfoRegisterAPIResponse 保存到 sync.Pool
func ReleaseTaobaoCarVehicleinfoRegisterAPIResponse(v *TaobaoCarVehicleinfoRegisterAPIResponse) {
	v.Reset()
	poolTaobaoCarVehicleinfoRegisterAPIResponse.Put(v)
}
