package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
企业下载列表 API请求
alibaba.alihealth.drug.download.entlist

获取企业的下载文件列表
*/
type AlibabaAlihealthDrugDownloadEntlistRequest struct {
    model.Params
    // appKey
    _appKeyN   string
}

// 初始化AlibabaAlihealthDrugDownloadEntlistRequest对象
func NewAlibabaAlihealthDrugDownloadEntlistRequest() *AlibabaAlihealthDrugDownloadEntlistRequest{
    return &AlibabaAlihealthDrugDownloadEntlistRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugDownloadEntlistRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.download.entlist"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugDownloadEntlistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AppKeyN Setter
// appKey
func (r *AlibabaAlihealthDrugDownloadEntlistRequest) SetAppKeyN(_appKeyN string) error {
    r._appKeyN = _appKeyN
    r.Set("app_key_n", _appKeyN)
    return nil
}

// AppKeyN Getter
func (r AlibabaAlihealthDrugDownloadEntlistRequest) GetAppKeyN() string {
    return r._appKeyN
}
