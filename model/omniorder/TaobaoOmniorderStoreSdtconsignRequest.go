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
    packageId   string
    // 发货标签号
    tagCode   string
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
func (r *TaobaoOmniorderStoreSdtconsignRequest) SetPackageId(packageId string) error {
    r.packageId = packageId
    r.Set("package_id", packageId)
    return nil
}

// PackageId Getter
func (r TaobaoOmniorderStoreSdtconsignRequest) GetPackageId() string {
    return r.packageId
}
// TagCode Setter
// 发货标签号
func (r *TaobaoOmniorderStoreSdtconsignRequest) SetTagCode(tagCode string) error {
    r.tagCode = tagCode
    r.Set("tag_code", tagCode)
    return nil
}

// TagCode Getter
func (r TaobaoOmniorderStoreSdtconsignRequest) GetTagCode() string {
    return r.tagCode
}
