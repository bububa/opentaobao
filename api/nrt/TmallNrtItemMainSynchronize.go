package nrt

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

// Tmallnrtitemmainsynchronize 家装新零售主商品同步至阿里
// tmall.nrt.item.main.synchronize
//
// 同步卖场存量线下商品到阿里
func Tmallnrtitemmainsynchronize(clt *core.SDKClient, req *nrt.TmallnrtitemmainsynchronizeAPIRequest, session string) (*nrt.TmallnrtitemmainsynchronizeAPIResponse, error) {
	var resp nrt.TmallnrtitemmainsynchronizeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
