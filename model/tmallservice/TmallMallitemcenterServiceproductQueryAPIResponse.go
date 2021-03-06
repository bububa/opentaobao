package tmallservice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallMallitemcenterServiceproductQueryAPIResponse 天猫服务商服务产品查询 API返回值
// tmall.mallitemcenter.serviceproduct.query
//
// 查询天猫服务的服务商发布的服务产品
type TmallMallitemcenterServiceproductQueryAPIResponse struct {
	model.CommonResponse
	TmallMallitemcenterServiceproductQueryAPIResponseModel
}

// TmallMallitemcenterServiceproductQueryAPIResponseModel is 天猫服务商服务产品查询 成功返回结果
type TmallMallitemcenterServiceproductQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_mallitemcenter_serviceproduct_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *TmallMallitemcenterServiceproductQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}
