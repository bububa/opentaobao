package icbushowcase

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbushowcase"
)

// Alibabascbpshowcaselist 橱窗查询
// alibaba.scbp.showcase.list
//
// 橱窗查询
func Alibabascbpshowcaselist(clt *core.SDKClient, req *icbushowcase.AlibabascbpshowcaselistAPIRequest, session string) (*icbushowcase.AlibabascbpshowcaselistAPIResponse, error) {
	var resp icbushowcase.AlibabascbpshowcaselistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
