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

// NewTaobaoAilabAicloudTopMemoNoteDeleteRequest 初始化TaobaoAilabAicloudTopMemoNoteDeleteAPIRequest对象
func NewTaobaoAilabAicloudTopMemoNoteDeleteRequest() *TaobaoAilabAicloudTopMemoNoteDeleteAPIRequest {
	return &TaobaoAilabAicloudTopMemoNoteDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopMemoNoteDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.memo.note.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopMemoNoteDeleteAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Schema Setter
// schema
func (r *TaobaoAilabAicloudTopMemoNoteDeleteAPIRequest) SetSchema(_schema string) error {
	r._schema = _schema
	r.Set("schema", _schema)
	return nil
}

// Get Schema Getter
func (r TaobaoAilabAicloudTopMemoNoteDeleteAPIRequest) GetSchema() string {
	return r._schema
}

// Set is UserId Setter
// 企业用户ID
func (r *TaobaoAilabAicloudTopMemoNoteDeleteAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// Get UserId Getter
func (r TaobaoAilabAicloudTopMemoNoteDeleteAPIRequest) GetUserId() string {
	return r._userId
}

// Set is UtdId Setter
// 手持设备ID
func (r *TaobaoAilabAicloudTopMemoNoteDeleteAPIRequest) SetUtdId(_utdId string) error {
	r._utdId = _utdId
	r.Set("utd_id", _utdId)
	return nil
}

// Get UtdId Getter
func (r TaobaoAilabAicloudTopMemoNoteDeleteAPIRequest) GetUtdId() string {
	return r._utdId
}

// Set is Ext Setter
// 扩展信息json段，用于存放APP类型，APP版本等等信息。
func (r *TaobaoAilabAicloudTopMemoNoteDeleteAPIRequest) SetExt(_ext string) error {
	r._ext = _ext
	r.Set("ext", _ext)
	return nil
}

// Get Ext Getter
func (r TaobaoAilabAicloudTopMemoNoteDeleteAPIRequest) GetExt() string {
	return r._ext
}

// Set is MemoId Setter
// 备忘录ID
func (r *TaobaoAilabAicloudTopMemoNoteDeleteAPIRequest) SetMemoId(_memoId int64) error {
	r._memoId = _memoId
	r.Set("memo_id", _memoId)
	return nil
}

// Get MemoId Getter
func (r TaobaoAilabAicloudTopMemoNoteDeleteAPIRequest) GetMemoId() int64 {
	return r._memoId
}
