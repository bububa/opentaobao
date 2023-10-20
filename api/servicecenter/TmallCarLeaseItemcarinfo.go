package servicecenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// TmallCarLeaseItemcarinfo 整车租赁商品四级车型信息
// tmall.car.lease.itemcarinfo
//
// 整车租赁项目发布宝贝需要4级车型库，4级车型库信息需要回传
func TmallCarLeaseItemcarinfo(clt *core.SDKClient, req *servicecenter.TmallCarLeaseItemcarinfoAPIRequest, session string) (*servicecenter.TmallCarLeaseItemcarinfoAPIResponse, error) {
	var resp servicecenter.TmallCarLeaseItemcarinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
