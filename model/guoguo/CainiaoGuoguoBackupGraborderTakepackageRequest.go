package guoguo

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
兜底派送订单的揽件接口 API请求
cainiao.guoguo.backup.graborder.takepackage

快递公司回传订单号和四位取件码给菜鸟裹裹
*/
type CainiaoGuoguoBackupGraborderTakepackageRequest struct {
    model.Params
    // 物流订单号
    orderCode   string
    // 包裹四位码
    packageCode   string
}

// 初始化CainiaoGuoguoBackupGraborderTakepackageRequest对象
func NewCainiaoGuoguoBackupGraborderTakepackageRequest() *CainiaoGuoguoBackupGraborderTakepackageRequest{
    return &CainiaoGuoguoBackupGraborderTakepackageRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoGuoguoBackupGraborderTakepackageRequest) GetApiMethodName() string {
    return "cainiao.guoguo.backup.graborder.takepackage"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoGuoguoBackupGraborderTakepackageRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderCode Setter
// 物流订单号
func (r *CainiaoGuoguoBackupGraborderTakepackageRequest) SetOrderCode(orderCode string) error {
    r.orderCode = orderCode
    r.Set("orderCode", orderCode)
    return nil
}

// OrderCode Getter
func (r CainiaoGuoguoBackupGraborderTakepackageRequest) GetOrderCode() string {
    return r.orderCode
}
// PackageCode Setter
// 包裹四位码
func (r *CainiaoGuoguoBackupGraborderTakepackageRequest) SetPackageCode(packageCode string) error {
    r.packageCode = packageCode
    r.Set("packageCode", packageCode)
    return nil
}

// PackageCode Getter
func (r CainiaoGuoguoBackupGraborderTakepackageRequest) GetPackageCode() string {
    return r.packageCode
}
