package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpadkeywordtagupdate 修改关键词所属分组
// alibaba.scbp.ad.keyword.tag.update
//
// 修改关键词所属分组
func Alibabascbpadkeywordtagupdate(clt *core.SDKClient, req *scbp.AlibabascbpadkeywordtagupdateAPIRequest, session string) (*scbp.AlibabascbpadkeywordtagupdateAPIResponse, error) {
	var resp scbp.AlibabascbpadkeywordtagupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
