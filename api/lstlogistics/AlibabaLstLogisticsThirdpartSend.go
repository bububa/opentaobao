package lstlogistics

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstlogistics"
)

// Alibabalstlogisticsthirdpartsend 供应商-异云-使用第三方物流发货
// alibaba.lst.logistics.thirdpart.send
//
// 异地云仓的订单，使用第三方物流的发货方式来变更订单发货状态
func Alibabalstlogisticsthirdpartsend(clt *core.SDKClient, req *lstlogistics.AlibabalstlogisticsthirdpartsendAPIRequest, session string) (*lstlogistics.AlibabalstlogisticsthirdpartsendAPIResponse, error) {
	var resp lstlogistics.AlibabalstlogisticsthirdpartsendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
