package qimen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// Taobaoqimenitemstorebanding 商品关联绑定接口
// taobao.qimen.itemstore.banding
//
// 商家在ERP等系统中调用该接口，将线上商品和线下门店“新建/删除”关联。这里的线上。每次只能单个商品关联多个门店，门店上限200
func Taobaoqimenitemstorebanding(clt *core.SDKClient, req *qimen.TaobaoqimenitemstorebandingAPIRequest, session string) (*qimen.TaobaoqimenitemstorebandingAPIResponse, error) {
	var resp qimen.TaobaoqimenitemstorebandingAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
