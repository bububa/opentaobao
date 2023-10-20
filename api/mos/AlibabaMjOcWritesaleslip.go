package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// AlibabaMjOcWritesaleslip 开票占库
// alibaba.mj.oc.writesaleslip
//
// 开票占库
func AlibabaMjOcWritesaleslip(clt *core.SDKClient, req *mos.AlibabaMjOcWritesaleslipAPIRequest, resp *mos.AlibabaMjOcWritesaleslipAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
