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
    _userInfo   string
    // 店铺名称
    _shopInfo   string
    // 售货员信息
    _salerInfo   string
    // 渠道
    _isvChannel   string
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
func (r *AlibabaAlihealthDrugWxinfoUploadRequest) SetUserInfo(_userInfo string) error {
    r._userInfo = _userInfo
    r.Set("user_info", _userInfo)
    return nil
}

// UserInfo Getter
func (r AlibabaAlihealthDrugWxinfoUploadRequest) GetUserInfo() string {
    return r._userInfo
}
// ShopInfo Setter
// 店铺名称
func (r *AlibabaAlihealthDrugWxinfoUploadRequest) SetShopInfo(_shopInfo string) error {
    r._shopInfo = _shopInfo
    r.Set("shop_info", _shopInfo)
    return nil
}

// ShopInfo Getter
func (r AlibabaAlihealthDrugWxinfoUploadRequest) GetShopInfo() string {
    return r._shopInfo
}
// SalerInfo Setter
// 售货员信息
func (r *AlibabaAlihealthDrugWxinfoUploadRequest) SetSalerInfo(_salerInfo string) error {
    r._salerInfo = _salerInfo
    r.Set("saler_info", _salerInfo)
    return nil
}

// SalerInfo Getter
func (r AlibabaAlihealthDrugWxinfoUploadRequest) GetSalerInfo() string {
    return r._salerInfo
}
// IsvChannel Setter
// 渠道
func (r *AlibabaAlihealthDrugWxinfoUploadRequest) SetIsvChannel(_isvChannel string) error {
    r._isvChannel = _isvChannel
    r.Set("isv_channel", _isvChannel)
    return nil
}

// IsvChannel Getter
func (r AlibabaAlihealthDrugWxinfoUploadRequest) GetIsvChannel() string {
    return r._isvChannel
}