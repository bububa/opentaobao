package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpadkeywordcreatekeywordbatch 关键词添加
// alibaba.scbp.ad.keyword.create.keyword.batch
//
// 关键词添加
func Alibabascbpadkeywordcreatekeywordbatch(clt *core.SDKClient, req *scbp.AlibabascbpadkeywordcreatekeywordbatchAPIRequest, session string) (*scbp.AlibabascbpadkeywordcreatekeywordbatchAPIResponse, error) {
	var resp scbp.AlibabascbpadkeywordcreatekeywordbatchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
