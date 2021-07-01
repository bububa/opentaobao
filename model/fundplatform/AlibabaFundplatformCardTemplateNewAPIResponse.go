package fundplatform

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
新增实体卡模板 API返回值 
alibaba.fundplatform.card.template.new

该接口由制卡商实现，当新增一个实体卡模板的时候，需要调用该接口，通知制卡商同步新增卡模板信息。
*/
type AlibabaFundplatformCardTemplateNewAPIResponse struct {
    model.CommonResponse
    AlibabaFundplatformCardTemplateNewAPIResponseModel
}

// 新增实体卡模板 成功返回结果
type AlibabaFundplatformCardTemplateNewAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_fundplatform_card_template_new_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 错误详情
    ResultMessage   string `json:"result_message,omitempty" xml:"result_message,omitempty"`
    // 错误CODE
    ResultCode   string `json:"result_code,omitempty" xml:"result_code,omitempty"`
    // 是否调用成功
    Success   bool `json:"success,omitempty" xml:"success,omitempty"`
}
