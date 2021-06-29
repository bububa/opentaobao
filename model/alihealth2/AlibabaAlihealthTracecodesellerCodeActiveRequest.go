package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
上传激活码的文件 API请求
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

// 初始化AlibabaAlihealthTracecodesellerCodeActiveRequest对象
func NewAlibabaAlihealthTracecodesellerCodeActiveRequest() *AlibabaAlihealthTracecodesellerCodeActiveRequest{
    return &AlibabaAlihealthTracecodesellerCodeActiveRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthTracecodesellerCodeActiveRequest) GetApiMethodName() string {
    return "alibaba.alihealth.tracecodeseller.code.active"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthTracecodesellerCodeActiveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// FileName Setter
// 文件名
func (r *AlibabaAlihealthTracecodesellerCodeActiveRequest) SetFileName(fileName string) error {
    r.fileName = fileName
    r.Set("file_name", fileName)
    return nil
}

// FileName Getter
func (r AlibabaAlihealthTracecodesellerCodeActiveRequest) GetFileName() string {
    return r.fileName
}
// ProductInfoId Setter
// 商品编号
func (r *AlibabaAlihealthTracecodesellerCodeActiveRequest) SetProductInfoId(productInfoId int64) error {
    r.productInfoId = productInfoId
    r.Set("product_info_id", productInfoId)
    return nil
}

// ProductInfoId Getter
func (r AlibabaAlihealthTracecodesellerCodeActiveRequest) GetProductInfoId() int64 {
    return r.productInfoId
}
// FileContent Setter
// 文件内容，十六进制编码
func (r *AlibabaAlihealthTracecodesellerCodeActiveRequest) SetFileContent(fileContent string) error {
    r.fileContent = fileContent
    r.Set("file_content", fileContent)
    return nil
}

// FileContent Getter
func (r AlibabaAlihealthTracecodesellerCodeActiveRequest) GetFileContent() string {
    return r.fileContent
}
// CorrelationType Setter
// 关联类型，0:无关联，1:前关联，2:后关联
func (r *AlibabaAlihealthTracecodesellerCodeActiveRequest) SetCorrelationType(correlationType int64) error {
    r.correlationType = correlationType
    r.Set("correlation_type", correlationType)
    return nil
}

// CorrelationType Getter
func (r AlibabaAlihealthTracecodesellerCodeActiveRequest) GetCorrelationType() int64 {
    return r.correlationType
}
// CorrelationRatio Setter
// 关联比例
func (r *AlibabaAlihealthTracecodesellerCodeActiveRequest) SetCorrelationRatio(correlationRatio string) error {
    r.correlationRatio = correlationRatio
    r.Set("correlation_ratio", correlationRatio)
    return nil
}

// CorrelationRatio Getter
func (r AlibabaAlihealthTracecodesellerCodeActiveRequest) GetCorrelationRatio() string {
    return r.correlationRatio
}
// Language Setter
// 语言标识
func (r *AlibabaAlihealthTracecodesellerCodeActiveRequest) SetLanguage(language string) error {
    r.language = language
    r.Set("language", language)
    return nil
}

// Language Getter
func (r AlibabaAlihealthTracecodesellerCodeActiveRequest) GetLanguage() string {
    return r.language
}
