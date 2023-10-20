package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// Taobaoqimenstoreupdate 门店更新接口
// taobao.qimen.store.update
//
// 商家在ERP等系统中调用该接口，更新门店信息
func Taobaoqimenstoreupdate(clt *core.SDKClient, req *qimen.TaobaoqimenstoreupdateAPIRequest, session string) (*qimen.TaobaoqimenstoreupdateAPIResponse, error) {
	var resp qimen.TaobaoqimenstoreupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
