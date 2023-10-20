package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkScmLrpOrderPredictAPIResponse 单量预测 API返回值
// alibaba.wdk.scm.lrp.order.predict
//
// 提供基于门店和配送站的订单量预测，可用于提前安排人力资源
type AlibabaWdkScmLrpOrderPredictAPIResponse struct {
	model.CommonResponse
	AlibabaWdkScmLrpOrderPredictAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkScmLrpOrderPredictAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkScmLrpOrderPredictAPIResponseModel).Reset()
}

// AlibabaWdkScmLrpOrderPredictAPIResponseModel is 单量预测 成功返回结果
type AlibabaWdkScmLrpOrderPredictAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_scm_lrp_order_predict_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaWdkScmLrpOrderPredictApiResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkScmLrpOrderPredictAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkScmLrpOrderPredictAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkScmLrpOrderPredictAPIResponse)
	},
}

// GetAlibabaWdkScmLrpOrderPredictAPIResponse 从 sync.Pool 获取 AlibabaWdkScmLrpOrderPredictAPIResponse
func GetAlibabaWdkScmLrpOrderPredictAPIResponse() *AlibabaWdkScmLrpOrderPredictAPIResponse {
	return poolAlibabaWdkScmLrpOrderPredictAPIResponse.Get().(*AlibabaWdkScmLrpOrderPredictAPIResponse)
}

// ReleaseAlibabaWdkScmLrpOrderPredictAPIResponse 将 AlibabaWdkScmLrpOrderPredictAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkScmLrpOrderPredictAPIResponse(v *AlibabaWdkScmLrpOrderPredictAPIResponse) {
	v.Reset()
	poolAlibabaWdkScmLrpOrderPredictAPIResponse.Put(v)
}
