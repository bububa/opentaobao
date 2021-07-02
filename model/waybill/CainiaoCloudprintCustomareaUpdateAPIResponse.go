package waybill

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// CainiaoCloudprintCustomareaUpdateAPIResponse 自定义区内容更新 API返回值
// cainiao.cloudprint.customarea.update
//
// 自定义区内容更新
type CainiaoCloudprintCustomareaUpdateAPIResponse struct {
	model.CommonResponse
	CainiaoCloudprintCustomareaUpdateAPIResponseModel
}

// CainiaoCloudprintCustomareaUpdateAPIResponseModel is 自定义区内容更新 成功返回结果
type CainiaoCloudprintCustomareaUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_cloudprint_customarea_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *CloudPrintBaseResult `json:"result,omitempty" xml:"result,omitempty"`
}
