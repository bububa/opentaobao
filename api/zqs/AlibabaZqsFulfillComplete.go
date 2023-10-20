package zqs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/zqs"
)

// AlibabaZqsFulfillComplete 周期购履约完成接口
// alibaba.zqs.fulfill.complete
//
// 周期购履约完成接口
func AlibabaZqsFulfillComplete(clt *core.SDKClient, req *zqs.AlibabaZqsFulfillCompleteAPIRequest, resp *zqs.AlibabaZqsFulfillCompleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
