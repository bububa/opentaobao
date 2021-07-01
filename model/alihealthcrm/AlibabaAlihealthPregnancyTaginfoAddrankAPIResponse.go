package alihealthcrm

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthPregnancyTaginfoAddrankAPIResponse
点击标签后排序接口 API返回值
alibaba.alihealth.pregnancy.taginfo.addrank

备孕管理--点击标签后排序接口 */
type AlibabaAlihealthPregnancyTaginfoAddrankAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthPregnancyTaginfoAddrankAPIResponseModel
}

// AlibabaAlihealthPregnancyTaginfoAddrankAPIResponseModel is 点击标签后排序接口 成功返回结果
type AlibabaAlihealthPregnancyTaginfoAddrankAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_pregnancy_taginfo_addrank_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果集
	Result *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}
