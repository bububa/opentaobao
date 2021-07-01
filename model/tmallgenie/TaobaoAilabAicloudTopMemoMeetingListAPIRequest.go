package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAilabAicloudTopMemoMeetingListAPIRequest
天猫精灵会议查询 API请求
taobao.ailab.aicloud.top.memo.meeting.list

查询天猫精灵用户设置的所有会议 */
type TaobaoAilabAicloudTopMemoMeetingListAPIRequest struct {
	model.Params
	// schema
	_schema string
	// 企业用户ID
	_userId string
	// 手持设备ID
	_utdId string
	// 扩展信息json段，用于存放APP类型，APP版本等等信息。
	_ext string
	// 闹钟ID
	_memoId int64
}

// New
