package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoUmpActivityUpdateAPIRequest
修改活动信息 API请求
taobao.ump.activity.update

修改营销活动 */
type TaobaoUmpActivityUpdateAPIRequest struct {
	model.Params
	// 活动id
	_actId int64
	// 营销活动内容，json格式，通过ump sdk 的marketingTool来生成
	_content string
}

// New
