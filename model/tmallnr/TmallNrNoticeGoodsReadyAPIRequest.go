package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallNrNoticeGoodsReadyAPIRequest
同步天猫配送人员信息 API请求
tmall.nr.notice.goods.ready

接收商家的配送人员信息，和第三公司信息及提货码 */
type TmallNrNoticeGoodsReadyAPIRequest struct {
	model.Params
	// 入参
	_param0 *NrZqsGoodsReadyReqDto
}

// New
