package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoUmpDetailUpdateAPIRequest
修改活动详情 API请求
taobao.ump.detail.update

更新活动详情 */
type TaobaoUmpDetailUpdateAPIRequest struct {
	model.Params
	// 活动详情id
	_detailId int64
	// 活动详情内容，可以通过ump sdk中的MarketingTool来生成这个内容
	_content string
}

// New
