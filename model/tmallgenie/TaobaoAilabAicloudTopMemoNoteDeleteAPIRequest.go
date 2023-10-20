package tmallgenie

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopMemoNoteDeleteAPIRequest 天猫精灵备忘录删除 API请求
// taobao.ailab.aicloud.top.memo.note.delete
//
// 删除天猫精灵用户设置的备忘录
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
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAilabAicloudTopMemoNoteDeleteAPIRequest) Reset() {
	r._schema = ""
	r._userId = ""
	r._utdId = ""
	r._ext = ""
	r._memoId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopMemoNoteDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.memo.note.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopMemoNoteDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAilabAicloudTopMemoNoteDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSchema is Schema Setter
// schema
func (r *TaobaoAilabAicloudTopMemoNoteDeleteAPIRequest) SetSchema(_schema string) error {
	r._schema = _schema
	r.Set("schema", _schema)
	return nil
}

// GetSchema Schema Getter
func (r TaobaoAilabAicloudTopMemoNoteDeleteAPIRequest) GetSchema() string {
	return r._schema
}

// SetUserId is UserId Setter
// 企业用户ID
func (r *TaobaoAilabAicloudTopMemoNoteDeleteAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r TaobaoAilabAicloudTopMemoNoteDeleteAPIRequest) GetUserId() string {
	return r._userId
}

// SetUtdId is UtdId Setter
// 手持设备ID
func (r *TaobaoAilabAicloudTopMemoNoteDeleteAPIRequest) SetUtdId(_utdId string) error {
	r._utdId = _utdId
	r.Set("utd_id", _utdId)
	return nil
}

// GetUtdId UtdId Getter
func (r TaobaoAilabAicloudTopMemoNoteDeleteAPIRequest) GetUtdId() string {
	return r._utdId
}

// SetExt is Ext Setter
// 扩展信息json段，用于存放APP类型，APP版本等等信息。
func (r *TaobaoAilabAicloudTopMemoNoteDeleteAPIRequest) SetExt(_ext string) error {
	r._ext = _ext
	r.Set("ext", _ext)
	return nil
}

// GetExt Ext Getter
func (r TaobaoAilabAicloudTopMemoNoteDeleteAPIRequest) GetExt() string {
	return r._ext
}

// SetMemoId is MemoId Setter
// 备忘录ID
func (r *TaobaoAilabAicloudTopMemoNoteDeleteAPIRequest) SetMemoId(_memoId int64) error {
	r._memoId = _memoId
	r.Set("memo_id", _memoId)
	return nil
}

// GetMemoId MemoId Getter
func (r TaobaoAilabAicloudTopMemoNoteDeleteAPIRequest) GetMemoId() int64 {
	return r._memoId
}

var poolTaobaoAilabAicloudTopMemoNoteDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAilabAicloudTopMemoNoteDeleteRequest()
	},
}

// GetTaobaoAilabAicloudTopMemoNoteDeleteRequest 从 sync.Pool 获取 TaobaoAilabAicloudTopMemoNoteDeleteAPIRequest
func GetTaobaoAilabAicloudTopMemoNoteDeleteAPIRequest() *TaobaoAilabAicloudTopMemoNoteDeleteAPIRequest {
	return poolTaobaoAilabAicloudTopMemoNoteDeleteAPIRequest.Get().(*TaobaoAilabAicloudTopMemoNoteDeleteAPIRequest)
}

// ReleaseTaobaoAilabAicloudTopMemoNoteDeleteAPIRequest 将 TaobaoAilabAicloudTopMemoNoteDeleteAPIRequest 放入 sync.Pool
func ReleaseTaobaoAilabAicloudTopMemoNoteDeleteAPIRequest(v *TaobaoAilabAicloudTopMemoNoteDeleteAPIRequest) {
	v.Reset()
	poolTaobaoAilabAicloudTopMemoNoteDeleteAPIRequest.Put(v)
}
