package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEleFengniaoCarrierCapacityQueryAPIRequest
按照门店查询骑手运力状态查询 API请求
alibaba.ele.fengniao.carrier.capacity.query

提供给大润发，用于按照站点纬度查询大润发每个配送站的实时上班骑手数、到店骑手数、活跃骑手数量 */
type AlibabaEleFengniaoCarrierCapacityQueryAPIRequest struct {
	model.Params
	// 系统自动生成
	_param *Param
}

// New
