package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytDrGetentinfoAPIRequest
通过企业名得到唯一标识（ref_ent_id）及企业ID(ent_id) API请求
alibaba.alihealth.drug.kyt.dr.getentinfo

根据企业名称查询ID */
type AlibabaAlihealthDrugKytDrGetentinfoAPIRequest struct {
	model.Params
	// 公司名称(全称)
	_entName string
}

// New
