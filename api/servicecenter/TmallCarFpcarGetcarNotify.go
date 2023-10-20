package servicecenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// Tmallcarfpcargetcarnotify 门店通知用户提车
// tmall.car.fpcar.getcar.notify
//
// 提供给外部(大搜或其它合作方)的接口-门店通知用户提车
func Tmallcarfpcargetcarnotify(clt *core.SDKClient, req *servicecenter.TmallcarfpcargetcarnotifyAPIRequest, session string) (*servicecenter.TmallcarfpcargetcarnotifyAPIResponse, error) {
	var resp servicecenter.TmallcarfpcargetcarnotifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
