package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpadkeywordpriceupdate 关键词改价
// alibaba.scbp.ad.keyword.price.update
//
// 关键词改价
func Alibabascbpadkeywordpriceupdate(clt *core.SDKClient, req *scbp.AlibabascbpadkeywordpriceupdateAPIRequest, session string) (*scbp.AlibabascbpadkeywordpriceupdateAPIResponse, error) {
	var resp scbp.AlibabascbpadkeywordpriceupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
