package retail

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaRetailVendingPriceWhitelistAddAPIResponse 贩卖机价格修改白名单 API返回值
// alibaba.retail.vending.price.whitelist.add
//
// 贩卖机价格修改白名单
type AlibabaRetailVendingPriceWhitelistAddAPIResponse struct {
	model.CommonResponse
	AlibabaRetailVendingPriceWhitelistAddAPIResponseModel
}

// AlibabaRetailVendingPriceWhitelistAddAPIResponseModel is 贩卖机价格修改白名单 成功返回结果
type AlibabaRetailVendingPriceWhitelistAddAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_retail_vending_price_whitelist_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaRetailVendingPriceWhitelistAddResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
