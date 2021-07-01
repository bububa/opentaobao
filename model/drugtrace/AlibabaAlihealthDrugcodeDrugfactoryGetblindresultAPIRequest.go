package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIRequest
获取盲底文件处理结果 API请求
alibaba.alihealth.drugcode.drugfactory.getblindresult

获取盲底文件处理结果 */
type AlibabaAlihealthDrugcodeDrugfactoryGetblindresultAPIRequest struct {
	model.Params
	// 企业id
	_refEntId string
	// 盲底文件名称
	_blindFileName string
}

// New
