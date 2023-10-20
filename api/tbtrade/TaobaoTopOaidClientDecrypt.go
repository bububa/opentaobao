package tbtrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbtrade"
)

// TaobaoTopOaidClientDecrypt 端侧OAID解密
// taobao.top.oaid.client.decrypt
//
// 解码OAID(Open Addressee ID)，返回收件人信息。该接口用于客户端直接查看订单隐私数据，解密数据不经过ISV服务器，且包含风控等安全检测。
func TaobaoTopOaidClientDecrypt(clt *core.SDKClient, req *tbtrade.TaobaoTopOaidClientDecryptAPIRequest, session string) (*tbtrade.TaobaoTopOaidClientDecryptAPIResponse, error) {
	var resp tbtrade.TaobaoTopOaidClientDecryptAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
