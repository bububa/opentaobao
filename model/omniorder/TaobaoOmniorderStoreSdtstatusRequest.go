package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
菜鸟裹裹运单状态查询 APIRequest
taobao.omniorder.store.sdtstatus

提供给商家查询运力单的状态。
*/
type TaobaoOmniorderStoreSdtstatusRequest struct {
    model.Params

    // 菜鸟裹裹的包裹ID
    packageId   int64 

}

func NewTaobaoOmniorderStoreSdtstatusRequest() *TaobaoOmniorderStoreSdtstatusRequest{
    return &TaobaoOmniorderStoreSdtstatusRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOmniorderStoreSdtstatusRequest) GetApiMethodName() string {
    return "taobao.omniorder.store.sdtstatus"
}

func (r TaobaoOmniorderStoreSdtstatusRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOmniorderStoreSdtstatusRequest) SetPackageId(packageId int64) error {
    r.packageId = packageId
    r.Set("package_id", packageId)
    return nil
}

func (r TaobaoOmniorderStoreSdtstatusRequest) GetPackageId() int64 {
    return r.packageId
}

