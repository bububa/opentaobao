package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkorderSharestockInsuranceCallbackAPIResponse 共享库存订单投保后回传保单号 API返回值
// alibaba.wdkorder.sharestock.insurance.callback
//
// 共享库存订单投保消息获取
type AlibabaWdkorderSharestockInsuranceCallbackAPIResponse struct {
	model.CommonResponse
	AlibabaWdkorderSharestockInsuranceCallbackAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkorderSharestockInsuranceCallbackAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkorderSharestockInsuranceCallbackAPIResponseModel).Reset()
}

// AlibabaWdkorderSharestockInsuranceCallbackAPIResponseModel is 共享库存订单投保后回传保单号 成功返回结果
type AlibabaWdkorderSharestockInsuranceCallbackAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdkorder_sharestock_insurance_callback_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 系统自动生成
	Result *MaochaoOrderInsuranceCallbackResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkorderSharestockInsuranceCallbackAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkorderSharestockInsuranceCallbackAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkorderSharestockInsuranceCallbackAPIResponse)
	},
}

// GetAlibabaWdkorderSharestockInsuranceCallbackAPIResponse 从 sync.Pool 获取 AlibabaWdkorderSharestockInsuranceCallbackAPIResponse
func GetAlibabaWdkorderSharestockInsuranceCallbackAPIResponse() *AlibabaWdkorderSharestockInsuranceCallbackAPIResponse {
	return poolAlibabaWdkorderSharestockInsuranceCallbackAPIResponse.Get().(*AlibabaWdkorderSharestockInsuranceCallbackAPIResponse)
}

// ReleaseAlibabaWdkorderSharestockInsuranceCallbackAPIResponse 将 AlibabaWdkorderSharestockInsuranceCallbackAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkorderSharestockInsuranceCallbackAPIResponse(v *AlibabaWdkorderSharestockInsuranceCallbackAPIResponse) {
	v.Reset()
	poolAlibabaWdkorderSharestockInsuranceCallbackAPIResponse.Put(v)
}
