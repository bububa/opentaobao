package drug

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drug"
)

// Taobaoalihealthdrugstoreget 根据店铺id获取店铺详情
// taobao.alihealth.drug.store.get
//
// 根据店铺id获取店铺详情
func Taobaoalihealthdrugstoreget(clt *core.SDKClient, req *drug.TaobaoalihealthdrugstoregetAPIRequest, session string) (*drug.TaobaoalihealthdrugstoregetAPIResponse, error) {
	var resp drug.TaobaoalihealthdrugstoregetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
