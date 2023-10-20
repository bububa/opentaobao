package scbp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// Alibabascbpadaccountlevelget 查询推广账户等级
// alibaba.scbp.ad.account.level.get
//
// 查询推广账户等级
func Alibabascbpadaccountlevelget(clt *core.SDKClient, req *scbp.AlibabascbpadaccountlevelgetAPIRequest, session string) (*scbp.AlibabascbpadaccountlevelgetAPIResponse, error) {
	var resp scbp.AlibabascbpadaccountlevelgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
