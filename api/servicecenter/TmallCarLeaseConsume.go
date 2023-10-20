package servicecenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// Tmallcarleaseconsume 汽车租赁核销
// tmall.car.lease.consume
//
// 租赁公司回传信息，核销
func Tmallcarleaseconsume(clt *core.SDKClient, req *servicecenter.TmallcarleaseconsumeAPIRequest, session string) (*servicecenter.TmallcarleaseconsumeAPIResponse, error) {
	var resp servicecenter.TmallcarleaseconsumeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
