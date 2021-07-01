package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIRequest
接收盲底文件删除日志 API请求
alibaba.alihealth.drugcode.drugfactory.blindfiledellog

临床用药试验-接收盲底文件删除日志 */
type AlibabaAlihealthDrugcodeDrugfactoryBlindfiledellogAPIRequest struct {
	model.Params
	// 药厂企业id
	_refEntId string
	// 操作人
	_operator string
	// 盲底文件删除时间
	_blindFileDeleteTime string
	// 盲底文件名称，多个盲底文件用,分隔
	_blindFileNames string
}

// New
