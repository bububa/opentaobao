package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
用户查询加固结果 APIRequest
alibaba.security.jaq.app.shieldresult.get

用户通过alibaba.security.jaq.app.shield接口提交应用加固后,通过该接口查询加固结果,下载加固包
*/
type AlibabaSecurityJaqAppShieldresultGetRequest struct {
    model.Params

    // 任务唯一标识
    itemId   string 

}

func NewAlibabaSecurityJaqAppShieldresultGetRequest() *AlibabaSecurityJaqAppShieldresultGetRequest{
    return &AlibabaSecurityJaqAppShieldresultGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaSecurityJaqAppShieldresultGetRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.app.shieldresult.get"
}

func (r AlibabaSecurityJaqAppShieldresultGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaSecurityJaqAppShieldresultGetRequest) SetItemId(itemId string) error {
    r.itemId = itemId
    r.Set("item_id", itemId)
    return nil
}

func (r AlibabaSecurityJaqAppShieldresultGetRequest) GetItemId() string {
    return r.itemId
}

