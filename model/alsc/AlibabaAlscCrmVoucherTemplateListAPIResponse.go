package alsc

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalsccrmvouchertemplatelistAPIResponse 获取优惠券模版列表 API返回值
// alibaba.alsc.crm.voucher.template.list
//
// 获取优惠券模版列表
type AlibabaalsccrmvouchertemplatelistAPIResponse struct {
	model.CommonResponse
	AlibabaalsccrmvouchertemplatelistAPIResponseModel
}

// AlibabaalsccrmvouchertemplatelistAPIResponseModel is 获取优惠券模版列表 成功返回结果
type AlibabaalsccrmvouchertemplatelistAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_crm_voucher_template_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 分页返回模型
	Result *CommonPageResult `json:"result,omitempty" xml:"result,omitempty"`
}
