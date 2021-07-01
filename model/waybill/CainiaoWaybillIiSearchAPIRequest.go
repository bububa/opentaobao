package waybill

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoWaybillIiSearchAPIRequest
查询面单服务订购及面单使用情况 API请求
cainiao.waybill.ii.search

获取发货地&CP开通状态&账户的使用情况 */
type CainiaoWaybillIiSearchAPIRequest struct {
	model.Params
	// 物流公司code
	_cpCode string
}

// New
