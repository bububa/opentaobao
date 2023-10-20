package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpadkeywordadd 外贸直通车加词
// alibaba.scbp.ad.keyword.add
//
// 外贸直通车加词服务
func Alibabascbpadkeywordadd(clt *core.SDKClient, req *scbp.AlibabascbpadkeywordaddAPIRequest, session string) (*scbp.AlibabascbpadkeywordaddAPIResponse, error) {
	var resp scbp.AlibabascbpadkeywordaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
