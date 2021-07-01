package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugtraceTopYljgListupoutDetailAPIRequest
上游出库单单据明细查询 API请求
alibaba.alihealth.drugtrace.top.yljg.listupout.detail

查询上游出库单明细(带追溯码信息) */
type AlibabaAlihealthDrugtraceTopYljgListupoutDetailAPIRequest struct {
	model.Params
	// 企业id
	_refEntId string
	// 单据编码
	_billCode string
	// 发货企业renEntId
	_fromRefUserId string
	// 收货企业refEntId
	_toRefUserId string
}

// New
