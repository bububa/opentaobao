package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthDrugKytDrQueryupbillcodeAPIRequest
查询上游企业出库单据号 API请求
alibaba.alihealth.drug.kyt.dr.queryupbillcode

疫苗温度合规补充需求-增加一个查询上游出库单号的接口。疾控在扫码入库时，接口通过扫到的码判定这个码对应所属的出库单据号 */
type AlibabaAlihealthDrugKytDrQueryupbillcodeAPIRequest struct {
	model.Params
	// 追溯码
	_code string
	// 企业ID （一般为要查询单据的收货企业）
	_refEntId string
}

// New
