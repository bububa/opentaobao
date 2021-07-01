package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytStorebilldeleteAPIRequest
零售端单据删除 API请求
alibaba.alihealth.drug.kyt.storebilldelete

零售端单据删除 */
type AlibabaAlihealthDrugKytStorebilldeleteAPIRequest struct {
	model.Params
	// 企业ID
	_refEntId string
	// 操作人编码
	_icCode string
	// 单据ID
	_billId string
	// 单据类型
	_billType string
}

// New
