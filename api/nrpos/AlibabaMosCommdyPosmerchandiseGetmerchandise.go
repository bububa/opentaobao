package nrpos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrpos"
)

// AlibabaMosCommdyPosmerchandiseGetmerchandise 去前置机商品在线查询
// alibaba.mos.commdy.posmerchandise.getmerchandise
//
// 去前置机商品在线查询接口
func AlibabaMosCommdyPosmerchandiseGetmerchandise(clt *core.SDKClient, req *nrpos.AlibabaMosCommdyPosmerchandiseGetmerchandiseAPIRequest, session string) (*nrpos.AlibabaMosCommdyPosmerchandiseGetmerchandiseAPIResponse, error) {
	var resp nrpos.AlibabaMosCommdyPosmerchandiseGetmerchandiseAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
