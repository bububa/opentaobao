package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// Alibabaaliqintasmsnumsend 短信发送
// alibaba.aliqin.ta.sms.num.send
//
// 短信发送
func Alibabaaliqintasmsnumsend(clt *core.SDKClient, req *alicom.AlibabaaliqintasmsnumsendAPIRequest, session string) (*alicom.AlibabaaliqintasmsnumsendAPIResponse, error) {
	var resp alicom.AlibabaaliqintasmsnumsendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
