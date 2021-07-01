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
type TaobaoOmniorderStoreSdtquerystationAPIRequest struct {
    model.Params
    // 取号时返回的packageId
    _paramLong2   int64
}

// 初始化TaobaoOmniorderStoreSdtquerystationAPIRequest对象
func NewTaobaoOmniorderStoreSdtquerystationRequest() *TaobaoOmniorderStoreSdtquerystationAPIRequest{
    return &TaobaoOmniorderStoreSdtquerystationAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOmniorderStoreSdtquerystationAPIRequest) GetApiMethodName() string {
    return "taobao.omniorder.store.sdtquerystation"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOmniorderStoreSdtquerystationAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamLong2 Setter
// 取号时返回的packageId
func (r *TaobaoOmniorderStoreSdtquerystationAPIRequest) SetParamLong2(_paramLong2 int64) error {
    r._paramLong2 = _paramLong2
    r.Set("param_long2", _paramLong2)
    return nil
}

// ParamLong2 Getter
func (r TaobaoOmniorderStoreSdtquerystationAPIRequest) GetParamLong2() int64 {
    return r._paramLong2
}
