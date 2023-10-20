package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpadkeywordrankget 获取外贸直通车关键词预估排名
// alibaba.scbp.ad.keyword.rank.get
//
// 获取外贸直通车关键词预估排名
func Alibabascbpadkeywordrankget(clt *core.SDKClient, req *scbp.AlibabascbpadkeywordrankgetAPIRequest, session string) (*scbp.AlibabascbpadkeywordrankgetAPIResponse, error) {
	var resp scbp.AlibabascbpadkeywordrankgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
