package scbp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbpeffectproductsinglegetAPIResponse 单个产品的报表 API返回值
// alibaba.scbp.effect.product.single.get
//
// 单个产品的报表
type AlibabascbpeffectproductsinglegetAPIResponse struct {
	model.CommonResponse
	AlibabascbpeffectproductsinglegetAPIResponseModel
}

// AlibabascbpeffectproductsinglegetAPIResponseModel is 单个产品的报表 成功返回结果
type AlibabascbpeffectproductsinglegetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_effect_product_single_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 单个产品的效果数据列表
	SProductEffectList []SingleProductEffectDto `json:"s_product_effect_list,omitempty" xml:"s_product_effect_list>single_product_effect_dto,omitempty"`
	// 总个数
	TotalNum int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
	// 总页数
	TotalPage int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`
}
