package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
菜鸟裹裹运单状态查询 API请求
taobao.omniorder.store.sdtstatus

提供给商家查询运力单的状态。
*/
type TaobaoOmniorderStoreSdtstatusRequest struct {
    model.Params
    // 菜鸟裹裹的包裹ID
    _packageId   int64
}

// 初始化TaobaoOmniorderStoreSdtstatusRequest对象
func NewTaobaoOmniorderStoreSdtstatusRequest() *TaobaoOmniorderStoreSdtstatusRequest{
    return &TaobaoOmniorderStoreSdtstatusRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOmniorderStoreSdtstatusRequest) GetApiMethodName() string {
    return "taobao.omniorder.store.sdtstatus"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOmniorderStoreSdtstatusRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PackageId Setter
// 菜鸟裹裹的包裹ID
func (r *TaobaoOmniorderStoreSdtstatusRequest) SetPackageId(_packageId int64) error {
    r._packageId = _packageId
    r.Set("package_id", _packageId)
    return nil
}

// PackageId Getter
func (r TaobaoOmniorderStoreSdtstatusRequest) GetPackageId() int64 {
    return r._packageId
}
