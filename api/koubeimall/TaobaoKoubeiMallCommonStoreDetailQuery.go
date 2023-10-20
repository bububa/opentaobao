package koubeimall

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/koubeimall"
)

// TaobaoKoubeiMallCommonStoreDetailQuery 查询综合体内的门店详细信息
// taobao.koubei.mall.common.store.detail.query
//
// 查询口碑综合体内的门店详情信息
func TaobaoKoubeiMallCommonStoreDetailQuery(clt *core.SDKClient, req *koubeimall.TaobaoKoubeiMallCommonStoreDetailQueryAPIRequest, session string) (*koubeimall.TaobaoKoubeiMallCommonStoreDetailQueryAPIResponse, error) {
	var resp koubeimall.TaobaoKoubeiMallCommonStoreDetailQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
