package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalsccrmcardqry 查询卡实例
// alibaba.alsc.crm.card.qry
//
// 查询卡实例（优先使用卡实例ID查询，没有则用物理卡号查询）
func Alibabaalsccrmcardqry(clt *core.SDKClient, req *alsc.AlibabaalsccrmcardqryAPIRequest, session string) (*alsc.AlibabaalsccrmcardqryAPIResponse, error) {
	var resp alsc.AlibabaalsccrmcardqryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
