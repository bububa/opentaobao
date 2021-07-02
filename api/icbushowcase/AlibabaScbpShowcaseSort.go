package icbushowcase

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbushowcase"
)

// AlibabaScbpShowcaseSort 橱窗顺序变更
// alibaba.scbp.showcase.sort
//
// 橱窗顺序变更
func AlibabaScbpShowcaseSort(clt *core.SDKClient, req *icbushowcase.AlibabaScbpShowcaseSortAPIRequest, session string) (*icbushowcase.AlibabaScbpShowcaseSortAPIResponse, error) {
	var resp icbushowcase.AlibabaScbpShowcaseSortAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
