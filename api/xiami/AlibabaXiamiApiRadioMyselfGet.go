package xiami

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xiami"
)

// AlibabaXiamiApiRadioMyselfGet 我的电台
// alibaba.xiami.api.radio.myself.get
//
// 我的电台
func AlibabaXiamiApiRadioMyselfGet(ctx context.Context, clt *core.SDKClient, req *xiami.AlibabaXiamiApiRadioMyselfGetAPIRequest, resp *xiami.AlibabaXiamiApiRadioMyselfGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
