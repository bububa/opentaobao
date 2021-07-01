package category

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
属性回填接口 API返回值 
alibaba.imap.pv.autofill

根据用户传入的标题、目标渠道id，目标渠道叶子类目，预测其对应的pv信息，返回给业务方，供其自动填充属性项属性值信息
*/
type AlibabaImapPvAutofillAPIResponse struct {
    model.CommonResponse
    AlibabaImapPvAutofillAPIResponseModel
}

// 属性回填接口 成功返回结果
type AlibabaImapPvAutofillAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_imap_pv_autofill_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // imap通用返回DO
    Result   *TopImapResultDo `json:"result,omitempty" xml:"result,omitempty"`
}
