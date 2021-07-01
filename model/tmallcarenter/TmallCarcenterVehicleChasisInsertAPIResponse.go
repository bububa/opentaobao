package tmallcarenter

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
EPC车型底盘压缩库新增接口 API返回值 
tmall.carcenter.vehicle.chasis.insert

EPC车型底盘压缩库新增接口
*/
type TmallCarcenterVehicleChasisInsertAPIResponse struct {
    model.CommonResponse
    TmallCarcenterVehicleChasisInsertAPIResponseModel
}

// EPC车型底盘压缩库新增接口 成功返回结果
type TmallCarcenterVehicleChasisInsertAPIResponseModel struct {
    XMLName xml.Name `xml:"tmall_carcenter_vehicle_chasis_insert_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *TmallCarcenterVehicleChasisInsertResult `json:"result,omitempty" xml:"result,omitempty"`
}
