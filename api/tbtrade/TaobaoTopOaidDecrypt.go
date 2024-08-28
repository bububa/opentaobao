package tbtrade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbtrade"
)

// TaobaoTopOaidDecrypt OAID解密
// taobao.top.oaid.decrypt
//
// 解码OAID(Open Addressee ID)，返回收件人信息。
func TaobaoTopOaidDecrypt(ctx context.Context, clt *core.SDKClient, req *tbtrade.TaobaoTopOaidDecryptAPIRequest, resp *tbtrade.TaobaoTopOaidDecryptAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
