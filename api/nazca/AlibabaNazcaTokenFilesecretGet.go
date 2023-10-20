package nazca

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nazca"
)

// AlibabaNazcaTokenFilesecretGet 获取文件秘钥
// alibaba.nazca.token.filesecret.get
//
// 获取文件秘钥
func AlibabaNazcaTokenFilesecretGet(clt *core.SDKClient, req *nazca.AlibabaNazcaTokenFilesecretGetAPIRequest, resp *nazca.AlibabaNazcaTokenFilesecretGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
