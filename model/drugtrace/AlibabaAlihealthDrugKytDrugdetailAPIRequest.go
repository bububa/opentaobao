package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytDrugdetailAPIRequest
查询药品详细信息 API请求
alibaba.alihealth.drug.kyt.drugdetail

查询药品详细信息 */
type AlibabaAlihealthDrugKytDrugdetailAPIRequest struct {
	model.Params
	// 企业ID
	_refEntId string
	// 药品ID
	_drugEntBaseInfoId string
}

// New
