package store

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/store"
)

// Taobaoplacestoregroupcreate 商户门店库创建接口
// taobao.place.storegroup.create
//
// 用于商家创建线下门店库
func Taobaoplacestoregroupcreate(clt *core.SDKClient, req *store.TaobaoplacestoregroupcreateAPIRequest, session string) (*store.TaobaoplacestoregroupcreateAPIResponse, error) {
	var resp store.TaobaoplacestoregroupcreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
