package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
正大鸡蛋激活追溯码 APIRequest
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

func NewAlibabaAlihealthTracecodeplatformCodeActiveRequest() *AlibabaAlihealthTracecodeplatformCodeActiveRequest{
    return &AlibabaAlihealthTracecodeplatformCodeActiveRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthTracecodeplatformCodeActiveRequest) GetApiMethodName() string {
    return "alibaba.alihealth.tracecodeplatform.code.active"
}

func (r AlibabaAlihealthTracecodeplatformCodeActiveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthTracecodeplatformCodeActiveRequest) SetFileInfo(fileInfo string) error {
    r.fileInfo = fileInfo
    r.Set("file_info", fileInfo)
    return nil
}

func (r AlibabaAlihealthTracecodeplatformCodeActiveRequest) GetFileInfo() string {
    return r.fileInfo
}

func (r *AlibabaAlihealthTracecodeplatformCodeActiveRequest) SetCallbackUrl(callbackUrl string) error {
    r.callbackUrl = callbackUrl
    r.Set("callback_url", callbackUrl)
    return nil
}

func (r AlibabaAlihealthTracecodeplatformCodeActiveRequest) GetCallbackUrl() string {
    return r.callbackUrl
}

func (r *AlibabaAlihealthTracecodeplatformCodeActiveRequest) SetFileName(fileName string) error {
    r.fileName = fileName
    r.Set("file_name", fileName)
    return nil
}

func (r AlibabaAlihealthTracecodeplatformCodeActiveRequest) GetFileName() string {
    return r.fileName
}

func (r *AlibabaAlihealthTracecodeplatformCodeActiveRequest) SetProdId(prodId int64) error {
    r.prodId = prodId
    r.Set("prod_id", prodId)
    return nil
}

func (r AlibabaAlihealthTracecodeplatformCodeActiveRequest) GetProdId() int64 {
    return r.prodId
}

