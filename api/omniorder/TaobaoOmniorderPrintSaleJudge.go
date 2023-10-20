package omniorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// Taobaoomniorderprintsalejudge 导购员判断
// taobao.omniorder.print.sale.judge
//
// 用于判断当前子账号是否导购员
func Taobaoomniorderprintsalejudge(clt *core.SDKClient, req *omniorder.TaobaoomniorderprintsalejudgeAPIRequest, session string) (*omniorder.TaobaoomniorderprintsalejudgeAPIResponse, error) {
	var resp omniorder.TaobaoomniorderprintsalejudgeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
