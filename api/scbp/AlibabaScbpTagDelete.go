package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbptagdelete 删除关键词分组
// alibaba.scbp.tag.delete
//
// 删除关键词分组
func Alibabascbptagdelete(clt *core.SDKClient, req *scbp.AlibabascbptagdeleteAPIRequest, session string) (*scbp.AlibabascbptagdeleteAPIResponse, error) {
	var resp scbp.AlibabascbptagdeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
