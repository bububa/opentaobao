package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// Alibabaaliqinfcvoicenumcancelcall 取消呼叫
// alibaba.aliqin.fc.voice.num.cancelcall
//
// 当通话通过阿里大于呼出后可以通过调用这个接口取消本次通话
func Alibabaaliqinfcvoicenumcancelcall(clt *core.SDKClient, req *alicom.AlibabaaliqinfcvoicenumcancelcallAPIRequest, session string) (*alicom.AlibabaaliqinfcvoicenumcancelcallAPIResponse, error) {
	var resp alicom.AlibabaaliqinfcvoicenumcancelcallAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
