package nrpos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrpos"
)

// AlibabaMosCommdyPosmerchandiseGetmerchandise 去前置机商品在线查询
// alibaba.mos.commdy.posmerchandise.getmerchandise
//
// 去前置机商品在线查询接口
func AlibabaMosCommdyPosmerchandiseGetmerchandise(clt *core.SDKClient, req *nrpos.AlibabaMosCommdyPosmerchandiseGetmerchandiseAPIRequest, resp *nrpos.AlibabaMosCommdyPosmerchandiseGetmerchandiseAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
