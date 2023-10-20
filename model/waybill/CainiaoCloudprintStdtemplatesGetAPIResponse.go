package waybill

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// CainiaocloudprintstdtemplatesgetAPIResponse 获取所有的菜鸟标准电子面单模板 API返回值
// cainiao.cloudprint.stdtemplates.get
//
// 获取菜鸟标准电子面单模板
type CainiaocloudprintstdtemplatesgetAPIResponse struct {
	model.CommonResponse
	CainiaocloudprintstdtemplatesgetAPIResponseModel
}

// CainiaocloudprintstdtemplatesgetAPIResponseModel is 获取所有的菜鸟标准电子面单模板 成功返回结果
type CainiaocloudprintstdtemplatesgetAPIResponseModel struct {
	XMLName xml.Name `xml:"cainiao_cloudprint_stdtemplates_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果集
	Result *CloudPrintBaseResult `json:"result,omitempty" xml:"result,omitempty"`
}
