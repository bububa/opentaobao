package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Taobaoscitemget 根据id查询商品
// taobao.scitem.get
//
// 根据id查询商品
func Taobaoscitemget(clt *core.SDKClient, req *fenxiao.TaobaoscitemgetAPIRequest, session string) (*fenxiao.TaobaoscitemgetAPIResponse, error) {
	var resp fenxiao.TaobaoscitemgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
