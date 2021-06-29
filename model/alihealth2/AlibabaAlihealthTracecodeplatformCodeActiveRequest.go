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
    fileInfo   string
    // 回调url
    callbackUrl   string
    // 文件名
    fileName   string
    // 商品id
    prodId   int64
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
func (r *AlibabaAlihealthTracecodeplatformCodeActiveRequest) SetFileInfo(fileInfo string) error {
    r.fileInfo = fileInfo
    r.Set("file_info", fileInfo)
    return nil
}

// FileInfo Getter
func (r AlibabaAlihealthTracecodeplatformCodeActiveRequest) GetFileInfo() string {
    return r.fileInfo
}
// CallbackUrl Setter
// 回调url
func (r *AlibabaAlihealthTracecodeplatformCodeActiveRequest) SetCallbackUrl(callbackUrl string) error {
    r.callbackUrl = callbackUrl
    r.Set("callback_url", callbackUrl)
    return nil
}

// CallbackUrl Getter
func (r AlibabaAlihealthTracecodeplatformCodeActiveRequest) GetCallbackUrl() string {
    return r.callbackUrl
}
// FileName Setter
// 文件名
func (r *AlibabaAlihealthTracecodeplatformCodeActiveRequest) SetFileName(fileName string) error {
    r.fileName = fileName
    r.Set("file_name", fileName)
    return nil
}

// FileName Getter
func (r AlibabaAlihealthTracecodeplatformCodeActiveRequest) GetFileName() string {
    return r.fileName
}
// ProdId Setter
// 商品id
func (r *AlibabaAlihealthTracecodeplatformCodeActiveRequest) SetProdId(prodId int64) error {
    r.prodId = prodId
    r.Set("prod_id", prodId)
    return nil
}

// ProdId Getter
func (r AlibabaAlihealthTracecodeplatformCodeActiveRequest) GetProdId() int64 {
    return r.prodId
}
