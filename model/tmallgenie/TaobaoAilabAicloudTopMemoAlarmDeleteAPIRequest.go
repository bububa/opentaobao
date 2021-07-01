package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest
天猫精灵闹钟删除 API请求
taobao.ailab.aicloud.top.memo.alarm.delete

天猫精灵闹钟删除 */
type TaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest struct {
	model.Params
	// schema
	_schema string
	// 手持设备ID
	_utdId string
	// 扩展信息json段，用于存放APP类型，APP版本等等信息。
	_ext string
	// 企业用户ID
	_userId string
	// 闹钟ID
	_memoId int64
}

// New
