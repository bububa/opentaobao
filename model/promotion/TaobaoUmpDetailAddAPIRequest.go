package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoUmpDetailAddAPIRequest
新增活动详情 API请求
taobao.ump.detail.add

增加活动详情。活动详情里面包括活动的范围（店铺，商品），活动的参数（比如具体的折扣），参与类型（全部，部分，部分不参加）等信息。当参与类型为部分或部分不参加的时候需要和taobao.ump.range.add来配合使用。 */
type TaobaoUmpDetailAddAPIRequest struct {
	model.Params
	// 增加工具详情
	_actId int64
	// 活动详情内容，json格式，可以通过ump sdk中的MarketingTool来进行处理
	_content string
}

// New
