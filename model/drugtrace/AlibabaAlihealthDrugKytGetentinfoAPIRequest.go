package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytGetentinfoAPIRequest
根据企业名称查询企业唯一标识【ref_ent_id】和企业ID【ent_id】 API请求
alibaba.alihealth.drug.kyt.getentinfo

根据企业名称查询企业唯一标识【ref_ent_id】和企业ID【ent_id】 */
type AlibabaAlihealthDrugKytGetentinfoAPIRequest struct {
	model.Params
	// 公司名称
	_entName string
}

// New
