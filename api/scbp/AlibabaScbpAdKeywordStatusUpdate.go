package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpadkeywordstatusupdate 关键词启动暂停推广
// alibaba.scbp.ad.keyword.status.update
//
// 关键词启动暂停推广
func Alibabascbpadkeywordstatusupdate(clt *core.SDKClient, req *scbp.AlibabascbpadkeywordstatusupdateAPIRequest, session string) (*scbp.AlibabascbpadkeywordstatusupdateAPIResponse, error) {
	var resp scbp.AlibabascbpadkeywordstatusupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
