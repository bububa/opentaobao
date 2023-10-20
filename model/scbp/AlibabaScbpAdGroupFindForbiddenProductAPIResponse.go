package scbp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpadgroupfindforbiddenproductAPIResponse 查询屏蔽品 API返回值
// alibaba.scbp.ad.group.find.forbidden.product
//
// 查询屏蔽品
type AlibabascbpadgroupfindforbiddenproductAPIResponse struct {
	model.CommonResponse
	AlibabascbpadgroupfindforbiddenproductAPIResponseModel
}

// AlibabascbpadgroupfindforbiddenproductAPIResponseModel is 查询屏蔽品 成功返回结果
type AlibabascbpadgroupfindforbiddenproductAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_group_find_forbidden_product_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回列表
	ResultList []ForbiddenProductDto `json:"result_list,omitempty" xml:"result_list>forbidden_product_dto,omitempty"`
}
