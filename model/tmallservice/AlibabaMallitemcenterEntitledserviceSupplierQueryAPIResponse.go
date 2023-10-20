package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabamallitemcenterentitledservicesupplierqueryAPIResponse 根据天猫id查询门店服务授权 API返回值
// alibaba.mallitemcenter.entitledservice.supplier.query
//
// 根据天猫id查询门店服务授权
type AlibabamallitemcenterentitledservicesupplierqueryAPIResponse struct {
	model.CommonResponse
	AlibabamallitemcenterentitledservicesupplierqueryAPIResponseModel
}

// AlibabamallitemcenterentitledservicesupplierqueryAPIResponseModel is 根据天猫id查询门店服务授权 成功返回结果
type AlibabamallitemcenterentitledservicesupplierqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_mallitemcenter_entitledservice_supplier_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 统一返回结果
	Result *AlibabamallitemcenterentitledservicesupplierqueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
