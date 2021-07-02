package icbu

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuProductGroupGetAPIResponse 分组信息获取 API返回值
// alibaba.icbu.product.group.get
//
// 分组信息获取
type AlibabaIcbuProductGroupGetAPIResponse struct {
	model.CommonResponse
	AlibabaIcbuProductGroupGetAPIResponseModel
}

// AlibabaIcbuProductGroupGetAPIResponseModel is 分组信息获取 成功返回结果
type AlibabaIcbuProductGroupGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_icbu_product_group_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 分组信息
	ProductGroup *ProductGroup `json:"product_group,omitempty" xml:"product_group,omitempty"`
}
