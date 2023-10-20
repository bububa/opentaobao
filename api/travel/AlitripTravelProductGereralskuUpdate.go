package travel

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/travel"
)

// AlitripTravelProductGereralskuUpdate (供销)船票通用类目sku新增&编辑API
// alitrip.travel.product.gereralsku.update
//
// 发布SKU信息（如果properties重复 则更新）
func AlitripTravelProductGereralskuUpdate(clt *core.SDKClient, req *travel.AlitripTravelProductGereralskuUpdateAPIRequest, session string) (*travel.AlitripTravelProductGereralskuUpdateAPIResponse, error) {
	var resp travel.AlitripTravelProductGereralskuUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
