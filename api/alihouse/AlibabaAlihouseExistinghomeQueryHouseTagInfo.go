package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihouseexistinghomequeryhousetaginfo 活动标查询接口
// alibaba.alihouse.existinghome.query.house.tag.info
//
// 活动标查询接口
func Alibabaalihouseexistinghomequeryhousetaginfo(clt *core.SDKClient, req *alihouse.AlibabaalihouseexistinghomequeryhousetaginfoAPIRequest, session string) (*alihouse.AlibabaalihouseexistinghomequeryhousetaginfoAPIResponse, error) {
	var resp alihouse.AlibabaalihouseexistinghomequeryhousetaginfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
