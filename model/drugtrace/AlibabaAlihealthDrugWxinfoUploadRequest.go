package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
小程序数据回传 APIRequest
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

func NewAlibabaAlihealthDrugWxinfoUploadRequest() *AlibabaAlihealthDrugWxinfoUploadRequest{
    return &AlibabaAlihealthDrugWxinfoUploadRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugWxinfoUploadRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.wxinfo.upload"
}

func (r AlibabaAlihealthDrugWxinfoUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugWxinfoUploadRequest) SetUserInfo(userInfo string) error {
    r.userInfo = userInfo
    r.Set("user_info", userInfo)
    return nil
}

func (r AlibabaAlihealthDrugWxinfoUploadRequest) GetUserInfo() string {
    return r.userInfo
}

func (r *AlibabaAlihealthDrugWxinfoUploadRequest) SetShopInfo(shopInfo string) error {
    r.shopInfo = shopInfo
    r.Set("shop_info", shopInfo)
    return nil
}

func (r AlibabaAlihealthDrugWxinfoUploadRequest) GetShopInfo() string {
    return r.shopInfo
}

func (r *AlibabaAlihealthDrugWxinfoUploadRequest) SetSalerInfo(salerInfo string) error {
    r.salerInfo = salerInfo
    r.Set("saler_info", salerInfo)
    return nil
}

func (r AlibabaAlihealthDrugWxinfoUploadRequest) GetSalerInfo() string {
    return r.salerInfo
}

func (r *AlibabaAlihealthDrugWxinfoUploadRequest) SetIsvChannel(isvChannel string) error {
    r.isvChannel = isvChannel
    r.Set("isv_channel", isvChannel)
    return nil
}

func (r AlibabaAlihealthDrugWxinfoUploadRequest) GetIsvChannel() string {
    return r.isvChannel
}

