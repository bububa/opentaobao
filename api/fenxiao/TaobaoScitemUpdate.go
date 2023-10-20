package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// Taobaoscitemupdate 根据商品ID或商家编码修改后端商品
// taobao.scitem.update
//
// 根据商品ID或商家编码修改后端商品
func Taobaoscitemupdate(clt *core.SDKClient, req *fenxiao.TaobaoscitemupdateAPIRequest, session string) (*fenxiao.TaobaoscitemupdateAPIResponse, error) {
	var resp fenxiao.TaobaoscitemupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
