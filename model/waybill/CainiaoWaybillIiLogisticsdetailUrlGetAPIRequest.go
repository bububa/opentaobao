package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoWaybillIiLogisticsdetailUrlGetAPIRequest
电子面单物流详情授权url获取 API请求
cainiao.waybill.ii.logisticsdetail.url.get

获取电子面单物流详情授权访问的H5 url */
type CainiaoWaybillIiLogisticsdetailUrlGetAPIRequest struct {
	model.Params
	// 快递公司编码
	_cpCode string
	// 电子面单单号
	_waybillCode string
}

// New
