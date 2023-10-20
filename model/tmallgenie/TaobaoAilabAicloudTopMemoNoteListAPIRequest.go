package tmallgenie

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopMemoNoteListAPIRequest 查询天猫精灵用户设置的所有备忘录 API请求
// taobao.ailab.aicloud.top.memo.note.list
//
// 查询天猫精灵用户设置的所有备忘录
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
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAilabAicloudTopMemoNoteListAPIRequest) Reset() {
	r._schema = ""
	r._userId = ""
	r._utdId = ""
	r._ext = ""
	r._memoId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopMemoNoteListAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.memo.note.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopMemoNoteListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAilabAicloudTopMemoNoteListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSchema is Schema Setter
// schema
func (r *TaobaoAilabAicloudTopMemoNoteListAPIRequest) SetSchema(_schema string) error {
	r._schema = _schema
	r.Set("schema", _schema)
	return nil
}

// GetSchema Schema Getter
func (r TaobaoAilabAicloudTopMemoNoteListAPIRequest) GetSchema() string {
	return r._schema
}

// SetUserId is UserId Setter
// 企业用户ID
func (r *TaobaoAilabAicloudTopMemoNoteListAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r TaobaoAilabAicloudTopMemoNoteListAPIRequest) GetUserId() string {
	return r._userId
}

// SetUtdId is UtdId Setter
// 手持设备ID
func (r *TaobaoAilabAicloudTopMemoNoteListAPIRequest) SetUtdId(_utdId string) error {
	r._utdId = _utdId
	r.Set("utd_id", _utdId)
	return nil
}

// GetUtdId UtdId Getter
func (r TaobaoAilabAicloudTopMemoNoteListAPIRequest) GetUtdId() string {
	return r._utdId
}

// SetExt is Ext Setter
// 扩展信息json段，用于存放APP类型，APP版本等等信息。
func (r *TaobaoAilabAicloudTopMemoNoteListAPIRequest) SetExt(_ext string) error {
	r._ext = _ext
	r.Set("ext", _ext)
	return nil
}

// GetExt Ext Getter
func (r TaobaoAilabAicloudTopMemoNoteListAPIRequest) GetExt() string {
	return r._ext
}

// SetMemoId is MemoId Setter
// 备忘录ID
func (r *TaobaoAilabAicloudTopMemoNoteListAPIRequest) SetMemoId(_memoId int64) error {
	r._memoId = _memoId
	r.Set("memo_id", _memoId)
	return nil
}

// GetMemoId MemoId Getter
func (r TaobaoAilabAicloudTopMemoNoteListAPIRequest) GetMemoId() int64 {
	return r._memoId
}

var poolTaobaoAilabAicloudTopMemoNoteListAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAilabAicloudTopMemoNoteListRequest()
	},
}

// GetTaobaoAilabAicloudTopMemoNoteListRequest 从 sync.Pool 获取 TaobaoAilabAicloudTopMemoNoteListAPIRequest
func GetTaobaoAilabAicloudTopMemoNoteListAPIRequest() *TaobaoAilabAicloudTopMemoNoteListAPIRequest {
	return poolTaobaoAilabAicloudTopMemoNoteListAPIRequest.Get().(*TaobaoAilabAicloudTopMemoNoteListAPIRequest)
}

// ReleaseTaobaoAilabAicloudTopMemoNoteListAPIRequest 将 TaobaoAilabAicloudTopMemoNoteListAPIRequest 放入 sync.Pool
func ReleaseTaobaoAilabAicloudTopMemoNoteListAPIRequest(v *TaobaoAilabAicloudTopMemoNoteListAPIRequest) {
	v.Reset()
	poolTaobaoAilabAicloudTopMemoNoteListAPIRequest.Put(v)
}
