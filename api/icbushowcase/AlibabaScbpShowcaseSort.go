package icbushowcase

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbushowcase"
)

// AlibabaScbpShowcaseSort 橱窗顺序变更
// alibaba.scbp.showcase.sort
//
// 橱窗顺序变更
func AlibabaScbpShowcaseSort(clt *core.SDKClient, req *icbushowcase.AlibabaScbpShowcaseSortAPIRequest, resp *icbushowcase.AlibabaScbpShowcaseSortAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
