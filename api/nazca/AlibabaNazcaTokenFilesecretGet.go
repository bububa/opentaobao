package nazca

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nazca"
)

// AlibabaNazcaTokenFilesecretGet 获取文件秘钥
// alibaba.nazca.token.filesecret.get
//
// 获取文件秘钥
func AlibabaNazcaTokenFilesecretGet(ctx context.Context, clt *core.SDKClient, req *nazca.AlibabaNazcaTokenFilesecretGetAPIRequest, resp *nazca.AlibabaNazcaTokenFilesecretGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
