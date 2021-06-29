package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
通知菜鸟裹裹发货 APIRequest
taobao.omniorder.store.sdtconsign

ISV取完单号后通知菜鸟裹裹发货
*/
type TaobaoOmniorderStoreSdtconsignRequest struct {
    model.Params

    // 取号接口返回的包裹id
    packageId   string 

    // 发货标签号
    tagCode   string 

}

func NewTaobaoOmniorderStoreSdtconsignRequest() *TaobaoOmniorderStoreSdtconsignRequest{
    return &TaobaoOmniorderStoreSdtconsignRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOmniorderStoreSdtconsignRequest) GetApiMethodName() string {
    return "taobao.omniorder.store.sdtconsign"
}

func (r TaobaoOmniorderStoreSdtconsignRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOmniorderStoreSdtconsignRequest) SetPackageId(packageId string) error {
    r.packageId = packageId
    r.Set("package_id", packageId)
    return nil
}

func (r TaobaoOmniorderStoreSdtconsignRequest) GetPackageId() string {
    return r.packageId
}

func (r *TaobaoOmniorderStoreSdtconsignRequest) SetTagCode(tagCode string) error {
    r.tagCode = tagCode
    r.Set("tag_code", tagCode)
    return nil
}

func (r TaobaoOmniorderStoreSdtconsignRequest) GetTagCode() string {
    return r.tagCode
}

