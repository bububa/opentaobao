package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// Alibabaalsccrmcardquerytemplate 查询卡模板详情
// alibaba.alsc.crm.card.query.template
//
// 查询卡模板详情
func Alibabaalsccrmcardquerytemplate(clt *core.SDKClient, req *alsc.AlibabaalsccrmcardquerytemplateAPIRequest, session string) (*alsc.AlibabaalsccrmcardquerytemplateAPIResponse, error) {
	var resp alsc.AlibabaalsccrmcardquerytemplateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
