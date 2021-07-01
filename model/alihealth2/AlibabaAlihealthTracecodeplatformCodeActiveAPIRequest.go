package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthTracecodeplatformCodeActiveAPIRequest
正大鸡蛋激活追溯码 API请求
alibaba.alihealth.tracecodeplatform.code.active

用于正大鸡蛋激活追溯码 */
type AlibabaAlihealthTracecodeplatformCodeActiveAPIRequest struct {
	model.Params
	// 文件信息（对文件内容16进制编码）
	_fileInfo string
	// 回调url
	_callbackUrl string
	// 文件名
	_fileName string
	// 商品id
	_prodId int64
}

// New
