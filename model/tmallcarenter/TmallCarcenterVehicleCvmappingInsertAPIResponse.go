package tmallcarenter

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallcarcentervehiclecvmappinginsertAPIResponse EPC车辆版本信息与底盘信息库关系绑定 API返回值
// tmall.carcenter.vehicle.cvmapping.insert
//
// EPC车辆版本信息与底盘信息库关系绑定
type TmallcarcentervehiclecvmappinginsertAPIResponse struct {
	model.CommonResponse
	TmallcarcentervehiclecvmappinginsertAPIResponseModel
}

// TmallcarcentervehiclecvmappinginsertAPIResponseModel is EPC车辆版本信息与底盘信息库关系绑定 成功返回结果
type TmallcarcentervehiclecvmappinginsertAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_carcenter_vehicle_cvmapping_insert_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TmallcarcentervehiclecvmappinginsertResult `json:"result,omitempty" xml:"result,omitempty"`
}
