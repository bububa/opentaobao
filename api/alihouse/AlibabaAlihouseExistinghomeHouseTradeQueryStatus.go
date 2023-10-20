package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihouseexistinghomehousetradequerystatus 查询房源状态接口
// alibaba.alihouse.existinghome.house.trade.query.status
//
// 查询房源状态接口
func Alibabaalihouseexistinghomehousetradequerystatus(clt *core.SDKClient, req *alihouse.AlibabaalihouseexistinghomehousetradequerystatusAPIRequest, session string) (*alihouse.AlibabaalihouseexistinghomehousetradequerystatusAPIResponse, error) {
	var resp alihouse.AlibabaalihouseexistinghomehousetradequerystatusAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
