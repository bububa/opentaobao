package nazca

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nazca"
)

// Alibabanazcatokenfilesecretget 获取文件秘钥
// alibaba.nazca.token.filesecret.get
//
// 获取文件秘钥
func Alibabanazcatokenfilesecretget(clt *core.SDKClient, req *nazca.AlibabanazcatokenfilesecretgetAPIRequest, session string) (*nazca.AlibabanazcatokenfilesecretgetAPIResponse, error) {
	var resp nazca.AlibabanazcatokenfilesecretgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
