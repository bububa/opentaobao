package category

import (
    "github.com/bububa/opentaobao/model"
)

/* 
属性回填接口 APIResponse
alibaba.imap.pv.autofill

根据用户传入的标题、目标渠道id，目标渠道叶子类目，预测其对应的pv信息，返回给业务方，供其自动填充属性项属性值信息
*/
type AlibabaImapPvAutofillAPIResponse struct {
    model.CommonResponse
    Response *AlibabaImapPvAutofillResponse `json:"alibaba_imap_pv_autofill_response,omitempty"`
}

type AlibabaImapPvAutofillResponse struct {

    // imap通用返回DO
    Result   *TopImapResultDo `json:"result,omitempty"`

}
