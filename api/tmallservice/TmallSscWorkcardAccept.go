package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallSscWorkcardAccept 服务商接单完成
// tmall.ssc.workcard.accept
//
// 工单完结
func TmallSscWorkcardAccept(clt *core.SDKClient, req *tmallservice.TmallSscWorkcardAcceptAPIRequest, session string) (*tmallservice.TmallSscWorkcardAcceptAPIResponse, error) {
	var resp tmallservice.TmallSscWorkcardAcceptAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
