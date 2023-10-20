package icbushowcase

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbushowcase"
)

// Alibabascbpshowcasesort 橱窗顺序变更
// alibaba.scbp.showcase.sort
//
// 橱窗顺序变更
func Alibabascbpshowcasesort(clt *core.SDKClient, req *icbushowcase.AlibabascbpshowcasesortAPIRequest, session string) (*icbushowcase.AlibabascbpshowcasesortAPIResponse, error) {
	var resp icbushowcase.AlibabascbpshowcasesortAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
