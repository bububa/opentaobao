package icbushowcase

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbushowcase"
)

/* AlibabaScbpShowcaseStatus
橱窗状态
alibaba.scbp.showcase.status

查询橱窗状态，如总数、可用数量 */
func AlibabaScbpShowcaseStatus(clt *core.SDKClient, req *icbushowcase.AlibabaScbpShowcaseStatusAPIRequest, session string) (*icbushowcase.AlibabaScbpShowcaseStatusAPIResponse, error) {
	var resp icbushowcase.AlibabaScbpShowcaseStatusAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
