package tmallcar

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcar"
)

// Tmallcarfinancestatussync 汽车金融状态同步
// tmall.car.finance.status.sync
//
// 汽车金融状态同步
func Tmallcarfinancestatussync(clt *core.SDKClient, req *tmallcar.TmallcarfinancestatussyncAPIRequest, session string) (*tmallcar.TmallcarfinancestatussyncAPIResponse, error) {
	var resp tmallcar.TmallcarfinancestatussyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
