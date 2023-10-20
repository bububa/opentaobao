package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalsccrmcardopen 标准开卡流程
// alibaba.alsc.crm.card.open
//
// 标准开卡流程
func Alibabaalsccrmcardopen(clt *core.SDKClient, req *alsc.AlibabaalsccrmcardopenAPIRequest, session string) (*alsc.AlibabaalsccrmcardopenAPIResponse, error) {
	var resp alsc.AlibabaalsccrmcardopenAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
