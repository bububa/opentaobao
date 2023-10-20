package xiami

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xiami"
)

// AlibabaXiamiApiRadioMyselfGet 我的电台
// alibaba.xiami.api.radio.myself.get
//
// 我的电台
func AlibabaXiamiApiRadioMyselfGet(clt *core.SDKClient, req *xiami.AlibabaXiamiApiRadioMyselfGetAPIRequest, resp *xiami.AlibabaXiamiApiRadioMyselfGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
