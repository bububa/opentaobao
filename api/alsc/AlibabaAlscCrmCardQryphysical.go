package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalsccrmcardqryphysical 查询物理卡
// alibaba.alsc.crm.card.qryphysical
//
// 查询物理卡
func Alibabaalsccrmcardqryphysical(clt *core.SDKClient, req *alsc.AlibabaalsccrmcardqryphysicalAPIRequest, session string) (*alsc.AlibabaalsccrmcardqryphysicalAPIResponse, error) {
	var resp alsc.AlibabaalsccrmcardqryphysicalAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
