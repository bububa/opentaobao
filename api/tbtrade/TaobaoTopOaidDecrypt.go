package tbtrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbtrade"
)

// TaobaoTopOaidDecrypt OAID解密
// taobao.top.oaid.decrypt
//
// 解码OAID(Open Addressee ID)，返回收件人信息。
func TaobaoTopOaidDecrypt(clt *core.SDKClient, req *tbtrade.TaobaoTopOaidDecryptAPIRequest, resp *tbtrade.TaobaoTopOaidDecryptAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
