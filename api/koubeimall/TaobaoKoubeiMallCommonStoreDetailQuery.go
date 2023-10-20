package koubeimall

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/koubeimall"
)

// TaobaoKoubeiMallCommonStoreDetailQuery 查询综合体内的门店详细信息
// taobao.koubei.mall.common.store.detail.query
//
// 查询口碑综合体内的门店详情信息
func TaobaoKoubeiMallCommonStoreDetailQuery(clt *core.SDKClient, req *koubeimall.TaobaoKoubeiMallCommonStoreDetailQueryAPIRequest, resp *koubeimall.TaobaoKoubeiMallCommonStoreDetailQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
