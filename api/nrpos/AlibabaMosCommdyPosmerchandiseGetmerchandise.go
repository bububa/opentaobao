package nrpos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrpos"
)

// Alibabamoscommdyposmerchandisegetmerchandise 去前置机商品在线查询
// alibaba.mos.commdy.posmerchandise.getmerchandise
//
// 去前置机商品在线查询接口
func Alibabamoscommdyposmerchandisegetmerchandise(clt *core.SDKClient, req *nrpos.AlibabamoscommdyposmerchandisegetmerchandiseAPIRequest, session string) (*nrpos.AlibabamoscommdyposmerchandisegetmerchandiseAPIResponse, error) {
	var resp nrpos.AlibabamoscommdyposmerchandisegetmerchandiseAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
