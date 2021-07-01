package traderate

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFliggyWrateGetmixratelistAPIRequest
飞猪通用评价接口 API请求
taobao.fliggy.wrate.getmixratelist

飞猪评价通用接口 */
type TaobaoFliggyWrateGetmixratelistAPIRequest struct {
	model.Params
	// 评论列表查询参数
	_paramTopGetMixRateListParam *TopGetMixRateListParam
}

// New
