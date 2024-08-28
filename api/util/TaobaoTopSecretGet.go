package util

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/util"
)

// TaobaoTopSecretGet 获取TOP通道解密秘钥
// taobao.top.secret.get
//
// top sdk通过api获取对应解密秘钥
func TaobaoTopSecretGet(ctx context.Context, clt *core.SDKClient, req *util.TaobaoTopSecretGetAPIRequest, resp *util.TaobaoTopSecretGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
