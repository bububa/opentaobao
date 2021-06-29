package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
速店通查询站点信息 API请求
taobao.omniorder.store.sdtquerystation

速店通查询站点信息
*/
type TaobaoOmniorderStoreSdtquerystationRequest struct {
    model.Params
    // 取号时返回的packageId
    paramLong2   int64
}

// 初始化TaobaoOmniorderStoreSdtquerystationRequest对象
func NewTaobaoOmniorderStoreSdtquerystationRequest() *TaobaoOmniorderStoreSdtquerystationRequest{
    return &TaobaoOmniorderStoreSdtquerystationRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOmniorderStoreSdtquerystationRequest) GetApiMethodName() string {
    return "taobao.omniorder.store.sdtquerystation"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOmniorderStoreSdtquerystationRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamLong2 Setter
// 取号时返回的packageId
func (r *TaobaoOmniorderStoreSdtquerystationRequest) SetParamLong2(paramLong2 int64) error {
    r.paramLong2 = paramLong2
    r.Set("param_long2", paramLong2)
    return nil
}

// ParamLong2 Getter
func (r TaobaoOmniorderStoreSdtquerystationRequest) GetParamLong2() int64 {
    return r.paramLong2
}
