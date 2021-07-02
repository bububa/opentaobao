package icbushowcase

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbushowcase"
)

// AlibabaScbpShowcaseList 橱窗查询
// alibaba.scbp.showcase.list
//
// 橱窗查询
func AlibabaScbpShowcaseList(clt *core.SDKClient, req *icbushowcase.AlibabaScbpShowcaseListAPIRequest, session string) (*icbushowcase.AlibabaScbpShowcaseListAPIResponse, error) {
	var resp icbushowcase.AlibabaScbpShowcaseListAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
