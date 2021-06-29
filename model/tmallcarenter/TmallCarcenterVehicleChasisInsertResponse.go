package tmallcarenter

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
EPC车型底盘压缩库新增接口 APIResponse
tmall.carcenter.vehicle.chasis.insert

EPC车型底盘压缩库新增接口
*/
type TmallCarcenterVehicleChasisInsertAPIResponse struct {
    model.CommonResponse
    TmallCarcenterVehicleChasisInsertResponse
}

type TmallCarcenterVehicleChasisInsertResponse struct {
    XMLName xml.Name `xml:"tmall_carcenter_vehicle_chasis_insert_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *TmallCarcenterVehicleChasisInsertResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
