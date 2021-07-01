package alihealth2

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthMsAreaProvinceListAPIResponse
疫苗预约省份列表查询 API返回值
alibaba.alihealth.ms.area.province.list

微信小程序疫苗预约省份列表查询 */
type AlibabaAlihealthMsAreaProvinceListAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthMsAreaProvinceListAPIResponseModel
}

// AlibabaAlihealthMsAreaProvinceListAPIResponseModel is 疫苗预约省份列表查询 成功返回结果
type AlibabaAlihealthMsAreaProvinceListAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_ms_area_province_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ServiceResult `json:"result,omitempty" xml:"result,omitempty"`
}
