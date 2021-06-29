package category

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
属性回填接口 API请求
alibaba.imap.pv.autofill

根据用户传入的标题、目标渠道id，目标渠道叶子类目，预测其对应的pv信息，返回给业务方，供其自动填充属性项属性值信息
*/
type AlibabaImapPvAutofillRequest struct {
    model.Params
    // 系统入参
    _topImapItemDo   *TopImapItemDo
}

// 初始化AlibabaImapPvAutofillRequest对象
func NewAlibabaImapPvAutofillRequest() *AlibabaImapPvAutofillRequest{
    return &AlibabaImapPvAutofillRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaImapPvAutofillRequest) GetApiMethodName() string {
    return "alibaba.imap.pv.autofill"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaImapPvAutofillRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TopImapItemDo Setter
// 系统入参
func (r *AlibabaImapPvAutofillRequest) SetTopImapItemDo(_topImapItemDo *TopImapItemDo) error {
    r._topImapItemDo = _topImapItemDo
    r.Set("top_imap_item_do", _topImapItemDo)
    return nil
}

// TopImapItemDo Getter
func (r AlibabaImapPvAutofillRequest) GetTopImapItemDo() *TopImapItemDo {
    return r._topImapItemDo
}
