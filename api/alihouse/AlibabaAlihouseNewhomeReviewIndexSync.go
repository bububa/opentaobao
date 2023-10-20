package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomereviewindexsync 新测评乐居指数接口
// alibaba.alihouse.newhome.review.index.sync
//
// 新测评乐居指数同步数据
func Alibabaalihousenewhomereviewindexsync(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomereviewindexsyncAPIRequest, session string) (*alihouse.AlibabaalihousenewhomereviewindexsyncAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomereviewindexsyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
