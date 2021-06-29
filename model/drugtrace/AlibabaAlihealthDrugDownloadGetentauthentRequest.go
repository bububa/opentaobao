package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取授权企业列表 API请求
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

// 初始化AlibabaAlihealthDrugDownloadGetentauthentRequest对象
func NewAlibabaAlihealthDrugDownloadGetentauthentRequest() *AlibabaAlihealthDrugDownloadGetentauthentRequest{
    return &AlibabaAlihealthDrugDownloadGetentauthentRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugDownloadGetentauthentRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.download.getentauthent"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugDownloadGetentauthentRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AuthBeginDate Setter
// 授权开始时间
func (r *AlibabaAlihealthDrugDownloadGetentauthentRequest) SetAuthBeginDate(authBeginDate string) error {
    r.authBeginDate = authBeginDate
    r.Set("auth_begin_date", authBeginDate)
    return nil
}

// AuthBeginDate Getter
func (r AlibabaAlihealthDrugDownloadGetentauthentRequest) GetAuthBeginDate() string {
    return r.authBeginDate
}
// AuthEndDate Setter
// 授权结束时间
func (r *AlibabaAlihealthDrugDownloadGetentauthentRequest) SetAuthEndDate(authEndDate string) error {
    r.authEndDate = authEndDate
    r.Set("auth_end_date", authEndDate)
    return nil
}

// AuthEndDate Getter
func (r AlibabaAlihealthDrugDownloadGetentauthentRequest) GetAuthEndDate() string {
    return r.authEndDate
}
