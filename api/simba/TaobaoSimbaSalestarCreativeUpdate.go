package simba

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/simba"
)

// TaobaoSimbaSalestarCreativeUpdate 销量明星更新创意相关接口
// taobao.simba.salestar.creative.update
//
// 更新一个创意的信息，可以设置创意标题、创意图片
func TaobaoSimbaSalestarCreativeUpdate(clt *core.SDKClient, req *simba.TaobaoSimbaSalestarCreativeUpdateAPIRequest, session string) (*simba.TaobaoSimbaSalestarCreativeUpdateAPIResponse, error) {
	var resp simba.TaobaoSimbaSalestarCreativeUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
