package aliqin

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aliqin"
)

// Alibabaaliqinfcsmsnumquery 短信发送记录查询
// alibaba.aliqin.fc.sms.num.query
//
// 短信发送记录查询。
func Alibabaaliqinfcsmsnumquery(clt *core.SDKClient, req *aliqin.AlibabaaliqinfcsmsnumqueryAPIRequest, session string) (*aliqin.AlibabaaliqinfcsmsnumqueryAPIResponse, error) {
	var resp aliqin.AlibabaaliqinfcsmsnumqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
