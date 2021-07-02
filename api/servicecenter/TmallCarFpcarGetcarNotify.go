package servicecenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// TmallCarFpcarGetcarNotify 门店通知用户提车
// tmall.car.fpcar.getcar.notify
//
// 提供给外部(大搜或其它合作方)的接口-门店通知用户提车
func TmallCarFpcarGetcarNotify(clt *core.SDKClient, req *servicecenter.TmallCarFpcarGetcarNotifyAPIRequest, session string) (*servicecenter.TmallCarFpcarGetcarNotifyAPIResponse, error) {
	var resp servicecenter.TmallCarFpcarGetcarNotifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
