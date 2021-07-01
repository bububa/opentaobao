package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytDrSinglerelationAPIRequest
多融单码关联关系查询 API请求
alibaba.alihealth.drug.kyt.dr.singlerelation

单码关联关系查询 */
type AlibabaAlihealthDrugKytDrSinglerelationAPIRequest struct {
	model.Params
	// 追溯码
	_code string
	// 接口调用企业的唯一标识（接口调用者）
	_refEntId string
	// 目标企业唯一标识（为哪个企业查询，一般与入参ref_ent_id一样）
	_desRefEntId string
}

// New
