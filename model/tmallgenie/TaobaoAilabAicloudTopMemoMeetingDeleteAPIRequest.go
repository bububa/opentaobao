package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAilabAicloudTopMemoMeetingDeleteAPIRequest
天猫精灵会议删除 API请求
taobao.ailab.aicloud.top.memo.meeting.delete

天猫精灵会议删除 */
type TaobaoAilabAicloudTopMemoMeetingDeleteAPIRequest struct {
	model.Params
	// schema
	_schema string
	// 企业用户ID
	_userId string
	// 手持设备ID
	_utdId string
	// 扩展信息json段，用于存放APP类型，APP版本等等信息。
	_ext string
	// 会议ID
	_memoId int64
}

// New
