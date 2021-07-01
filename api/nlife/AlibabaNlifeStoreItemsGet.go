package nlife

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nlife"
)

/* AlibabaNlifeStoreItemsGet
获取门店的商品列表(在售|已下架|全部)
alibaba.nlife.store.items.get

利用该接口可以获取到零售+商品服务中符合条件的商品列表，包括在售的、已下架的或者是所有状态的商品。 */
func AlibabaNlifeStoreItemsGet(clt *core.SDKClient, req *nlife.AlibabaNlifeStoreItemsGetAPIRequest, session string) (*nlife.AlibabaNlifeStoreItemsGetAPIResponse, error) {
	var resp nlife.AlibabaNlifeStoreItemsGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
