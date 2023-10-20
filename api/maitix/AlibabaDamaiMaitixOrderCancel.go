package maitix

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/maitix"
)

// Alibabadamaimaitixordercancel 大麦-库存释放
// alibaba.damai.maitix.order.cancel
//
// 库存释放
func Alibabadamaimaitixordercancel(clt *core.SDKClient, req *maitix.AlibabadamaimaitixordercancelAPIRequest, session string) (*maitix.AlibabadamaimaitixordercancelAPIResponse, error) {
	var resp maitix.AlibabadamaimaitixordercancelAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
