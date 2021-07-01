package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthTracecodesellerCodeActiveAPIRequest
上传激活码的文件 API请求
alibaba.alihealth.tracecodeseller.code.active

上传商品的激活码文件，存到系统中 */
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

// New
