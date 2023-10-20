package waybill

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// CainiaocloudprintsinglecustomareagetAPIResponse 获取商家单一自定义区 API返回值
// cainiao.cloudprint.single.customarea.get
//
// 商家所有快递公司模板只有一个自定义区
type CainiaocloudprintsinglecustomareagetAPIResponse struct {
	model.CommonResponse
	CainiaocloudprintsinglecustomareagetAPIResponseModel
}

// CainiaocloudprintsinglecustomareagetAPIResponseModel is 获取商家单一自定义区 成功返回结果
type CainiaocloudprintsinglecustomareagetAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_cloudprint_single_customarea_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *CloudPrintBaseResult `json:"result,omitempty" xml:"result,omitempty"`
}
