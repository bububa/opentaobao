package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
企业上传回执 API请求
alibaba.alihealth.drug.download.fileacceptret

拿到企业下载回执，将企业已下载的和未下载成功的条目都相应的改变状态
*/
type AlibabaAlihealthDrugDownloadFileacceptretRequest struct {
    model.Params
    // appKey
    _appKeyN   string
    // fileResultJson
    _fileResultJson   string
}

// 初始化AlibabaAlihealthDrugDownloadFileacceptretRequest对象
func NewAlibabaAlihealthDrugDownloadFileacceptretRequest() *AlibabaAlihealthDrugDownloadFileacceptretRequest{
    return &AlibabaAlihealthDrugDownloadFileacceptretRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugDownloadFileacceptretRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.download.fileacceptret"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugDownloadFileacceptretRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AppKeyN Setter
// appKey
func (r *AlibabaAlihealthDrugDownloadFileacceptretRequest) SetAppKeyN(_appKeyN string) error {
    r._appKeyN = _appKeyN
    r.Set("app_key_n", _appKeyN)
    return nil
}

// AppKeyN Getter
func (r AlibabaAlihealthDrugDownloadFileacceptretRequest) GetAppKeyN() string {
    return r._appKeyN
}
// FileResultJson Setter
// fileResultJson
func (r *AlibabaAlihealthDrugDownloadFileacceptretRequest) SetFileResultJson(_fileResultJson string) error {
    r._fileResultJson = _fileResultJson
    r.Set("file_result_json", _fileResultJson)
    return nil
}

// FileResultJson Getter
func (r AlibabaAlihealthDrugDownloadFileacceptretRequest) GetFileResultJson() string {
    return r._fileResultJson
}
