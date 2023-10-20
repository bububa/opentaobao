package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpadkeywordpricebatchupdate 关键词批量改价
// alibaba.scbp.ad.keyword.price.batchupdate
//
// 关键词批量改价
func Alibabascbpadkeywordpricebatchupdate(clt *core.SDKClient, req *scbp.AlibabascbpadkeywordpricebatchupdateAPIRequest, session string) (*scbp.AlibabascbpadkeywordpricebatchupdateAPIResponse, error) {
	var resp scbp.AlibabascbpadkeywordpricebatchupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
