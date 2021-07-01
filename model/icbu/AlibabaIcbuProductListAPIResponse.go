package icbu

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIcbuProductListAPIResponse
商品查询 API返回值
alibaba.icbu.product.list

根据类目ID和商品名称查询商品概要信息。结果以修改时间倒序返回，支持分页，每页最多30个。每次调用都是独立的请求，不记录调用的上下文。 */
type AlibabaIcbuProductListAPIResponse struct {
	model.CommonResponse
	AlibabaIcbuProductListAPIResponseModel
}

// AlibabaIcbuProductListAPIResponseModel is 商品查询 成功返回结果
type AlibabaIcbuProductListAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_icbu_product_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 总数
	TotalItem int64 `json:"total_item,omitempty" xml:"total_item,omitempty"`
	// 当前页
	CurrentPage int64 `json:"current_page,omitempty" xml:"current_page,omitempty"`
	// 每页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 商品概要信息列表
	Products []AlibabaProductBriefResponse `json:"products,omitempty" xml:"products>alibaba_product_brief_response,omitempty"`
}
