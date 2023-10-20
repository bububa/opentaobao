package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// Alibabaaliqintasmsnumquery 短信查询
// alibaba.aliqin.ta.sms.num.query
//
// 查询短信发送揭露
func Alibabaaliqintasmsnumquery(clt *core.SDKClient, req *alicom.AlibabaaliqintasmsnumqueryAPIRequest, session string) (*alicom.AlibabaaliqintasmsnumqueryAPIResponse, error) {
	var resp alicom.AlibabaaliqintasmsnumqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
