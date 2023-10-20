package tbitem

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaitemoperatedownshelfAPIResponse 商品下架 API返回值
// alibaba.item.operate.downshelf
//
// 商品下架
type AlibabaitemoperatedownshelfAPIResponse struct {
	model.CommonResponse
	AlibabaitemoperatedownshelfAPIResponseModel
}

// AlibabaitemoperatedownshelfAPIResponseModel is 商品下架 成功返回结果
type AlibabaitemoperatedownshelfAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_item_operate_downshelf_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品下架是否成功
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}
