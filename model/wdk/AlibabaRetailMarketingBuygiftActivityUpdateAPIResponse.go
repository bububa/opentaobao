package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaRetailMarketingBuygiftActivityUpdateAPIResponse
更新单品买赠活动 API返回值
alibaba.retail.marketing.buygift.activity.update

同城零售单品买赠活动更新 */
type AlibabaRetailMarketingBuygiftActivityUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaRetailMarketingBuygiftActivityUpdateAPIResponseModel
}

// AlibabaRetailMarketingBuygiftActivityUpdateAPIResponseModel is 更新单品买赠活动 成功返回结果
type AlibabaRetailMarketingBuygiftActivityUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_retail_marketing_buygift_activity_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作结果
	Result *OctopusOpenResult `json:"result,omitempty" xml:"result,omitempty"`
}
