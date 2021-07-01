package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailAPIRequest
根据单据号查询单据的详情信息【注意：查询的是本企业的单据】 API请求
alibaba.alihealth.drugtrace.top.yljg.query.upbilldetail

根据单据号查询单据的详情信息【注意：这个接口查询的是本企业的单据，如果是查询上游的单据明细信息，使用xxxxxxx.listupout.detail接口】。 */
type AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailAPIRequest struct {
	model.Params
	// 单据号码
	_billCode string
	// 本企业refEntId
	_refEntId string
}

// New
