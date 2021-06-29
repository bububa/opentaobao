package eticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取包基本信息 API请求
taobao.vmarket.eticket.package.base.get

获取包基本信息
*/
type TaobaoVmarketEticketPackageBaseGetRequest struct {
    model.Params
    // 包id
    packageId   int64
}

// 初始化TaobaoVmarketEticketPackageBaseGetRequest对象
func NewTaobaoVmarketEticketPackageBaseGetRequest() *TaobaoVmarketEticketPackageBaseGetRequest{
    return &TaobaoVmarketEticketPackageBaseGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoVmarketEticketPackageBaseGetRequest) GetApiMethodName() string {
    return "taobao.vmarket.eticket.package.base.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoVmarketEticketPackageBaseGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// PackageId Setter
// 包id
func (r *TaobaoVmarketEticketPackageBaseGetRequest) SetPackageId(packageId int64) error {
    r.packageId = packageId
    r.Set("package_id", packageId)
    return nil
}

// PackageId Getter
func (r TaobaoVmarketEticketPackageBaseGetRequest) GetPackageId() int64 {
    return r.packageId
}
