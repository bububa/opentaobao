package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpadkeyworddeletekeywordbatch 删除关键词
// alibaba.scbp.ad.keyword.delete.keyword.batch
//
// 删除关键词
func Alibabascbpadkeyworddeletekeywordbatch(clt *core.SDKClient, req *scbp.AlibabascbpadkeyworddeletekeywordbatchAPIRequest, session string) (*scbp.AlibabascbpadkeyworddeletekeywordbatchAPIResponse, error) {
	var resp scbp.AlibabascbpadkeyworddeletekeywordbatchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
