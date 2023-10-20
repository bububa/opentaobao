package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalsccrmcardbatchsell 批量开卡（售卡）
// alibaba.alsc.crm.card.batch.sell
//
// 批量开卡（售卡）
func Alibabaalsccrmcardbatchsell(clt *core.SDKClient, req *alsc.AlibabaalsccrmcardbatchsellAPIRequest, session string) (*alsc.AlibabaalsccrmcardbatchsellAPIResponse, error) {
	var resp alsc.AlibabaalsccrmcardbatchsellAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
