package aliqin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliqin"
)

// Alibabaaliqinfcivrnumcall ivr呼叫
// alibaba.aliqin.fc.ivr.num.call
//
// ivr呼叫
func Alibabaaliqinfcivrnumcall(clt *core.SDKClient, req *aliqin.AlibabaaliqinfcivrnumcallAPIRequest, session string) (*aliqin.AlibabaaliqinfcivrnumcallAPIResponse, error) {
	var resp aliqin.AlibabaaliqinfcivrnumcallAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
