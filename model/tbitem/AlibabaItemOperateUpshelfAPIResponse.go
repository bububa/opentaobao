package tbitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaitemoperateupshelfAPIResponse 商品上架 API返回值
// alibaba.item.operate.upshelf
//
// 商品上架
type AlibabaitemoperateupshelfAPIResponse struct {
	model.CommonResponse
	AlibabaitemoperateupshelfAPIResponseModel
}

// AlibabaitemoperateupshelfAPIResponseModel is 商品上架 成功返回结果
type AlibabaitemoperateupshelfAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_item_operate_upshelf_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品上架是否成功
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}
