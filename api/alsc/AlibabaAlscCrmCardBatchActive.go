package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalsccrmcardbatchactive 批量激活卡
// alibaba.alsc.crm.card.batch.active
//
// 批量激活卡
func Alibabaalsccrmcardbatchactive(clt *core.SDKClient, req *alsc.AlibabaalsccrmcardbatchactiveAPIRequest, session string) (*alsc.AlibabaalsccrmcardbatchactiveAPIResponse, error) {
	var resp alsc.AlibabaalsccrmcardbatchactiveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
