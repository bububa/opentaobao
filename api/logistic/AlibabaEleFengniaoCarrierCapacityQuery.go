package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

/* AlibabaEleFengniaoCarrierCapacityQuery
按照门店查询骑手运力状态查询
alibaba.ele.fengniao.carrier.capacity.query

提供给大润发，用于按照站点纬度查询大润发每个配送站的实时上班骑手数、到店骑手数、活跃骑手数量 */
func AlibabaEleFengniaoCarrierCapacityQuery(clt *core.SDKClient, req *logistic.AlibabaEleFengniaoCarrierCapacityQueryAPIRequest, session string) (*logistic.AlibabaEleFengniaoCarrierCapacityQueryAPIResponse, error) {
	var resp logistic.AlibabaEleFengniaoCarrierCapacityQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
