package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIResponse
营销商品数据同步 API返回值
alibaba.wdk.marketing.open.darunfa.activity.sku.sync

大润发营销商品数据同步 */
type AlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIResponseModel
}

// AlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIResponseModel is 营销商品数据同步 成功返回结果
type AlibabaWdkMarketingOpenDarunfaActivitySkuSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_open_darunfa_activity_sku_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果信息
	Result *WdkMarketOpenResult `json:"result,omitempty" xml:"result,omitempty"`
}
