package baichuan

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/baichuan"
)

// Taobaobaichuanuserloginbytoken 百川手淘信任登录
// taobao.baichuan.user.loginbytoken
//
// 百川手淘信任登录
func Taobaobaichuanuserloginbytoken(clt *core.SDKClient, req *baichuan.TaobaobaichuanuserloginbytokenAPIRequest, session string) (*baichuan.TaobaobaichuanuserloginbytokenAPIResponse, error) {
	var resp baichuan.TaobaobaichuanuserloginbytokenAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
