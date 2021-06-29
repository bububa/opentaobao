package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
正大鸡蛋激活追溯码 API请求
alibaba.alihealth.tracecodeplatform.code.active

用于正大鸡蛋激活追溯码
*/
type AlibabaAlihealthTracecodeplatformCodeActiveRequest struct {
    model.Params
    // 文件信息（对文件内容16进制编码）
    _fileInfo   string
    // 回调url
    _callbackUrl   string
    // 文件名
    _fileName   string
    // 商品id
    _prodId   int64
}

// 初始化AlibabaAlihealthTracecodeplatformCodeActiveRequest对象
func NewAlibabaAlihealthTracecodeplatformCodeActiveRequest() *AlibabaAlihealthTracecodeplatformCodeActiveRequest{
    return &AlibabaAlihealthTracecodeplatformCodeActiveRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthTracecodeplatformCodeActiveRequest) GetApiMethodName() string {
    return "alibaba.alihealth.tracecodeplatform.code.active"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthTracecodeplatformCodeActiveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// FileInfo Setter
// 文件信息（对文件内容16进制编码）
func (r *AlibabaAlihealthTracecodeplatformCodeActiveRequest) SetFileInfo(_fileInfo string) error {
    r._fileInfo = _fileInfo
    r.Set("file_info", _fileInfo)
    return nil
}

// FileInfo Getter
func (r AlibabaAlihealthTracecodeplatformCodeActiveRequest) GetFileInfo() string {
    return r._fileInfo
}
// CallbackUrl Setter
// 回调url
func (r *AlibabaAlihealthTracecodeplatformCodeActiveRequest) SetCallbackUrl(_callbackUrl string) error {
    r._callbackUrl = _callbackUrl
    r.Set("callback_url", _callbackUrl)
    return nil
}

// CallbackUrl Getter
func (r AlibabaAlihealthTracecodeplatformCodeActiveRequest) GetCallbackUrl() string {
    return r._callbackUrl
}
// FileName Setter
// 文件名
func (r *AlibabaAlihealthTracecodeplatformCodeActiveRequest) SetFileName(_fileName string) error {
    r._fileName = _fileName
    r.Set("file_name", _fileName)
    return nil
}

// FileName Getter
func (r AlibabaAlihealthTracecodeplatformCodeActiveRequest) GetFileName() string {
    return r._fileName
}
// ProdId Setter
// 商品id
func (r *AlibabaAlihealthTracecodeplatformCodeActiveRequest) SetProdId(_prodId int64) error {
    r._prodId = _prodId
    r.Set("prod_id", _prodId)
    return nil
}

// ProdId Getter
func (r AlibabaAlihealthTracecodeplatformCodeActiveRequest) GetProdId() int64 {
    return r._prodId
}