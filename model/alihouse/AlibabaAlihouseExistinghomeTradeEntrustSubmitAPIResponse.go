package alihouse

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseExistinghomeTradeEntrustSubmitAPIResponse 交易委托信息更新接口 API返回值
// alibaba.alihouse.existinghome.trade.entrust.submit
//
// 交易委托信息更新接口
type AlibabaAlihouseExistinghomeTradeEntrustSubmitAPIResponse struct {
	model.CommonResponse
	AlibabaAlihouseExistinghomeTradeEntrustSubmitAPIResponseModel
}

// AlibabaAlihouseExistinghomeTradeEntrustSubmitAPIResponseModel is 交易委托信息更新接口 成功返回结果
type AlibabaAlihouseExistinghomeTradeEntrustSubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihouse_existinghome_trade_entrust_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAlihouseExistinghomeTradeEntrustSubmitResult `json:"result,omitempty" xml:"result,omitempty"`
}
