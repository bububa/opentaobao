package baichuanctg

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaBaichuanCtgToutiaoContentAPIRequest
微博输出头条数据 API请求
alibaba.baichuan.ctg.toutiao.content

百川头条内容获取 */
type AlibabaBaichuanCtgToutiaoContentAPIRequest struct {
	model.Params
	// param0
	_param0 *CtgRequest
}

// New
