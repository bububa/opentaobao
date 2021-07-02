package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthTracecodesellerCodeActiveAPIRequest 上传激活码的文件 API请求
// alibaba.alihealth.tracecodeseller.code.active
//
// 上传商品的激活码文件，存到系统中
type AlibabaAlihealthTracecodesellerCodeActiveAPIRequest struct {
	model.Params
	// 文件名
	_fileName string
	// 商品编号
	_productInfoId int64
	// 文件内容，十六进制编码
	_fileContent string
	// 关联类型，0:无关联，1:前关联，2:后关联
	_correlationType int64
	// 关联比例
	_correlationRatio string
	// 语言标识
	_language string
}

// NewAlibabaAlihealthTracecodesellerCodeActiveRequest 初始化AlibabaAlihealthTracecodesellerCodeActiveAPIRequest对象
func NewAlibabaAlihealthTracecodesellerCodeActiveRequest() *AlibabaAlihealthTracecodesellerCodeActiveAPIRequest {
	return &AlibabaAlihealthTracecodesellerCodeActiveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthTracecodesellerCodeActiveAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.tracecodeseller.code.active"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthTracecodesellerCodeActiveAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is FileName Setter
// 文件名
func (r *AlibabaAlihealthTracecodesellerCodeActiveAPIRequest) SetFileName(_fileName string) error {
	r._fileName = _fileName
	r.Set("file_name", _fileName)
	return nil
}

// Get FileName Getter
func (r AlibabaAlihealthTracecodesellerCodeActiveAPIRequest) GetFileName() string {
	return r._fileName
}

// Set is ProductInfoId Setter
// 商品编号
func (r *AlibabaAlihealthTracecodesellerCodeActiveAPIRequest) SetProductInfoId(_productInfoId int64) error {
	r._productInfoId = _productInfoId
	r.Set("product_info_id", _productInfoId)
	return nil
}

// Get ProductInfoId Getter
func (r AlibabaAlihealthTracecodesellerCodeActiveAPIRequest) GetProductInfoId() int64 {
	return r._productInfoId
}

// Set is FileContent Setter
// 文件内容，十六进制编码
func (r *AlibabaAlihealthTracecodesellerCodeActiveAPIRequest) SetFileContent(_fileContent string) error {
	r._fileContent = _fileContent
	r.Set("file_content", _fileContent)
	return nil
}

// Get FileContent Getter
func (r AlibabaAlihealthTracecodesellerCodeActiveAPIRequest) GetFileContent() string {
	return r._fileContent
}

// Set is CorrelationType Setter
// 关联类型，0:无关联，1:前关联，2:后关联
func (r *AlibabaAlihealthTracecodesellerCodeActiveAPIRequest) SetCorrelationType(_correlationType int64) error {
	r._correlationType = _correlationType
	r.Set("correlation_type", _correlationType)
	return nil
}

// Get CorrelationType Getter
func (r AlibabaAlihealthTracecodesellerCodeActiveAPIRequest) GetCorrelationType() int64 {
	return r._correlationType
}

// Set is CorrelationRatio Setter
// 关联比例
func (r *AlibabaAlihealthTracecodesellerCodeActiveAPIRequest) SetCorrelationRatio(_correlationRatio string) error {
	r._correlationRatio = _correlationRatio
	r.Set("correlation_ratio", _correlationRatio)
	return nil
}

// Get CorrelationRatio Getter
func (r AlibabaAlihealthTracecodesellerCodeActiveAPIRequest) GetCorrelationRatio() string {
	return r._correlationRatio
}

// Set is Language Setter
// 语言标识
func (r *AlibabaAlihealthTracecodesellerCodeActiveAPIRequest) SetLanguage(_language string) error {
	r._language = _language
	r.Set("language", _language)
	return nil
}

// Get Language Getter
func (r AlibabaAlihealthTracecodesellerCodeActiveAPIRequest) GetLanguage() string {
	return r._language
}
