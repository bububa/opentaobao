package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

/* TaobaoQimenStoreUpdate
门店更新接口
taobao.qimen.store.update

商家在ERP等系统中调用该接口，更新门店信息 */
func TaobaoQimenStoreUpdate(clt *core.SDKClient, req *qimen.TaobaoQimenStoreUpdateAPIRequest, session string) (*qimen.TaobaoQimenStoreUpdateAPIResponse, error) {
	var resp qimen.TaobaoQimenStoreUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
