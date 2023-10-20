package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// Taobaoqimencombineitemsynchronize 组合商品接口
// taobao.qimen.combineitem.synchronize
//
// ERP调用奇门的接口,将商品信息同步给WMS
func Taobaoqimencombineitemsynchronize(clt *core.SDKClient, req *qimen.TaobaoqimencombineitemsynchronizeAPIRequest, session string) (*qimen.TaobaoqimencombineitemsynchronizeAPIResponse, error) {
	var resp qimen.TaobaoqimencombineitemsynchronizeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
