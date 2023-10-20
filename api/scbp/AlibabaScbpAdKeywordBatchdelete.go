package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpadkeywordbatchdelete 外贸直通车批量删除关键词
// alibaba.scbp.ad.keyword.batchdelete
//
// 外贸直通车批量删除关键词
func Alibabascbpadkeywordbatchdelete(clt *core.SDKClient, req *scbp.AlibabascbpadkeywordbatchdeleteAPIRequest, session string) (*scbp.AlibabascbpadkeywordbatchdeleteAPIResponse, error) {
	var resp scbp.AlibabascbpadkeywordbatchdeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
