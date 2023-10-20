package koubeimall

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/koubeimall"
)

// Taobaokoubeimallcommonmallauthpage 分页查询已授权的商圈列表信息
// taobao.koubei.mall.common.mall.auth.page
//
// 分页查询口碑已授权商圈的列表信息
func Taobaokoubeimallcommonmallauthpage(clt *core.SDKClient, req *koubeimall.TaobaokoubeimallcommonmallauthpageAPIRequest, session string) (*koubeimall.TaobaokoubeimallcommonmallauthpageAPIResponse, error) {
	var resp koubeimall.TaobaokoubeimallcommonmallauthpageAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
