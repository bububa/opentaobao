package tmallnr

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// Tmallnrnoticegoodsready 同步天猫配送人员信息
// tmall.nr.notice.goods.ready
//
// 接收商家的配送人员信息，和第三公司信息及提货码
func Tmallnrnoticegoodsready(clt *core.SDKClient, req *tmallnr.TmallnrnoticegoodsreadyAPIRequest, session string) (*tmallnr.TmallnrnoticegoodsreadyAPIResponse, error) {
	var resp tmallnr.TmallnrnoticegoodsreadyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
