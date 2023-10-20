package tblogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tblogistics"
)

// TaobaoLogisticsTraceGet 用户根据交易号查询物流流转信息
// taobao.logistics.trace.get
//
// 用户根据交易号查询物流流转信息
func TaobaoLogisticsTraceGet(clt *core.SDKClient, req *tblogistics.TaobaoLogisticsTraceGetAPIRequest, session string) (*tblogistics.TaobaoLogisticsTraceGetAPIResponse, error) {
	var resp tblogistics.TaobaoLogisticsTraceGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
