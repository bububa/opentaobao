package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoAPIRequest
通过企业名得到唯一标识ref_ent_id及企业ent_id API请求
alibaba.alihealth.drugtrace.top.yljg.query.getentinfo

根据企业名称查询ID */
type AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoAPIRequest struct {
	model.Params
	// 公司名称(全称)
	_entName string
}

// New
