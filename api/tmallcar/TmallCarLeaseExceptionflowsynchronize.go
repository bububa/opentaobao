package tmallcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// Tmallcarleaseexceptionflowsynchronize 天猫开新车租后异常流线下处理状态通知接口
// tmall.car.lease.exceptionflowsynchronize
//
// 天猫开新车租后异常流线下处理状态通知接口
func Tmallcarleaseexceptionflowsynchronize(clt *core.SDKClient, req *tmallcar.TmallcarleaseexceptionflowsynchronizeAPIRequest, session string) (*tmallcar.TmallcarleaseexceptionflowsynchronizeAPIResponse, error) {
	var resp tmallcar.TmallcarleaseexceptionflowsynchronizeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
