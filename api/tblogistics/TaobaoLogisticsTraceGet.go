package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// TaobaoLogisticsTraceGet 用户根据交易号查询物流流转信息
// taobao.logistics.trace.get
//
// 用户根据交易号查询物流流转信息
func TaobaoLogisticsTraceGet(clt *core.SDKClient, req *tblogistics.TaobaoLogisticsTraceGetAPIRequest, resp *tblogistics.TaobaoLogisticsTraceGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
