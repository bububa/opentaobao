package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// Alibabaaliqintanumbersinglecallbyvoice 根据号码tts单呼
// alibaba.aliqin.ta.number.singlecallbyvoice
//
// 根据号码语音单呼
func Alibabaaliqintanumbersinglecallbyvoice(clt *core.SDKClient, req *alicom.AlibabaaliqintanumbersinglecallbyvoiceAPIRequest, session string) (*alicom.AlibabaaliqintanumbersinglecallbyvoiceAPIResponse, error) {
	var resp alicom.AlibabaaliqintanumbersinglecallbyvoiceAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
