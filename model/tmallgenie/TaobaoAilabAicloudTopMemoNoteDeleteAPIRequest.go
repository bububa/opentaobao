package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAilabAicloudTopMemoNoteDeleteAPIRequest
天猫精灵备忘录删除 API请求
taobao.ailab.aicloud.top.memo.note.delete

删除天猫精灵用户设置的备忘录 */
type TaobaoAilabAicloudTopMemoNoteDeleteAPIRequest struct {
	model.Params
	// schema
	_schema string
	// 企业用户ID
	_userId string
	// 手持设备ID
	_utdId string
	// 扩展信息json段，用于存放APP类型，APP版本等等信息。
	_ext string
	// 备忘录ID
	_memoId int64
}

// New
