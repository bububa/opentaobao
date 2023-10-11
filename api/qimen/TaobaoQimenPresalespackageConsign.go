package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenPresalespackageConsign 预售预包尾款推单发货
// taobao.qimen.presalespackage.consign
//
// 预售预包尾款推单发货
func TaobaoQimenPresalespackageConsign(clt *core.SDKClient, req *qimen.TaobaoQimenPresalespackageConsignAPIRequest, session string) (*qimen.TaobaoQimenPresalespackageConsignAPIResponse, error) {
	var resp qimen.TaobaoQimenPresalespackageConsignAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
