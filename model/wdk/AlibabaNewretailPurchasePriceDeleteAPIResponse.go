package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabanewretailpurchasepricedeleteAPIResponse 共享库存 商户删除采购价 API返回值
// alibaba.newretail.purchase.price.delete
//
// 共享库存 商户删除采购价
type AlibabanewretailpurchasepricedeleteAPIResponse struct {
	model.CommonResponse
	AlibabanewretailpurchasepricedeleteAPIResponseModel
}

// AlibabanewretailpurchasepricedeleteAPIResponseModel is 共享库存 商户删除采购价 成功返回结果
type AlibabanewretailpurchasepricedeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_newretail_purchase_price_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 拆单结果对象
	Result *TopBaseResult `json:"result,omitempty" xml:"result,omitempty"`
}
