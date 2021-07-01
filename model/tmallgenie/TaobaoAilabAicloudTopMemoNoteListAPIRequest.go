package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAilabAicloudTopMemoNoteListAPIRequest
查询天猫精灵用户设置的所有备忘录 API请求
taobao.ailab.aicloud.top.memo.note.list

查询天猫精灵用户设置的所有备忘录 */
type TaobaoAilabAicloudTopMemoNoteListAPIRequest struct {
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

// NewTaobaoAilabAicloudTopMemoNoteListRequest 初始化TaobaoAilabAicloudTopMemoNoteListAPIRequest对象
func NewTaobaoAilabAicloudTopMemoNoteListRequest() *TaobaoAilabAicloudTopMemoNoteListAPIRequest {
	return &TaobaoAilabAicloudTopMemoNoteListAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopMemoNoteListAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.memo.note.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopMemoNoteListAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Schema Setter
// schema
func (r *TaobaoAilabAicloudTopMemoNoteListAPIRequest) SetSchema(_schema string) error {
	r._schema = _schema
	r.Set("schema", _schema)
	return nil
}

// Get Schema Getter
func (r TaobaoAilabAicloudTopMemoNoteListAPIRequest) GetSchema() string {
	return r._schema
}

// Set is UserId Setter
// 企业用户ID
func (r *TaobaoAilabAicloudTopMemoNoteListAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// Get UserId Getter
func (r TaobaoAilabAicloudTopMemoNoteListAPIRequest) GetUserId() string {
	return r._userId
}

// Set is UtdId Setter
// 手持设备ID
func (r *TaobaoAilabAicloudTopMemoNoteListAPIRequest) SetUtdId(_utdId string) error {
	r._utdId = _utdId
	r.Set("utd_id", _utdId)
	return nil
}

// Get UtdId Getter
func (r TaobaoAilabAicloudTopMemoNoteListAPIRequest) GetUtdId() string {
	return r._utdId
}

// Set is Ext Setter
// 扩展信息json段，用于存放APP类型，APP版本等等信息。
func (r *TaobaoAilabAicloudTopMemoNoteListAPIRequest) SetExt(_ext string) error {
	r._ext = _ext
	r.Set("ext", _ext)
	return nil
}

// Get Ext Getter
func (r TaobaoAilabAicloudTopMemoNoteListAPIRequest) GetExt() string {
	return r._ext
}

// Set is MemoId Setter
// 备忘录ID
func (r *TaobaoAilabAicloudTopMemoNoteListAPIRequest) SetMemoId(_memoId int64) error {
	r._memoId = _memoId
	r.Set("memo_id", _memoId)
	return nil
}

// Get MemoId Getter
func (r TaobaoAilabAicloudTopMemoNoteListAPIRequest) GetMemoId() int64 {
	return r._memoId
}
