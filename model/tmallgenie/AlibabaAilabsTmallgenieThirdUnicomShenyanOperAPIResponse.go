package tmallgenie

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsTmallgenieThirdUnicomShenyanOperAPIResponse 联通神眼注册操作 API返回值
// alibaba.ailabs.tmallgenie.third.unicom.shenyan.oper
//
// 联通神眼注册操作
type AlibabaAilabsTmallgenieThirdUnicomShenyanOperAPIResponse struct {
	model.CommonResponse
	AlibabaAilabsTmallgenieThirdUnicomShenyanOperAPIResponseModel
}

// AlibabaAilabsTmallgenieThirdUnicomShenyanOperAPIResponseModel is 联通神眼注册操作 成功返回结果
type AlibabaAilabsTmallgenieThirdUnicomShenyanOperAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ailabs_tmallgenie_third_unicom_shenyan_oper_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 响应码（0000表示成功，其他表示失败）
	Result *AlibabaAilabsTmallgenieThirdUnicomShenyanOperResult `json:"result,omitempty" xml:"result,omitempty"`
}
