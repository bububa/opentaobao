package caipiao

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/caipiao"
)

// Taobaocaipiaogoodsinfoget 根据卖家id与appkey获取商品信息
// taobao.caipiao.goods.info.get
//
// 根据卖家id与appkey获取商品信息。
func Taobaocaipiaogoodsinfoget(clt *core.SDKClient, req *caipiao.TaobaocaipiaogoodsinfogetAPIRequest, session string) (*caipiao.TaobaocaipiaogoodsinfogetAPIResponse, error) {
	var resp caipiao.TaobaocaipiaogoodsinfogetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
