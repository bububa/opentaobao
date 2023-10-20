package store

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/store"
)

// Taobaoplacestoreextendupdate 商户门店拓展信息更新接口
// taobao.place.store.extend.update
//
// 更新商户门店拓展信息（tags、attribute、bizAtrribute）更新接口
func Taobaoplacestoreextendupdate(clt *core.SDKClient, req *store.TaobaoplacestoreextendupdateAPIRequest, session string) (*store.TaobaoplacestoreextendupdateAPIResponse, error) {
	var resp store.TaobaoplacestoreextendupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
