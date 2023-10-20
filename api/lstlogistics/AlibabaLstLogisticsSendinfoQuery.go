package lstlogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstlogistics"
)

// AlibabaLstLogisticsSendinfoQuery 供应商-异云-查询主订单包含的物流单
// alibaba.lst.logistics.sendinfo.query
//
// 查询主订单包含的物流单
func AlibabaLstLogisticsSendinfoQuery(clt *core.SDKClient, req *lstlogistics.AlibabaLstLogisticsSendinfoQueryAPIRequest, resp *lstlogistics.AlibabaLstLogisticsSendinfoQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
