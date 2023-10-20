package cainiaohandover

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/cainiaohandover"
)

// Cainiaoglobalhandovercontentquery 查询大包详情
// cainiao.global.handover.content.query
//
// 查询大包详情
func Cainiaoglobalhandovercontentquery(clt *core.SDKClient, req *cainiaohandover.CainiaoglobalhandovercontentqueryAPIRequest, session string) (*cainiaohandover.CainiaoglobalhandovercontentqueryAPIResponse, error) {
	var resp cainiaohandover.CainiaoglobalhandovercontentqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
