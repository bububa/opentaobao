package wangwang

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wangwang"
)

// Taobaoqianniukefuevalget 客服评价详情接口
// taobao.qianniu.kefueval.get
//
// 获取买家对客服的服务评价
func Taobaoqianniukefuevalget(clt *core.SDKClient, req *wangwang.TaobaoqianniukefuevalgetAPIRequest, session string) (*wangwang.TaobaoqianniukefuevalgetAPIResponse, error) {
	var resp wangwang.TaobaoqianniukefuevalgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
