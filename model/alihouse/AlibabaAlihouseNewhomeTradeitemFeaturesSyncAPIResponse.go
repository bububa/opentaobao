package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeTradeitemFeaturesSyncAPIResponse 同步品活动标 API返回值
// alibaba.alihouse.newhome.tradeitem.features.sync
//
// 同步品活动标
type AlibabaAlihouseNewhomeTradeitemFeaturesSyncAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseNewhomeTradeitemFeaturesSyncAPIResponseModel
}

// AlibabaAlihouseNewhomeTradeitemFeaturesSyncAPIResponseModel is 同步品活动标 成功返回结果
type AlibabaAlihouseNewhomeTradeitemFeaturesSyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_newhome_tradeitem_features_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseNewhomeTradeitemFeaturesSyncResult `json:"result,omitempty" xml:"result,omitempty"`
}
