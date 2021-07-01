package fans

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallFansArenaRecordAPIRequest
记录完成擂台的用户 API请求
tmall.fans.arena.record

记录完成擂台的用户和完成分数 */
type TmallFansArenaRecordAPIRequest struct {
	model.Params
	// 资金池id
	_cashPoolId int64
	// 用户得分
	_score int64
	// mixnick
	_mixNick string
}

// New
