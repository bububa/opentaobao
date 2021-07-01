package alsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmCardQryAPIResponse
查询卡实例 API返回值
alibaba.alsc.crm.card.qry

查询卡实例（优先使用卡实例ID查询，没有则用物理卡号查询） */
type AlibabaAlscCrmCardQryAPIResponse struct {
	model.CommonResponse
	AlibabaAlscCrmCardQryAPIResponseModel
}

// AlibabaAlscCrmCardQryAPIResponseModel is 查询卡实例 成功返回结果
type AlibabaAlscCrmCardQryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_card_qry_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口结果
	Result *CommonResult `json:"result,omitempty" xml:"result,omitempty"`
}
