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
    _fileName   string
    // 商品编号
    _productInfoId   int64
    // 文件内容，十六进制编码
    _fileContent   string
    // 关联类型，0:无关联，1:前关联，2:后关联
    _correlationType   int64
    // 关联比例
    _correlationRatio   string
    // 语言标识
    _language   string
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
func (r *AlibabaAlihealthTracecodesellerCodeActiveRequest) SetFileName(_fileName string) error {
    r._fileName = _fileName
    r.Set("file_name", _fileName)
    return nil
}

// FileName Getter
func (r AlibabaAlihealthTracecodesellerCodeActiveRequest) GetFileName() string {
    return r._fileName
}
// ProductInfoId Setter
// 商品编号
func (r *AlibabaAlihealthTracecodesellerCodeActiveRequest) SetProductInfoId(_productInfoId int64) error {
    r._productInfoId = _productInfoId
    r.Set("product_info_id", _productInfoId)
    return nil
}

// ProductInfoId Getter
func (r AlibabaAlihealthTracecodesellerCodeActiveRequest) GetProductInfoId() int64 {
    return r._productInfoId
}
// FileContent Setter
// 文件内容，十六进制编码
func (r *AlibabaAlihealthTracecodesellerCodeActiveRequest) SetFileContent(_fileContent string) error {
    r._fileContent = _fileContent
    r.Set("file_content", _fileContent)
    return nil
}

// FileContent Getter
func (r AlibabaAlihealthTracecodesellerCodeActiveRequest) GetFileContent() string {
    return r._fileContent
}
// CorrelationType Setter
// 关联类型，0:无关联，1:前关联，2:后关联
func (r *AlibabaAlihealthTracecodesellerCodeActiveRequest) SetCorrelationType(_correlationType int64) error {
    r._correlationType = _correlationType
    r.Set("correlation_type", _correlationType)
    return nil
}

// CorrelationType Getter
func (r AlibabaAlihealthTracecodesellerCodeActiveRequest) GetCorrelationType() int64 {
    return r._correlationType
}
// CorrelationRatio Setter
// 关联比例
func (r *AlibabaAlihealthTracecodesellerCodeActiveRequest) SetCorrelationRatio(_correlationRatio string) error {
    r._correlationRatio = _correlationRatio
    r.Set("correlation_ratio", _correlationRatio)
    return nil
}

// CorrelationRatio Getter
func (r AlibabaAlihealthTracecodesellerCodeActiveRequest) GetCorrelationRatio() string {
    return r._correlationRatio
}
// Language Setter
// 语言标识
func (r *AlibabaAlihealthTracecodesellerCodeActiveRequest) SetLanguage(_language string) error {
    r._language = _language
    r.Set("language", _language)
    return nil
}

// Language Getter
func (r AlibabaAlihealthTracecodesellerCodeActiveRequest) GetLanguage() string {
    return r._language
}