package fenxiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoScitemGet 根据id查询商品
// taobao.scitem.get
//
// 根据id查询商品
func TaobaoScitemGet(clt *core.SDKClient, req *fenxiao.TaobaoScitemGetAPIRequest, session string) (*fenxiao.TaobaoScitemGetAPIResponse, error) {
	var resp fenxiao.TaobaoScitemGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
