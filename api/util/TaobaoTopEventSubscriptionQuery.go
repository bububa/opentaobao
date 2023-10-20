package util

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// Taobaotopeventsubscriptionquery 三方事件订阅查询
// taobao.top.event.subscription.query
//
// 三方事件订阅查询
func Taobaotopeventsubscriptionquery(clt *core.SDKClient, req *util.TaobaotopeventsubscriptionqueryAPIRequest, session string) (*util.TaobaotopeventsubscriptionqueryAPIResponse, error) {
	var resp util.TaobaotopeventsubscriptionqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
