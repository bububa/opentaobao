package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallSscWorkcardAccept 服务商接单完成
// tmall.ssc.workcard.accept
//
// 工单完结
func TmallSscWorkcardAccept(clt *core.SDKClient, req *tmallservice.TmallSscWorkcardAcceptAPIRequest, resp *tmallservice.TmallSscWorkcardAcceptAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
