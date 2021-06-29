package dt

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
天猫汽车线索产品潜客数据 API请求
alibaba.dt.tmllcar.leadsinfo

1.	线索分发是天猫汽车行业流量端最中要的产品，经过前两年的业务和数据端的积累已经对整体业务流程和方案有了清晰的思路；目前数据段已经产沉淀2000W汽车潜客数据，通过运营尝试得到了较好的效果，今年将通过与商家端合作（大搜车-卖车管家）完成潜客分发-商家报价-潜客触达-线索分发-线下核销等一整个汽车人群运营闭环；这个接口反馈大搜车线下门店周围潜客规模及热门车型数据
*/
type AlibabaDtTmllcarLeadsinfoRequest struct {
    model.Params
    // shopcode
    shopCode   string
    // app_name
    appName   string
    // name
    name   string
    // pssword
    password   string
}

// 初始化AlibabaDtTmllcarLeadsinfoRequest对象
func NewAlibabaDtTmllcarLeadsinfoRequest() *AlibabaDtTmllcarLeadsinfoRequest{
    return &AlibabaDtTmllcarLeadsinfoRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaDtTmllcarLeadsinfoRequest) GetApiMethodName() string {
    return "alibaba.dt.tmllcar.leadsinfo"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaDtTmllcarLeadsinfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ShopCode Setter
// shopcode
func (r *AlibabaDtTmllcarLeadsinfoRequest) SetShopCode(shopCode string) error {
    r.shopCode = shopCode
    r.Set("shop_code", shopCode)
    return nil
}

// ShopCode Getter
func (r AlibabaDtTmllcarLeadsinfoRequest) GetShopCode() string {
    return r.shopCode
}
// AppName Setter
// app_name
func (r *AlibabaDtTmllcarLeadsinfoRequest) SetAppName(appName string) error {
    r.appName = appName
    r.Set("app_name", appName)
    return nil
}

// AppName Getter
func (r AlibabaDtTmllcarLeadsinfoRequest) GetAppName() string {
    return r.appName
}
// Name Setter
// name
func (r *AlibabaDtTmllcarLeadsinfoRequest) SetName(name string) error {
    r.name = name
    r.Set("name", name)
    return nil
}

// Name Getter
func (r AlibabaDtTmllcarLeadsinfoRequest) GetName() string {
    return r.name
}
// Password Setter
// pssword
func (r *AlibabaDtTmllcarLeadsinfoRequest) SetPassword(password string) error {
    r.password = password
    r.Set("password", password)
    return nil
}

// Password Getter
func (r AlibabaDtTmllcarLeadsinfoRequest) GetPassword() string {
    return r.password
}
