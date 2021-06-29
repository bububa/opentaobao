package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
上传激活码的文件 APIRequest
alibaba.alihealth.tracecodeseller.code.active

上传商品的激活码文件，存到系统中
*/
type AlibabaAlihealthTracecodesellerCodeActiveRequest struct {
    model.Params

    // 文件名
    fileName   string 

    // 商品编号
    productInfoId   int64 

    // 文件内容，十六进制编码
    fileContent   string 

    // 关联类型，0:无关联，1:前关联，2:后关联
    correlationType   int64 

    // 关联比例
    correlationRatio   string 

    // 语言标识
    language   string 

}

func NewAlibabaAlihealthTracecodesellerCodeActiveRequest() *AlibabaAlihealthTracecodesellerCodeActiveRequest{
    return &AlibabaAlihealthTracecodesellerCodeActiveRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthTracecodesellerCodeActiveRequest) GetApiMethodName() string {
    return "alibaba.alihealth.tracecodeseller.code.active"
}

func (r AlibabaAlihealthTracecodesellerCodeActiveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthTracecodesellerCodeActiveRequest) SetFileName(fileName string) error {
    r.fileName = fileName
    r.Set("file_name", fileName)
    return nil
}

func (r AlibabaAlihealthTracecodesellerCodeActiveRequest) GetFileName() string {
    return r.fileName
}

func (r *AlibabaAlihealthTracecodesellerCodeActiveRequest) SetProductInfoId(productInfoId int64) error {
    r.productInfoId = productInfoId
    r.Set("product_info_id", productInfoId)
    return nil
}

func (r AlibabaAlihealthTracecodesellerCodeActiveRequest) GetProductInfoId() int64 {
    return r.productInfoId
}

func (r *AlibabaAlihealthTracecodesellerCodeActiveRequest) SetFileContent(fileContent string) error {
    r.fileContent = fileContent
    r.Set("file_content", fileContent)
    return nil
}

func (r AlibabaAlihealthTracecodesellerCodeActiveRequest) GetFileContent() string {
    return r.fileContent
}

func (r *AlibabaAlihealthTracecodesellerCodeActiveRequest) SetCorrelationType(correlationType int64) error {
    r.correlationType = correlationType
    r.Set("correlation_type", correlationType)
    return nil
}

func (r AlibabaAlihealthTracecodesellerCodeActiveRequest) GetCorrelationType() int64 {
    return r.correlationType
}

func (r *AlibabaAlihealthTracecodesellerCodeActiveRequest) SetCorrelationRatio(correlationRatio string) error {
    r.correlationRatio = correlationRatio
    r.Set("correlation_ratio", correlationRatio)
    return nil
}

func (r AlibabaAlihealthTracecodesellerCodeActiveRequest) GetCorrelationRatio() string {
    return r.correlationRatio
}

func (r *AlibabaAlihealthTracecodesellerCodeActiveRequest) SetLanguage(language string) error {
    r.language = language
    r.Set("language", language)
    return nil
}

func (r AlibabaAlihealthTracecodesellerCodeActiveRequest) GetLanguage() string {
    return r.language
}

