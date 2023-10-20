package servicecenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// Tmallcarleaseriskcallback 整车租赁风控模型回调
// tmall.car.lease.riskcallback
//
// 租赁公司回调风控结果
func Tmallcarleaseriskcallback(clt *core.SDKClient, req *servicecenter.TmallcarleaseriskcallbackAPIRequest, session string) (*servicecenter.TmallcarleaseriskcallbackAPIResponse, error) {
	var resp servicecenter.TmallcarleaseriskcallbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
