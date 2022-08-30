package topoaid

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/topoaid"
)

// TaobaoTopOaidClientDecrypt 端侧OAID解密
// taobao.top.oaid.client.decrypt
//
// 解码OAID(Open Addressee ID)，返回收件人信息。该接口用于客户端直接查看订单隐私数据，解密数据不经过ISV服务器，且包含风控等安全检测。
func TaobaoTopOaidClientDecrypt(clt *core.SDKClient, req *topoaid.TaobaoTopOaidClientDecryptAPIRequest, session string) (*topoaid.TaobaoTopOaidClientDecryptAPIResponse, error) {
	var resp topoaid.TaobaoTopOaidClientDecryptAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
