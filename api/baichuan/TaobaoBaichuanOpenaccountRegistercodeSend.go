package baichuan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

/* TaobaoBaichuanOpenaccountRegistercodeSend
百川发送注册验证码
taobao.baichuan.openaccount.registercode.send

百川发送注册验证码 */
func TaobaoBaichuanOpenaccountRegistercodeSend(clt *core.SDKClient, req *baichuan.TaobaoBaichuanOpenaccountRegistercodeSendAPIRequest, session string) (*baichuan.TaobaoBaichuanOpenaccountRegistercodeSendAPIResponse, error) {
	var resp baichuan.TaobaoBaichuanOpenaccountRegistercodeSendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
