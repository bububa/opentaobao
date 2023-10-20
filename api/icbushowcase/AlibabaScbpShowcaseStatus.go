package icbushowcase

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbushowcase"
)

// Alibabascbpshowcasestatus 橱窗状态
// alibaba.scbp.showcase.status
//
// 查询橱窗状态，如总数、可用数量
func Alibabascbpshowcasestatus(clt *core.SDKClient, req *icbushowcase.AlibabascbpshowcasestatusAPIRequest, session string) (*icbushowcase.AlibabascbpshowcasestatusAPIResponse, error) {
	var resp icbushowcase.AlibabascbpshowcasestatusAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
