package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbptagadd 创建关键词分组
// alibaba.scbp.tag.add
//
// 创建关键词分组
func Alibabascbptagadd(clt *core.SDKClient, req *scbp.AlibabascbptagaddAPIRequest, session string) (*scbp.AlibabascbptagaddAPIResponse, error) {
	var resp scbp.AlibabascbptagaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
