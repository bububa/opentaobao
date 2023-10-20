package tmallfcbox

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallfcbox"
)

// TmallFcboxNotify 丰巢通知接口
// tmall.fcbox.notify
//
// tmax接收丰巢快递通知
func TmallFcboxNotify(clt *core.SDKClient, req *tmallfcbox.TmallFcboxNotifyAPIRequest, resp *tmallfcbox.TmallFcboxNotifyAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
