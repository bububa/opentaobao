package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpadkeywordget 外贸直通车查询关键词
// alibaba.scbp.ad.keyword.get
//
// 外贸直通车查询关键词
func Alibabascbpadkeywordget(clt *core.SDKClient, req *scbp.AlibabascbpadkeywordgetAPIRequest, session string) (*scbp.AlibabascbpadkeywordgetAPIResponse, error) {
	var resp scbp.AlibabascbpadkeywordgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
