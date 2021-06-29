package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
通知速店通取消取号 APIRequest
taobao.omniorder.store.sdtcancel

通知速店通取消取号
*/
type TaobaoOmniorderStoreSdtcancelRequest struct {
    model.Params

    // 取号返回的packageId
    packageId   int64 

}

func NewTaobaoOmniorderStoreSdtcancelRequest() *TaobaoOmniorderStoreSdtcancelRequest{
    return &TaobaoOmniorderStoreSdtcancelRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOmniorderStoreSdtcancelRequest) GetApiMethodName() string {
    return "taobao.omniorder.store.sdtcancel"
}

func (r TaobaoOmniorderStoreSdtcancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOmniorderStoreSdtcancelRequest) SetPackageId(packageId int64) error {
    r.packageId = packageId
    r.Set("package_id", packageId)
    return nil
}

func (r TaobaoOmniorderStoreSdtcancelRequest) GetPackageId() int64 {
    return r.packageId
}

