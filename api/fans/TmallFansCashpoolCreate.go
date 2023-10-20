package fans

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fans"
)

// Tmallfanscashpoolcreate 创建资金池
// tmall.fans.cashpool.create
//
// 商家创建资金池接口
func Tmallfanscashpoolcreate(clt *core.SDKClient, req *fans.TmallfanscashpoolcreateAPIRequest, session string) (*fans.TmallfanscashpoolcreateAPIResponse, error) {
	var resp fans.TmallfanscashpoolcreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
