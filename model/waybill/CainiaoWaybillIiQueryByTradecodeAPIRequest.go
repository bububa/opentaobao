package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoWaybillIiQueryByTradecodeAPIRequest
通过订单号查询电子面单通接口 API请求
cainiao.waybill.ii.query.by.tradecode

通过订单号查看面单的信息 */
type CainiaoWaybillIiQueryByTradecodeAPIRequest struct {
	model.Params
	// 订单号列表
	_paramList []WaybillDetailQueryByBizSubCodeRequest
}

// New
