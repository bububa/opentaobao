package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomeadvisermessagenotice 催促小B发送短信
// alibaba.alihouse.newhome.adviser.message.notice
//
// 催促小B发送短信
func Alibabaalihousenewhomeadvisermessagenotice(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomeadvisermessagenoticeAPIRequest, session string) (*alihouse.AlibabaalihousenewhomeadvisermessagenoticeAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomeadvisermessagenoticeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
