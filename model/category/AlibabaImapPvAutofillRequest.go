package category

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
属性回填接口 APIRequest
alibaba.imap.pv.autofill

根据用户传入的标题、目标渠道id，目标渠道叶子类目，预测其对应的pv信息，返回给业务方，供其自动填充属性项属性值信息
*/
type AlibabaImapPvAutofillRequest struct {
    model.Params

    // 系统入参
    topImapItemDo   *TopImapItemDo 

}

func NewAlibabaImapPvAutofillRequest() *AlibabaImapPvAutofillRequest{
    return &AlibabaImapPvAutofillRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaImapPvAutofillRequest) GetApiMethodName() string {
    return "alibaba.imap.pv.autofill"
}

func (r AlibabaImapPvAutofillRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaImapPvAutofillRequest) SetTopImapItemDo(topImapItemDo *TopImapItemDo) error {
    r.topImapItemDo = topImapItemDo
    r.Set("top_imap_item_do", topImapItemDo)
    return nil
}

func (r AlibabaImapPvAutofillRequest) GetTopImapItemDo() *TopImapItemDo {
    return r.topImapItemDo
}

