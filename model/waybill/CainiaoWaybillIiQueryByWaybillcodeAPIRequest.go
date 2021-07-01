package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoWaybillIiQueryByWaybillcodeAPIRequest
通过面单号查询面单信息 API请求
cainiao.waybill.ii.query.by.waybillcode

通过面单号查看面单号的当前状态，如签收、发货、失效等。 */
type CainiaoWaybillIiQueryByWaybillcodeAPIRequest struct {
	model.Params
	// 系统自动生成
	_paramList []WaybillDetailQueryByWaybillCodeRequest
}

// New
