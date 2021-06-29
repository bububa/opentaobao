package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
企业上传回执 APIRequest
alibaba.alihealth.drug.download.fileacceptret

拿到企业下载回执，将企业已下载的和未下载成功的条目都相应的改变状态
*/
type AlibabaAlihealthDrugDownloadFileacceptretRequest struct {
    model.Params

    // appKey
    appKeyN   string 

    // fileResultJson
    fileResultJson   string 

}

func NewAlibabaAlihealthDrugDownloadFileacceptretRequest() *AlibabaAlihealthDrugDownloadFileacceptretRequest{
    return &AlibabaAlihealthDrugDownloadFileacceptretRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugDownloadFileacceptretRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.download.fileacceptret"
}

func (r AlibabaAlihealthDrugDownloadFileacceptretRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugDownloadFileacceptretRequest) SetAppKeyN(appKeyN string) error {
    r.appKeyN = appKeyN
    r.Set("app_key_n", appKeyN)
    return nil
}

func (r AlibabaAlihealthDrugDownloadFileacceptretRequest) GetAppKeyN() string {
    return r.appKeyN
}

func (r *AlibabaAlihealthDrugDownloadFileacceptretRequest) SetFileResultJson(fileResultJson string) error {
    r.fileResultJson = fileResultJson
    r.Set("file_result_json", fileResultJson)
    return nil
}

func (r AlibabaAlihealthDrugDownloadFileacceptretRequest) GetFileResultJson() string {
    return r.fileResultJson
}

