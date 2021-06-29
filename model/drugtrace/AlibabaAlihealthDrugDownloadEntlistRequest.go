package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
企业下载列表 APIRequest
alibaba.alihealth.drug.download.entlist

获取企业的下载文件列表
*/
type AlibabaAlihealthDrugDownloadEntlistRequest struct {
    model.Params

    // appKey
    appKeyN   string 

}

func NewAlibabaAlihealthDrugDownloadEntlistRequest() *AlibabaAlihealthDrugDownloadEntlistRequest{
    return &AlibabaAlihealthDrugDownloadEntlistRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugDownloadEntlistRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.download.entlist"
}

func (r AlibabaAlihealthDrugDownloadEntlistRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugDownloadEntlistRequest) SetAppKeyN(appKeyN string) error {
    r.appKeyN = appKeyN
    r.Set("app_key_n", appKeyN)
    return nil
}

func (r AlibabaAlihealthDrugDownloadEntlistRequest) GetAppKeyN() string {
    return r.appKeyN
}

