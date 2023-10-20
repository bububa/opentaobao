package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbptaglist 查询所有分组
// alibaba.scbp.tag.list
//
// 查询所有分组
func Alibabascbptaglist(clt *core.SDKClient, req *scbp.AlibabascbptaglistAPIRequest, session string) (*scbp.AlibabascbptaglistAPIResponse, error) {
	var resp scbp.AlibabascbptaglistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
