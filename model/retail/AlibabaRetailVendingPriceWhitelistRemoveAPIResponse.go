package retail

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailVendingPriceWhitelistRemoveAPIResponse 价格管控白名单去除 API返回值
// alibaba.retail.vending.price.whitelist.remove
//
// 商家价格管控白名单去除
type AlibabaRetailVendingPriceWhitelistRemoveAPIResponse struct {
	model.CommonResponse
	AlibabaRetailVendingPriceWhitelistRemoveAPIResponseModel
}

// AlibabaRetailVendingPriceWhitelistRemoveAPIResponseModel is 价格管控白名单去除 成功返回结果
type AlibabaRetailVendingPriceWhitelistRemoveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_retail_vending_price_whitelist_remove_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaRetailVendingPriceWhitelistRemoveResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
