package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpadkeywordupdatekeywordstatusbatch 修改关键词状态
// alibaba.scbp.ad.keyword.update.keyword.status.batch
//
// 修改关键词状态
func Alibabascbpadkeywordupdatekeywordstatusbatch(clt *core.SDKClient, req *scbp.AlibabascbpadkeywordupdatekeywordstatusbatchAPIRequest, session string) (*scbp.AlibabascbpadkeywordupdatekeywordstatusbatchAPIResponse, error) {
	var resp scbp.AlibabascbpadkeywordupdatekeywordstatusbatchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
