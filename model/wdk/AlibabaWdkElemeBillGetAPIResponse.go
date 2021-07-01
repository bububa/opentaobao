package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkElemeBillGetAPIResponse
饿了么日维度对账单查询 API返回值
alibaba.wdk.eleme.bill.get

查询饿了么日维度对账单信息 */
type AlibabaWdkElemeBillGetAPIResponse struct {
	model.CommonResponse
	AlibabaWdkElemeBillGetAPIResponseModel
}

// AlibabaWdkElemeBillGetAPIResponseModel is 饿了么日维度对账单查询 成功返回结果
type AlibabaWdkElemeBillGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_eleme_bill_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaWdkElemeBillGetApiResult `json:"result,omitempty" xml:"result,omitempty"`
}
