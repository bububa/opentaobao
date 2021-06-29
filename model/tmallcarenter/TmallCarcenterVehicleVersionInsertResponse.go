package tmallcarenter

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
汽车EPC版本压缩库新增接口 APIResponse
tmall.carcenter.vehicle.version.insert

汽车EPC版本压缩库新增接口
*/
type TmallCarcenterVehicleVersionInsertAPIResponse struct {
    model.CommonResponse
    TmallCarcenterVehicleVersionInsertResponse
}

type TmallCarcenterVehicleVersionInsertResponse struct {
    XMLName xml.Name `xml:"tmall_carcenter_vehicle_version_insert_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // result
    
    Result   *TmallCarcenterVehicleVersionInsertResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
