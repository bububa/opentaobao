package vms

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
新能源车--外部车辆信息回传 API返回值 
cainiao.vms.service.vehicleinfo.upload

新能源车--外部车辆信息回传
*/
type CainiaoVmsServiceVehicleinfoUploadAPIResponse struct {
    model.CommonResponse
    CainiaoVmsServiceVehicleinfoUploadAPIResponseModel
}

// 新能源车--外部车辆信息回传 成功返回结果
type CainiaoVmsServiceVehicleinfoUploadAPIResponseModel struct {
    XMLName xml.Name `xml:"cainiao_vms_service_vehicleinfo_upload_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // result
    Result   *AppBaseResponse `json:"result,omitempty" xml:"result,omitempty"`
}
