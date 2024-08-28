package alicom

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// AlibabaAliqinTaNumberSinglecallbyvoice 根据号码tts单呼
// alibaba.aliqin.ta.number.singlecallbyvoice
//
// 根据号码语音单呼
func AlibabaAliqinTaNumberSinglecallbyvoice(ctx context.Context, clt *core.SDKClient, req *alicom.AlibabaAliqinTaNumberSinglecallbyvoiceAPIRequest, resp *alicom.AlibabaAliqinTaNumberSinglecallbyvoiceAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
