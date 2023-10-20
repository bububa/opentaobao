package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalsccrmcardactive 标准激活卡
// alibaba.alsc.crm.card.active
//
// 激活卡
func Alibabaalsccrmcardactive(clt *core.SDKClient, req *alsc.AlibabaalsccrmcardactiveAPIRequest, session string) (*alsc.AlibabaalsccrmcardactiveAPIResponse, error) {
	var resp alsc.AlibabaalsccrmcardactiveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
