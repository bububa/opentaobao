package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihouseexistinghomestorelevelquery 门店等级评分查询
// alibaba.alihouse.existinghome.store.level.query
//
// 门店等级评分查询
func Alibabaalihouseexistinghomestorelevelquery(clt *core.SDKClient, req *alihouse.AlibabaalihouseexistinghomestorelevelqueryAPIRequest, session string) (*alihouse.AlibabaalihouseexistinghomestorelevelqueryAPIResponse, error) {
	var resp alihouse.AlibabaalihouseexistinghomestorelevelqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
