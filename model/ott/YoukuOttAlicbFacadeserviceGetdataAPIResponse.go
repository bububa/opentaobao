package ott

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// YoukuottalicbfacadeservicegetdataAPIResponse 影视SDK获取设备能力值 API返回值
// youku.ott.alicb.facadeservice.getdata
//
// 影视SDK获取设备能力值
type YoukuottalicbfacadeservicegetdataAPIResponse struct {
	model.CommonResponse
	YoukuottalicbfacadeservicegetdataAPIResponseModel
}

// YoukuottalicbfacadeservicegetdataAPIResponseModel is 影视SDK获取设备能力值 成功返回结果
type YoukuottalicbfacadeservicegetdataAPIResponseModel struct {
	XMLName xml.Name `xml:"youku_ott_alicb_facadeservice_getdata_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 设备能力JSON
	Model string `json:"model,omitempty" xml:"model,omitempty"`
}
