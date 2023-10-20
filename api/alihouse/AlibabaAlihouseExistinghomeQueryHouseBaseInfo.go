package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihouseexistinghomequeryhousebaseinfo 查询房源基本信息
// alibaba.alihouse.existinghome.query.house.base.info
//
// 查询房源基本信息
func Alibabaalihouseexistinghomequeryhousebaseinfo(clt *core.SDKClient, req *alihouse.AlibabaalihouseexistinghomequeryhousebaseinfoAPIRequest, session string) (*alihouse.AlibabaalihouseexistinghomequeryhousebaseinfoAPIResponse, error) {
	var resp alihouse.AlibabaalihouseexistinghomequeryhousebaseinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
