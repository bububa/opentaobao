package store

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/store"
)

// Taobaoplacestoregroupupdate 门店库修改基本信息
// taobao.place.storegroup.update
//
// 门店库修改基本信息
func Taobaoplacestoregroupupdate(clt *core.SDKClient, req *store.TaobaoplacestoregroupupdateAPIRequest, session string) (*store.TaobaoplacestoregroupupdateAPIResponse, error) {
	var resp store.TaobaoplacestoregroupupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
