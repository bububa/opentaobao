package itpolicy

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/itpolicy"
)

// Taobaoalitripitfareget 【国际机票自有政策】单条查询
// taobao.alitrip.it.fare.get
//
// 通过此接口可以查询单条政策的详情，可以根据fareId或outId查询，用户outId查询时，如果outId不唯一，只返回最新添加的一条数据
func Taobaoalitripitfareget(clt *core.SDKClient, req *itpolicy.TaobaoalitripitfaregetAPIRequest, session string) (*itpolicy.TaobaoalitripitfaregetAPIResponse, error) {
	var resp itpolicy.TaobaoalitripitfaregetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
