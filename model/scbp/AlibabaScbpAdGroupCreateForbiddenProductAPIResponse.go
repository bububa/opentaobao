package scbp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpadgroupcreateforbiddenproductAPIResponse 创建屏蔽品 API返回值
// alibaba.scbp.ad.group.create.forbidden.product
//
// 创建屏蔽品
type AlibabascbpadgroupcreateforbiddenproductAPIResponse struct {
	model.CommonResponse
	AlibabascbpadgroupcreateforbiddenproductAPIResponseModel
}

// AlibabascbpadgroupcreateforbiddenproductAPIResponseModel is 创建屏蔽品 成功返回结果
type AlibabascbpadgroupcreateforbiddenproductAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_group_create_forbidden_product_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result int64 `json:"result,omitempty" xml:"result,omitempty"`
}
