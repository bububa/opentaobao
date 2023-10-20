package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpadkeywordrankpriceget 外贸直通车关键词前五名排价
// alibaba.scbp.ad.keyword.rank.price.get
//
// 外贸直通车关键词前五名排价
func Alibabascbpadkeywordrankpriceget(clt *core.SDKClient, req *scbp.AlibabascbpadkeywordrankpricegetAPIRequest, session string) (*scbp.AlibabascbpadkeywordrankpricegetAPIResponse, error) {
	var resp scbp.AlibabascbpadkeywordrankpricegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
