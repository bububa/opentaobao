package nlife

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nlife"
)

// AlibabaNlifeStoreItemdetailsGet 批量获取零售加商品详情信息
// alibaba.nlife.store.itemdetails.get
//
// 批量获取零售加平台上的商品详情信息
func AlibabaNlifeStoreItemdetailsGet(clt *core.SDKClient, req *nlife.AlibabaNlifeStoreItemdetailsGetAPIRequest, session string) (*nlife.AlibabaNlifeStoreItemdetailsGetAPIResponse, error) {
	var resp nlife.AlibabaNlifeStoreItemdetailsGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
