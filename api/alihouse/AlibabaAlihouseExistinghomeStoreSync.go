package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihouseexistinghomestoresync 二手房标准门店数据同步
// alibaba.alihouse.existinghome.store.sync
//
// 二手房标准门店数据同步
func Alibabaalihouseexistinghomestoresync(clt *core.SDKClient, req *alihouse.AlibabaalihouseexistinghomestoresyncAPIRequest, session string) (*alihouse.AlibabaalihouseexistinghomestoresyncAPIResponse, error) {
	var resp alihouse.AlibabaalihouseexistinghomestoresyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
