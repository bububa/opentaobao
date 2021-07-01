package store

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/store"
)

/* TaobaoPlaceStoreUpdateLabel
商户门店标签更新接口
taobao.place.store.update.label

更新商户门店标签（服务、权益、标签）接口 */
func TaobaoPlaceStoreUpdateLabel(clt *core.SDKClient, req *store.TaobaoPlaceStoreUpdateLabelAPIRequest, session string) (*store.TaobaoPlaceStoreUpdateLabelAPIResponse, error) {
	var resp store.TaobaoPlaceStoreUpdateLabelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
