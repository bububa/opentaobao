package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkbmpaiyangsupplypublishAPIResponse 派样商品库存变更同步接口 API返回值
// alibaba.wdk.bm.paiyang.supply.publish
//
// 淘鲜达接入第三方进行派样，第三方同步大仓和门店的库存变更信息。
type AlibabawdkbmpaiyangsupplypublishAPIResponse struct {
	model.CommonResponse
	AlibabawdkbmpaiyangsupplypublishAPIResponseModel
}

// AlibabawdkbmpaiyangsupplypublishAPIResponseModel is 派样商品库存变更同步接口 成功返回结果
type AlibabawdkbmpaiyangsupplypublishAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_bm_paiyang_supply_publish_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求出参
	Result *BmResult `json:"result,omitempty" xml:"result,omitempty"`
}
