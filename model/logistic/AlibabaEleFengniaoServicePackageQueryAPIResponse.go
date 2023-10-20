package logistic

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaelefengniaoservicepackagequeryAPIResponse 预采购服务包查询接口 API返回值
// alibaba.ele.fengniao.service.package.query
//
// 查询门店所在经纬度可用服务包的接口
type AlibabaelefengniaoservicepackagequeryAPIResponse struct {
	model.CommonResponse
	AlibabaelefengniaoservicepackagequeryAPIResponseModel
}

// AlibabaelefengniaoservicepackagequeryAPIResponseModel is 预采购服务包查询接口 成功返回结果
type AlibabaelefengniaoservicepackagequeryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ele_fengniao_service_package_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// servicePackages
	ServicePackages []AlibabaelefengniaoservicepackagequeryResult `json:"service_packages,omitempty" xml:"service_packages>alibabaelefengniaoservicepackagequery_result,omitempty"`
}
