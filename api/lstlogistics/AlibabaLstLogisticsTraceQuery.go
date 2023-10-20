package lstlogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstlogistics"
)

// Alibabalstlogisticstracequery 供应商-异云-查询运单物流追踪信息
// alibaba.lst.logistics.trace.query
//
// 查询LP单物流追踪信息
func Alibabalstlogisticstracequery(clt *core.SDKClient, req *lstlogistics.AlibabalstlogisticstracequeryAPIRequest, session string) (*lstlogistics.AlibabalstlogisticstracequeryAPIResponse, error) {
	var resp lstlogistics.AlibabalstlogisticstracequeryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
