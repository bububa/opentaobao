package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkBmPaiyangSupplyPublishAPIResponse 派样商品库存变更同步接口 API返回值
// alibaba.wdk.bm.paiyang.supply.publish
//
// 淘鲜达接入第三方进行派样，第三方同步大仓和门店的库存变更信息。
type AlibabaWdkBmPaiyangSupplyPublishAPIResponse struct {
	model.CommonResponse
	AlibabaWdkBmPaiyangSupplyPublishAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkBmPaiyangSupplyPublishAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkBmPaiyangSupplyPublishAPIResponseModel).Reset()
}

// AlibabaWdkBmPaiyangSupplyPublishAPIResponseModel is 派样商品库存变更同步接口 成功返回结果
type AlibabaWdkBmPaiyangSupplyPublishAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_bm_paiyang_supply_publish_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 请求出参
	Result *BmResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkBmPaiyangSupplyPublishAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkBmPaiyangSupplyPublishAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkBmPaiyangSupplyPublishAPIResponse)
	},
}

// GetAlibabaWdkBmPaiyangSupplyPublishAPIResponse 从 sync.Pool 获取 AlibabaWdkBmPaiyangSupplyPublishAPIResponse
func GetAlibabaWdkBmPaiyangSupplyPublishAPIResponse() *AlibabaWdkBmPaiyangSupplyPublishAPIResponse {
	return poolAlibabaWdkBmPaiyangSupplyPublishAPIResponse.Get().(*AlibabaWdkBmPaiyangSupplyPublishAPIResponse)
}

// ReleaseAlibabaWdkBmPaiyangSupplyPublishAPIResponse 将 AlibabaWdkBmPaiyangSupplyPublishAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkBmPaiyangSupplyPublishAPIResponse(v *AlibabaWdkBmPaiyangSupplyPublishAPIResponse) {
	v.Reset()
	poolAlibabaWdkBmPaiyangSupplyPublishAPIResponse.Put(v)
}
