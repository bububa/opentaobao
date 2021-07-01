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
type AlibabaAlihealthDrugDownloadGetentauthentAPIRequest struct {
    model.Params
    // 授权开始时间
    _authBeginDate   string
    // 授权结束时间
    _authEndDate   string
}

// 初始化AlibabaAlihealthDrugDownloadGetentauthentAPIRequest对象
func NewAlibabaAlihealthDrugDownloadGetentauthentRequest() *AlibabaAlihealthDrugDownloadGetentauthentAPIRequest{
    return &AlibabaAlihealthDrugDownloadGetentauthentAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugDownloadGetentauthentAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.download.getentauthent"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugDownloadGetentauthentAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AuthBeginDate Setter
// 授权开始时间
func (r *AlibabaAlihealthDrugDownloadGetentauthentAPIRequest) SetAuthBeginDate(_authBeginDate string) error {
    r._authBeginDate = _authBeginDate
    r.Set("auth_begin_date", _authBeginDate)
    return nil
}

// AuthBeginDate Getter
func (r AlibabaAlihealthDrugDownloadGetentauthentAPIRequest) GetAuthBeginDate() string {
    return r._authBeginDate
}
// AuthEndDate Setter
// 授权结束时间
func (r *AlibabaAlihealthDrugDownloadGetentauthentAPIRequest) SetAuthEndDate(_authEndDate string) error {
    r._authEndDate = _authEndDate
    r.Set("auth_end_date", _authEndDate)
    return nil
}

// AuthEndDate Getter
func (r AlibabaAlihealthDrugDownloadGetentauthentAPIRequest) GetAuthEndDate() string {
    return r._authEndDate
}
