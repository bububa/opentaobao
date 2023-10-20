package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Taobaoscitemadd 发布后端商品
// taobao.scitem.add
//
// 发布后端商品
func Taobaoscitemadd(clt *core.SDKClient, req *fenxiao.TaobaoscitemaddAPIRequest, session string) (*fenxiao.TaobaoscitemaddAPIResponse, error) {
	var resp fenxiao.TaobaoscitemaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
