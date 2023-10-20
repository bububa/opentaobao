package tmallfcbox

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallfcbox"
)

// Tmallfcboxnotify 丰巢通知接口
// tmall.fcbox.notify
//
// tmax接收丰巢快递通知
func Tmallfcboxnotify(clt *core.SDKClient, req *tmallfcbox.TmallfcboxnotifyAPIRequest, session string) (*tmallfcbox.TmallfcboxnotifyAPIResponse, error) {
	var resp tmallfcbox.TmallfcboxnotifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
