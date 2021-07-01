package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytYyGetentinfoAPIRequest
得到企业信息 API请求
alibaba.alihealth.drug.kyt.yy.getentinfo

根据企业名称查询企业唯一标识【ref_ent_id】和企业ID【ent_id】 */
type AlibabaAlihealthDrugKytYyGetentinfoAPIRequest struct {
	model.Params
	// 公司名称
	_entName string
}

// New
