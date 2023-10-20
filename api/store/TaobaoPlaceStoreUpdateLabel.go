package store

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/store"
)

// Taobaoplacestoreupdatelabel 商户门店标签更新接口
// taobao.place.store.update.label
//
// 更新商户门店标签（服务、权益、标签）接口
func Taobaoplacestoreupdatelabel(clt *core.SDKClient, req *store.TaobaoplacestoreupdatelabelAPIRequest, session string) (*store.TaobaoplacestoreupdatelabelAPIResponse, error) {
	var resp store.TaobaoplacestoreupdatelabelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
