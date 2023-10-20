package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpadkeyworddelete 外贸直通车删除关键词
// alibaba.scbp.ad.keyword.delete
//
// 外贸直通车删除关键词
func Alibabascbpadkeyworddelete(clt *core.SDKClient, req *scbp.AlibabascbpadkeyworddeleteAPIRequest, session string) (*scbp.AlibabascbpadkeyworddeleteAPIResponse, error) {
	var resp scbp.AlibabascbpadkeyworddeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
