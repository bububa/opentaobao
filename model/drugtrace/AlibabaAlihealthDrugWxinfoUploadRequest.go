package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
小程序数据回传 API请求
alibaba.alihealth.drug.wxinfo.upload

小程序数据回传
*/
type AlibabaAlihealthDrugWxinfoUploadRequest struct {
    model.Params
    // 用户信息
    userInfo   string
    // 店铺名称
    shopInfo   string
    // 售货员信息
    salerInfo   string
    // 渠道
    isvChannel   string
}

// 初始化AlibabaAlihealthDrugWxinfoUploadRequest对象
func NewAlibabaAlihealthDrugWxinfoUploadRequest() *AlibabaAlihealthDrugWxinfoUploadRequest{
    return &AlibabaAlihealthDrugWxinfoUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugWxinfoUploadRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.wxinfo.upload"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugWxinfoUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UserInfo Setter
// 用户信息
func (r *AlibabaAlihealthDrugWxinfoUploadRequest) SetUserInfo(userInfo string) error {
    r.userInfo = userInfo
    r.Set("user_info", userInfo)
    return nil
}

// UserInfo Getter
func (r AlibabaAlihealthDrugWxinfoUploadRequest) GetUserInfo() string {
    return r.userInfo
}
// ShopInfo Setter
// 店铺名称
func (r *AlibabaAlihealthDrugWxinfoUploadRequest) SetShopInfo(shopInfo string) error {
    r.shopInfo = shopInfo
    r.Set("shop_info", shopInfo)
    return nil
}

// ShopInfo Getter
func (r AlibabaAlihealthDrugWxinfoUploadRequest) GetShopInfo() string {
    return r.shopInfo
}
// SalerInfo Setter
// 售货员信息
func (r *AlibabaAlihealthDrugWxinfoUploadRequest) SetSalerInfo(salerInfo string) error {
    r.salerInfo = salerInfo
    r.Set("saler_info", salerInfo)
    return nil
}

// SalerInfo Getter
func (r AlibabaAlihealthDrugWxinfoUploadRequest) GetSalerInfo() string {
    return r.salerInfo
}
// IsvChannel Setter
// 渠道
func (r *AlibabaAlihealthDrugWxinfoUploadRequest) SetIsvChannel(isvChannel string) error {
    r.isvChannel = isvChannel
    r.Set("isv_channel", isvChannel)
    return nil
}

// IsvChannel Getter
func (r AlibabaAlihealthDrugWxinfoUploadRequest) GetIsvChannel() string {
    return r.isvChannel
}
