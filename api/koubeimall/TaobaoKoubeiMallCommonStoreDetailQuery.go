package koubeimall

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/koubeimall"
)

// Taobaokoubeimallcommonstoredetailquery 查询综合体内的门店详细信息
// taobao.koubei.mall.common.store.detail.query
//
// 查询口碑综合体内的门店详情信息
func Taobaokoubeimallcommonstoredetailquery(clt *core.SDKClient, req *koubeimall.TaobaokoubeimallcommonstoredetailqueryAPIRequest, session string) (*koubeimall.TaobaokoubeimallcommonstoredetailqueryAPIResponse, error) {
	var resp koubeimall.TaobaokoubeimallcommonstoredetailqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
