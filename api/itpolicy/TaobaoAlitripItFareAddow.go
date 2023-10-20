package itpolicy

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/itpolicy"
)

// Taobaoalitripitfareaddow 【国际机票自有政策】单条单程添加
// taobao.alitrip.it.fare.addow
//
// 自有政策单程添加接口，重复的老数据会被删除，重复判断规则同excel
func Taobaoalitripitfareaddow(clt *core.SDKClient, req *itpolicy.TaobaoalitripitfareaddowAPIRequest, session string) (*itpolicy.TaobaoalitripitfareaddowAPIResponse, error) {
	var resp itpolicy.TaobaoalitripitfareaddowAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
