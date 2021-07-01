package drug

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drug"
)

/* TaobaoAlihealthDrugStoreGet
根据店铺id获取店铺详情
taobao.alihealth.drug.store.get

根据店铺id获取店铺详情 */
func TaobaoAlihealthDrugStoreGet(clt *core.SDKClient, req *drug.TaobaoAlihealthDrugStoreGetAPIRequest, session string) (*drug.TaobaoAlihealthDrugStoreGetAPIResponse, error) {
	var resp drug.TaobaoAlihealthDrugStoreGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
