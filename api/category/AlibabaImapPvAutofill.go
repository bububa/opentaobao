package category

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/category"
)

/* 
属性回填接口 
alibaba.imap.pv.autofill

根据用户传入的标题、目标渠道id，目标渠道叶子类目，预测其对应的pv信息，返回给业务方，供其自动填充属性项属性值信息
*/
func AlibabaImapPvAutofill(clt *core.SDKClient, req *category.AlibabaImapPvAutofillRequest, session string) (*category.AlibabaImapPvAutofillResponse, error) {
    var resp category.AlibabaImapPvAutofillAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return resp.Response, nil
}
