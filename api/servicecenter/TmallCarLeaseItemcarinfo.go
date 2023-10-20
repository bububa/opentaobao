package servicecenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// Tmallcarleaseitemcarinfo 整车租赁商品四级车型信息
// tmall.car.lease.itemcarinfo
//
// 整车租赁项目发布宝贝需要4级车型库，4级车型库信息需要回传
func Tmallcarleaseitemcarinfo(clt *core.SDKClient, req *servicecenter.TmallcarleaseitemcarinfoAPIRequest, session string) (*servicecenter.TmallcarleaseitemcarinfoAPIResponse, error) {
	var resp servicecenter.TmallcarleaseitemcarinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
