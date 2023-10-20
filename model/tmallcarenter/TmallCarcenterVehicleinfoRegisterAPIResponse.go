package tmallcarenter

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallCarcenterVehicleinfoRegisterAPIResponse 车型数据更新 API返回值
// tmall.carcenter.vehicleinfo.register
//
// 基本车型信息维护
type TmallCarcenterVehicleinfoRegisterAPIResponse struct {
	model.CommonResponse
	TmallCarcenterVehicleinfoRegisterAPIResponseModel
}

// Reset 清空结构体
func (m *TmallCarcenterVehicleinfoRegisterAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallCarcenterVehicleinfoRegisterAPIResponseModel).Reset()
}

// TmallCarcenterVehicleinfoRegisterAPIResponseModel is 车型数据更新 成功返回结果
type TmallCarcenterVehicleinfoRegisterAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_carcenter_vehicleinfo_register_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TmallCarcenterVehicleinfoRegisterResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallCarcenterVehicleinfoRegisterAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallCarcenterVehicleinfoRegisterAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallCarcenterVehicleinfoRegisterAPIResponse)
	},
}

// GetTmallCarcenterVehicleinfoRegisterAPIResponse 从 sync.Pool 获取 TmallCarcenterVehicleinfoRegisterAPIResponse
func GetTmallCarcenterVehicleinfoRegisterAPIResponse() *TmallCarcenterVehicleinfoRegisterAPIResponse {
	return poolTmallCarcenterVehicleinfoRegisterAPIResponse.Get().(*TmallCarcenterVehicleinfoRegisterAPIResponse)
}

// ReleaseTmallCarcenterVehicleinfoRegisterAPIResponse 将 TmallCarcenterVehicleinfoRegisterAPIResponse 保存到 sync.Pool
func ReleaseTmallCarcenterVehicleinfoRegisterAPIResponse(v *TmallCarcenterVehicleinfoRegisterAPIResponse) {
	v.Reset()
	poolTmallCarcenterVehicleinfoRegisterAPIResponse.Put(v)
}
