package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取授权企业列表 APIRequest
alibaba.alihealth.drug.download.getentauthent

D2D数据落地获取授权企业列表
*/
type AlibabaAlihealthDrugDownloadGetentauthentRequest struct {
    model.Params

    // 授权开始时间
    authBeginDate   string 

    // 授权结束时间
    authEndDate   string 

}

func NewAlibabaAlihealthDrugDownloadGetentauthentRequest() *AlibabaAlihealthDrugDownloadGetentauthentRequest{
    return &AlibabaAlihealthDrugDownloadGetentauthentRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugDownloadGetentauthentRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.download.getentauthent"
}

func (r AlibabaAlihealthDrugDownloadGetentauthentRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugDownloadGetentauthentRequest) SetAuthBeginDate(authBeginDate string) error {
    r.authBeginDate = authBeginDate
    r.Set("auth_begin_date", authBeginDate)
    return nil
}

func (r AlibabaAlihealthDrugDownloadGetentauthentRequest) GetAuthBeginDate() string {
    return r.authBeginDate
}

func (r *AlibabaAlihealthDrugDownloadGetentauthentRequest) SetAuthEndDate(authEndDate string) error {
    r.authEndDate = authEndDate
    r.Set("auth_end_date", authEndDate)
    return nil
}

func (r AlibabaAlihealthDrugDownloadGetentauthentRequest) GetAuthEndDate() string {
    return r.authEndDate
}

