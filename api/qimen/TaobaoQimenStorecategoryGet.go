package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// Taobaoqimenstorecategoryget 门店类目获取接口
// taobao.qimen.storecategory.get
//
// 商家在ERP中调用该接口，获取门店类目
func Taobaoqimenstorecategoryget(clt *core.SDKClient, req *qimen.TaobaoqimenstorecategorygetAPIRequest, session string) (*qimen.TaobaoqimenstorecategorygetAPIResponse, error) {
	var resp qimen.TaobaoqimenstorecategorygetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
