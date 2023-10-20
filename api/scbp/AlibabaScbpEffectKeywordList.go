package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpeffectkeywordlist 关键词报表
// alibaba.scbp.effect.keyword.list
//
// 关键词报表
func Alibabascbpeffectkeywordlist(clt *core.SDKClient, req *scbp.AlibabascbpeffectkeywordlistAPIRequest, session string) (*scbp.AlibabascbpeffectkeywordlistAPIResponse, error) {
	var resp scbp.AlibabascbpeffectkeywordlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
