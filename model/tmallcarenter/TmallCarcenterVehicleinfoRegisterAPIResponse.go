package tmallcarenter

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
车型数据更新 API返回值 
tmall.carcenter.vehicleinfo.register

基本车型信息维护
*/
type TmallCarcenterVehicleinfoRegisterAPIResponse struct {
    model.CommonResponse
    TmallCarcenterVehicleinfoRegisterAPIResponseModel
}

// 车型数据更新 成功返回结果
type TmallCarcenterVehicleinfoRegisterAPIResponseModel struct {
    XMLName xml.Name `xml:"tmall_carcenter_vehicleinfo_register_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *TmallCarcenterVehicleinfoRegisterResult `json:"result,omitempty" xml:"result,omitempty"`
}
