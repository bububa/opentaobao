package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoWaybillIiProductAPIRequest
商家查询物流商产品类型接口 API请求
cainiao.waybill.ii.product

商家可以查询物流商的产品类型和服务能力。 */
type CainiaoWaybillIiProductAPIRequest struct {
	model.Params
	// 快递公司code
	_cpCode string
}

// New
