package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
速店通查询站点信息 APIRequest
taobao.omniorder.store.sdtquerystation

速店通查询站点信息
*/
type TaobaoOmniorderStoreSdtquerystationRequest struct {
    model.Params

    // 取号时返回的packageId
    paramLong2   int64 

}

func NewTaobaoOmniorderStoreSdtquerystationRequest() *TaobaoOmniorderStoreSdtquerystationRequest{
    return &TaobaoOmniorderStoreSdtquerystationRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOmniorderStoreSdtquerystationRequest) GetApiMethodName() string {
    return "taobao.omniorder.store.sdtquerystation"
}

func (r TaobaoOmniorderStoreSdtquerystationRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOmniorderStoreSdtquerystationRequest) SetParamLong2(paramLong2 int64) error {
    r.paramLong2 = paramLong2
    r.Set("param_long2", paramLong2)
    return nil
}

func (r TaobaoOmniorderStoreSdtquerystationRequest) GetParamLong2() int64 {
    return r.paramLong2
}

