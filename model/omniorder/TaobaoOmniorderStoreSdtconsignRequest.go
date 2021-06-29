package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
通知菜鸟裹裹发货 API请求
taobao.omniorder.store.sdtconsign

ISV取完单号后通知菜鸟裹裹发货
*/
type TaobaoOmniorderStoreSdtconsignRequest struct {
    model.Params
    // 取号接口返回的包裹id
    _packageId   string
    // 发货标签号
    _tagCode   string
}

// 初始化TaobaoOmniorderStoreSdtconsignRequest对象
func NewTaobaoOmniorderStoreSdtconsignRequest() *TaobaoOmniorderStoreSdtconsignRequest{
    return &TaobaoOmniorderStoreSdtconsignRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOmniorderStoreSdtconsignRequest) GetApiMethodName() string {
    return "taobao.omniorder.store.sdtconsign"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOmniorderStoreSdtconsignRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PackageId Setter
// 取号接口返回的包裹id
func (r *TaobaoOmniorderStoreSdtconsignRequest) SetPackageId(_packageId string) error {
    r._packageId = _packageId
    r.Set("package_id", _packageId)
    return nil
}

// PackageId Getter
func (r TaobaoOmniorderStoreSdtconsignRequest) GetPackageId() string {
    return r._packageId
}
// TagCode Setter
// 发货标签号
func (r *TaobaoOmniorderStoreSdtconsignRequest) SetTagCode(_tagCode string) error {
    r._tagCode = _tagCode
    r.Set("tag_code", _tagCode)
    return nil
}

// TagCode Getter
func (r TaobaoOmniorderStoreSdtconsignRequest) GetTagCode() string {
    return r._tagCode
}
