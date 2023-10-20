package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbptagrename 重命名关键词分组
// alibaba.scbp.tag.rename
//
// 重命名关键词分组
func Alibabascbptagrename(clt *core.SDKClient, req *scbp.AlibabascbptagrenameAPIRequest, session string) (*scbp.AlibabascbptagrenameAPIResponse, error) {
	var resp scbp.AlibabascbptagrenameAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
