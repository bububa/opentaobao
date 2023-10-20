package trade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstVasTradeflowSaveAPIResponse 交易信息回流 API返回值
// alibaba.lst.vas.tradeflow.save
//
// 自动售货机交易信息同步接口，ISV通过此接口上传售货机交易信息。
type AlibabaLstVasTradeflowSaveAPIResponse struct {
	model.CommonResponse
	AlibabaLstVasTradeflowSaveAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLstVasTradeflowSaveAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLstVasTradeflowSaveAPIResponseModel).Reset()
}

// AlibabaLstVasTradeflowSaveAPIResponseModel is 交易信息回流 成功返回结果
type AlibabaLstVasTradeflowSaveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_vas_tradeflow_save_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaLstVasTradeflowSaveResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLstVasTradeflowSaveAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaLstVasTradeflowSaveAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLstVasTradeflowSaveAPIResponse)
	},
}

// GetAlibabaLstVasTradeflowSaveAPIResponse 从 sync.Pool 获取 AlibabaLstVasTradeflowSaveAPIResponse
func GetAlibabaLstVasTradeflowSaveAPIResponse() *AlibabaLstVasTradeflowSaveAPIResponse {
	return poolAlibabaLstVasTradeflowSaveAPIResponse.Get().(*AlibabaLstVasTradeflowSaveAPIResponse)
}

// ReleaseAlibabaLstVasTradeflowSaveAPIResponse 将 AlibabaLstVasTradeflowSaveAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLstVasTradeflowSaveAPIResponse(v *AlibabaLstVasTradeflowSaveAPIResponse) {
	v.Reset()
	poolAlibabaLstVasTradeflowSaveAPIResponse.Put(v)
}
