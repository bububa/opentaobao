package store

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/store"
)

// Taobaoplacestoregroupdelete 删除门店库
// taobao.place.storegroup.delete
//
// 删除门店库
func Taobaoplacestoregroupdelete(clt *core.SDKClient, req *store.TaobaoplacestoregroupdeleteAPIRequest, session string) (*store.TaobaoplacestoregroupdeleteAPIResponse, error) {
	var resp store.TaobaoplacestoregroupdeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
