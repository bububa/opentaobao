package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpadkeywordrankpricebatchget 外贸直通车关键词前五名批量排价
// alibaba.scbp.ad.keyword.rank.price.batchget
//
// 外贸直通车关键词前五名批量排价
func Alibabascbpadkeywordrankpricebatchget(clt *core.SDKClient, req *scbp.AlibabascbpadkeywordrankpricebatchgetAPIRequest, session string) (*scbp.AlibabascbpadkeywordrankpricebatchgetAPIResponse, error) {
	var resp scbp.AlibabascbpadkeywordrankpricebatchgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
