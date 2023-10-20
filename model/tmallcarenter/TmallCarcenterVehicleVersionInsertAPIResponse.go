package tmallcarenter

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallCarcenterVehicleVersionInsertAPIResponse 汽车EPC版本压缩库新增接口 API返回值
// tmall.carcenter.vehicle.version.insert
//
// 汽车EPC版本压缩库新增接口
type TmallCarcenterVehicleVersionInsertAPIResponse struct {
	model.CommonResponse
	TmallCarcenterVehicleVersionInsertAPIResponseModel
}

// Reset 清空结构体
func (m *TmallCarcenterVehicleVersionInsertAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TmallCarcenterVehicleVersionInsertAPIResponseModel).Reset()
}

// TmallCarcenterVehicleVersionInsertAPIResponseModel is 汽车EPC版本压缩库新增接口 成功返回结果
type TmallCarcenterVehicleVersionInsertAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_carcenter_vehicle_version_insert_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TmallCarcenterVehicleVersionInsertResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TmallCarcenterVehicleVersionInsertAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTmallCarcenterVehicleVersionInsertAPIResponse = sync.Pool{
	New: func() any {
		return new(TmallCarcenterVehicleVersionInsertAPIResponse)
	},
}

// GetTmallCarcenterVehicleVersionInsertAPIResponse 从 sync.Pool 获取 TmallCarcenterVehicleVersionInsertAPIResponse
func GetTmallCarcenterVehicleVersionInsertAPIResponse() *TmallCarcenterVehicleVersionInsertAPIResponse {
	return poolTmallCarcenterVehicleVersionInsertAPIResponse.Get().(*TmallCarcenterVehicleVersionInsertAPIResponse)
}

// ReleaseTmallCarcenterVehicleVersionInsertAPIResponse 将 TmallCarcenterVehicleVersionInsertAPIResponse 保存到 sync.Pool
func ReleaseTmallCarcenterVehicleVersionInsertAPIResponse(v *TmallCarcenterVehicleVersionInsertAPIResponse) {
	v.Reset()
	poolTmallCarcenterVehicleVersionInsertAPIResponse.Put(v)
}
