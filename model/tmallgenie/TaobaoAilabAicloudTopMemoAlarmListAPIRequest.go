package tmallgenie

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopMemoAlarmListAPIRequest 天猫精灵闹钟查询 API请求
// taobao.ailab.aicloud.top.memo.alarm.list
//
// 查询天猫精灵用户设置的所有闹钟
type TaobaoAilabAicloudTopMemoAlarmListAPIRequest struct {
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

// NewTaobaoAilabAicloudTopMemoAlarmListRequest 初始化TaobaoAilabAicloudTopMemoAlarmListAPIRequest对象
func NewTaobaoAilabAicloudTopMemoAlarmListRequest() *TaobaoAilabAicloudTopMemoAlarmListAPIRequest {
	return &TaobaoAilabAicloudTopMemoAlarmListAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAilabAicloudTopMemoAlarmListAPIRequest) Reset() {
	r._schema = ""
	r._userId = ""
	r._utdId = ""
	r._ext = ""
	r._memoId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopMemoAlarmListAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.memo.alarm.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopMemoAlarmListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAilabAicloudTopMemoAlarmListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSchema is Schema Setter
// schema
func (r *TaobaoAilabAicloudTopMemoAlarmListAPIRequest) SetSchema(_schema string) error {
	r._schema = _schema
	r.Set("schema", _schema)
	return nil
}

// GetSchema Schema Getter
func (r TaobaoAilabAicloudTopMemoAlarmListAPIRequest) GetSchema() string {
	return r._schema
}

// SetUserId is UserId Setter
// 企业用户ID
func (r *TaobaoAilabAicloudTopMemoAlarmListAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r TaobaoAilabAicloudTopMemoAlarmListAPIRequest) GetUserId() string {
	return r._userId
}

// SetUtdId is UtdId Setter
// 手持设备ID
func (r *TaobaoAilabAicloudTopMemoAlarmListAPIRequest) SetUtdId(_utdId string) error {
	r._utdId = _utdId
	r.Set("utd_id", _utdId)
	return nil
}

// GetUtdId UtdId Getter
func (r TaobaoAilabAicloudTopMemoAlarmListAPIRequest) GetUtdId() string {
	return r._utdId
}

// SetExt is Ext Setter
// 扩展信息json段，用于存放APP类型，APP版本等等信息。
func (r *TaobaoAilabAicloudTopMemoAlarmListAPIRequest) SetExt(_ext string) error {
	r._ext = _ext
	r.Set("ext", _ext)
	return nil
}

// GetExt Ext Getter
func (r TaobaoAilabAicloudTopMemoAlarmListAPIRequest) GetExt() string {
	return r._ext
}

// SetMemoId is MemoId Setter
// 闹钟ID
func (r *TaobaoAilabAicloudTopMemoAlarmListAPIRequest) SetMemoId(_memoId int64) error {
	r._memoId = _memoId
	r.Set("memo_id", _memoId)
	return nil
}

// GetMemoId MemoId Getter
func (r TaobaoAilabAicloudTopMemoAlarmListAPIRequest) GetMemoId() int64 {
	return r._memoId
}

var poolTaobaoAilabAicloudTopMemoAlarmListAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAilabAicloudTopMemoAlarmListRequest()
	},
}

// GetTaobaoAilabAicloudTopMemoAlarmListRequest 从 sync.Pool 获取 TaobaoAilabAicloudTopMemoAlarmListAPIRequest
func GetTaobaoAilabAicloudTopMemoAlarmListAPIRequest() *TaobaoAilabAicloudTopMemoAlarmListAPIRequest {
	return poolTaobaoAilabAicloudTopMemoAlarmListAPIRequest.Get().(*TaobaoAilabAicloudTopMemoAlarmListAPIRequest)
}

// ReleaseTaobaoAilabAicloudTopMemoAlarmListAPIRequest 将 TaobaoAilabAicloudTopMemoAlarmListAPIRequest 放入 sync.Pool
func ReleaseTaobaoAilabAicloudTopMemoAlarmListAPIRequest(v *TaobaoAilabAicloudTopMemoAlarmListAPIRequest) {
	v.Reset()
	poolTaobaoAilabAicloudTopMemoAlarmListAPIRequest.Put(v)
}
