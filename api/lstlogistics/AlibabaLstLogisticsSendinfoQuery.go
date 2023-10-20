package lstlogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstlogistics"
)

// Alibabalstlogisticssendinfoquery 供应商-异云-查询主订单包含的物流单
// alibaba.lst.logistics.sendinfo.query
//
// 查询主订单包含的物流单
func Alibabalstlogisticssendinfoquery(clt *core.SDKClient, req *lstlogistics.AlibabalstlogisticssendinfoqueryAPIRequest, session string) (*lstlogistics.AlibabalstlogisticssendinfoqueryAPIResponse, error) {
	var resp lstlogistics.AlibabalstlogisticssendinfoqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
