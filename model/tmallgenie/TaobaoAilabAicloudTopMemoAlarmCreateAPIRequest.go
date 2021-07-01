package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAilabAicloudTopMemoAlarmCreateAPIRequest
天猫精灵闹钟创建 API请求
taobao.ailab.aicloud.top.memo.alarm.create

天猫精灵闹钟创建 */
type TaobaoAilabAicloudTopMemoAlarmCreateAPIRequest struct {
	model.Params
	// 扩展信息json段，用于存放APP类型，APP版本等等信息。
	_ext string
	// schema
	_schema string
	// 企业用户ID
	_userId string
	// 手持设备ID
	_utdId string
	// 创建闹钟入参
	_paramCreateAlarmParam *CreateAlarmParam
}

// New
