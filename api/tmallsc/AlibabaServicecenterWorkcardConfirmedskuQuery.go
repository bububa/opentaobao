package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// Alibabaservicecenterworkcardconfirmedskuquery 查询确认履行的服务项
// alibaba.servicecenter.workcard.confirmedsku.query
//
// 查询确认履行的服务项
func Alibabaservicecenterworkcardconfirmedskuquery(clt *core.SDKClient, req *tmallsc.AlibabaservicecenterworkcardconfirmedskuqueryAPIRequest, session string) (*tmallsc.AlibabaservicecenterworkcardconfirmedskuqueryAPIResponse, error) {
	var resp tmallsc.AlibabaservicecenterworkcardconfirmedskuqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
