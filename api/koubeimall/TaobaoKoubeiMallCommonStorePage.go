package koubeimall

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/koubeimall"
)

// Taobaokoubeimallcommonstorepage 分页查询综合体内的门店列表信息
// taobao.koubei.mall.common.store.page
//
// 分页查询综合体内的门店列表信息
func Taobaokoubeimallcommonstorepage(clt *core.SDKClient, req *koubeimall.TaobaokoubeimallcommonstorepageAPIRequest, session string) (*koubeimall.TaobaokoubeimallcommonstorepageAPIResponse, error) {
	var resp koubeimall.TaobaokoubeimallcommonstorepageAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
