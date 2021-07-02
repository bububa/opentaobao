package product

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallProductSpecGetAPIResponse 根据产品规格的Id号获取当个的规格信息 API返回值
// tmall.product.spec.get
//
// 通过当个的spec_id获取到这个产品规格的信息，主要是因为产品规格是要经过审核的，所以通过这个接口可以获取到是否通过审核<br/>通过参看这个ProductSpec的status判断：<br/>1:表示审核通过<br/>3:表示等待审核。<br/>如果你的id找不到数据，那么就是审核被拒绝。
type TmallProductSpecGetAPIResponse struct {
	model.CommonResponse
	TmallProductSpecGetAPIResponseModel
}

// TmallProductSpecGetAPIResponseModel is 根据产品规格的Id号获取当个的规格信息 成功返回结果
type TmallProductSpecGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_product_spec_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回的产品规格信息，注意，这个产品规格信息可能是等待审核的，不一定可用。根据状态判断1：表示审核通过<br/>3：表示等待审核。
	ProductSpec *ProductSpec `json:"product_spec,omitempty" xml:"product_spec,omitempty"`
}
