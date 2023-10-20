package drug

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drug"
)

// Taobaoalihealthdrugstoresearch 药品店内搜索
// taobao.alihealth.drug.store.search
//
// 提供给千牛智能客服，在阿里健康O2O店铺内搜索药品
func Taobaoalihealthdrugstoresearch(clt *core.SDKClient, req *drug.TaobaoalihealthdrugstoresearchAPIRequest, session string) (*drug.TaobaoalihealthdrugstoresearchAPIResponse, error) {
	var resp drug.TaobaoalihealthdrugstoresearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
