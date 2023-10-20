package icbushowcase

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbushowcase"
)

// AlibabaScbpShowcaseList 橱窗查询
// alibaba.scbp.showcase.list
//
// 橱窗查询
func AlibabaScbpShowcaseList(clt *core.SDKClient, req *icbushowcase.AlibabaScbpShowcaseListAPIRequest, resp *icbushowcase.AlibabaScbpShowcaseListAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
