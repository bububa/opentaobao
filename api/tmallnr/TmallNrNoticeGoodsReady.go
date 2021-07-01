package tmallnr

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

/* TmallNrNoticeGoodsReady
同步天猫配送人员信息
tmall.nr.notice.goods.ready

接收商家的配送人员信息，和第三公司信息及提货码 */
func TmallNrNoticeGoodsReady(clt *core.SDKClient, req *tmallnr.TmallNrNoticeGoodsReadyAPIRequest, session string) (*tmallnr.TmallNrNoticeGoodsReadyAPIResponse, error) {
	var resp tmallnr.TmallNrNoticeGoodsReadyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
