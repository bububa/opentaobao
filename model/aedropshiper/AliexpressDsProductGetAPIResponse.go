package aedropshiper

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AliexpressDsProductGetAPIResponse 商品信息查询 API返回值
// aliexpress.ds.product.get
//
// 商品信息查询
type AliexpressDsProductGetAPIResponse struct {
	model.CommonResponse
	AliexpressDsProductGetAPIResponseModel
}

// AliexpressDsProductGetAPIResponseModel is 商品信息查询 成功返回结果
type AliexpressDsProductGetAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_ds_product_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// Error message
	RspMsg string `json:"rsp_msg,omitempty" xml:"rsp_msg,omitempty"`
	// Error code
	RspCode string `json:"rsp_code,omitempty" xml:"rsp_code,omitempty"`
	// Product search results
	Result *AeItemQueryResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
