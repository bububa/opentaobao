package tmallgenie

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest 天猫精灵闹钟删除 API请求
// taobao.ailab.aicloud.top.memo.alarm.delete
//
// 天猫精灵闹钟删除
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

// NewTaobaoAilabAicloudTopMemoAlarmDeleteRequest 初始化TaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest对象
func NewTaobaoAilabAicloudTopMemoAlarmDeleteRequest() *TaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest {
	return &TaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest) Reset() {
	r._schema = ""
	r._utdId = ""
	r._ext = ""
	r._userId = ""
	r._memoId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.memo.alarm.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSchema is Schema Setter
// schema
func (r *TaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest) SetSchema(_schema string) error {
	r._schema = _schema
	r.Set("schema", _schema)
	return nil
}

// GetSchema Schema Getter
func (r TaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest) GetSchema() string {
	return r._schema
}

// SetUtdId is UtdId Setter
// 手持设备ID
func (r *TaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest) SetUtdId(_utdId string) error {
	r._utdId = _utdId
	r.Set("utd_id", _utdId)
	return nil
}

// GetUtdId UtdId Getter
func (r TaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest) GetUtdId() string {
	return r._utdId
}

// SetExt is Ext Setter
// 扩展信息json段，用于存放APP类型，APP版本等等信息。
func (r *TaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest) SetExt(_ext string) error {
	r._ext = _ext
	r.Set("ext", _ext)
	return nil
}

// GetExt Ext Getter
func (r TaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest) GetExt() string {
	return r._ext
}

// SetUserId is UserId Setter
// 企业用户ID
func (r *TaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r TaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest) GetUserId() string {
	return r._userId
}

// SetMemoId is MemoId Setter
// 闹钟ID
func (r *TaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest) SetMemoId(_memoId int64) error {
	r._memoId = _memoId
	r.Set("memo_id", _memoId)
	return nil
}

// GetMemoId MemoId Getter
func (r TaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest) GetMemoId() int64 {
	return r._memoId
}

var poolTaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAilabAicloudTopMemoAlarmDeleteRequest()
	},
}

// GetTaobaoAilabAicloudTopMemoAlarmDeleteRequest 从 sync.Pool 获取 TaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest
func GetTaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest() *TaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest {
	return poolTaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest.Get().(*TaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest)
}

// ReleaseTaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest 将 TaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest 放入 sync.Pool
func ReleaseTaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest(v *TaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest) {
	v.Reset()
	poolTaobaoAilabAicloudTopMemoAlarmDeleteAPIRequest.Put(v)
}
