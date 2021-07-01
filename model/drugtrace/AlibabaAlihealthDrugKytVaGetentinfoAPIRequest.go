package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytVaGetentinfoAPIRequest
通过企业名得到唯一标识（ref_ent_id）及企业ID(ent_id) API请求
alibaba.alihealth.drug.kyt.va.getentinfo

根据企业名称查询企业唯一标识（ref_ent_id）及企业ID(ent_id) */
type AlibabaAlihealthDrugKytVaGetentinfoAPIRequest struct {
	model.Params
	// 公司名称(全称)
	_entName string
}

// New
