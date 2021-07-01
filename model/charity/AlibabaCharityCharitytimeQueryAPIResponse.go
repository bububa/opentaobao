package charity

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCharityCharitytimeQueryAPIResponse
查询公益3小时公益时汇总 API返回值
alibaba.charity.charitytime.query

查询公益3小时公益时汇总 */
type AlibabaCharityCharitytimeQueryAPIResponse struct {
	model.CommonResponse
	AlibabaCharityCharitytimeQueryAPIResponseModel
}

// AlibabaCharityCharitytimeQueryAPIResponseModel is 查询公益3小时公益时汇总 成功返回结果
type AlibabaCharityCharitytimeQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_charity_charitytime_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码
	Apicode int64 `json:"apicode,omitempty" xml:"apicode,omitempty"`
	// 数据
	Data int64 `json:"data,omitempty" xml:"data,omitempty"`
	// 失败
	Fail bool `json:"fail,omitempty" xml:"fail,omitempty"`
	// 消息
	Apimsg string `json:"apimsg,omitempty" xml:"apimsg,omitempty"`
	// 成功
	Apisuccess bool `json:"apisuccess,omitempty" xml:"apisuccess,omitempty"`
}
