package fenxiao

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TmallsupplychainchannelproductupshelfAPIResponse 产品上架 API返回值
// tmall.supplychain.channel.product.upshelf
//
// 上架渠道产品
type TmallsupplychainchannelproductupshelfAPIResponse struct {
	model.CommonResponse
	TmallsupplychainchannelproductupshelfAPIResponseModel
}

// TmallsupplychainchannelproductupshelfAPIResponseModel is 产品上架 成功返回结果
type TmallsupplychainchannelproductupshelfAPIResponseModel struct {
	XMLName xml.Name `xml:"tmall_supplychain_channel_product_upshelf_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *TmallsupplychainchannelproductupshelfResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
