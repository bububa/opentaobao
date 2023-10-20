package tmallcarenter

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallCarcenterVehicleChasisInsertAPIResponse EPC车型底盘压缩库新增接口 API返回值
// tmall.carcenter.vehicle.chasis.insert
//
// EPC车型底盘压缩库新增接口
type TmallCarcenterVehicleChasisInsertAPIResponse struct {
	model.CommonResponse
	TmallCarcenterVehicleChasisInsertAPIResponseModel
}

// Reset 清空结构体
func (m *TmallCarcenterVehicleChasisInsertAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallCarcenterVehicleChasisInsertAPIResponseModel).Reset()
}

// TmallCarcenterVehicleChasisInsertAPIResponseModel is EPC车型底盘压缩库新增接口 成功返回结果
type TmallCarcenterVehicleChasisInsertAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_carcenter_vehicle_chasis_insert_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TmallCarcenterVehicleChasisInsertResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallCarcenterVehicleChasisInsertAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallCarcenterVehicleChasisInsertAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallCarcenterVehicleChasisInsertAPIResponse)
	},
}

// GetTmallCarcenterVehicleChasisInsertAPIResponse 从 sync.Pool 获取 TmallCarcenterVehicleChasisInsertAPIResponse
func GetTmallCarcenterVehicleChasisInsertAPIResponse() *TmallCarcenterVehicleChasisInsertAPIResponse {
	return poolTmallCarcenterVehicleChasisInsertAPIResponse.Get().(*TmallCarcenterVehicleChasisInsertAPIResponse)
}

// ReleaseTmallCarcenterVehicleChasisInsertAPIResponse 将 TmallCarcenterVehicleChasisInsertAPIResponse 保存到 sync.Pool
func ReleaseTmallCarcenterVehicleChasisInsertAPIResponse(v *TmallCarcenterVehicleChasisInsertAPIResponse) {
	v.Reset()
	poolTmallCarcenterVehicleChasisInsertAPIResponse.Put(v)
}
