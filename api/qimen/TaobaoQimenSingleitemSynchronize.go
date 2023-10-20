package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// Taobaoqimensingleitemsynchronize 商品同步接口
// taobao.qimen.singleitem.synchronize
//
// taobao.qimen.singleitem.synchronize
func Taobaoqimensingleitemsynchronize(clt *core.SDKClient, req *qimen.TaobaoqimensingleitemsynchronizeAPIRequest, session string) (*qimen.TaobaoqimensingleitemsynchronizeAPIResponse, error) {
	var resp qimen.TaobaoqimensingleitemsynchronizeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
