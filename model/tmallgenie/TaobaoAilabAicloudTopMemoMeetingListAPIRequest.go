package tmallgenie

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopMemoMeetingListAPIRequest 天猫精灵会议查询 API请求
// taobao.ailab.aicloud.top.memo.meeting.list
//
// 查询天猫精灵用户设置的所有会议
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

// NewTaobaoAilabAicloudTopMemoMeetingListRequest 初始化TaobaoAilabAicloudTopMemoMeetingListAPIRequest对象
func NewTaobaoAilabAicloudTopMemoMeetingListRequest() *TaobaoAilabAicloudTopMemoMeetingListAPIRequest {
	return &TaobaoAilabAicloudTopMemoMeetingListAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAilabAicloudTopMemoMeetingListAPIRequest) Reset() {
	r._schema = ""
	r._userId = ""
	r._utdId = ""
	r._ext = ""
	r._memoId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopMemoMeetingListAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.memo.meeting.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopMemoMeetingListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAilabAicloudTopMemoMeetingListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSchema is Schema Setter
// schema
func (r *TaobaoAilabAicloudTopMemoMeetingListAPIRequest) SetSchema(_schema string) error {
	r._schema = _schema
	r.Set("schema", _schema)
	return nil
}

// GetSchema Schema Getter
func (r TaobaoAilabAicloudTopMemoMeetingListAPIRequest) GetSchema() string {
	return r._schema
}

// SetUserId is UserId Setter
// 企业用户ID
func (r *TaobaoAilabAicloudTopMemoMeetingListAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r TaobaoAilabAicloudTopMemoMeetingListAPIRequest) GetUserId() string {
	return r._userId
}

// SetUtdId is UtdId Setter
// 手持设备ID
func (r *TaobaoAilabAicloudTopMemoMeetingListAPIRequest) SetUtdId(_utdId string) error {
	r._utdId = _utdId
	r.Set("utd_id", _utdId)
	return nil
}

// GetUtdId UtdId Getter
func (r TaobaoAilabAicloudTopMemoMeetingListAPIRequest) GetUtdId() string {
	return r._utdId
}

// SetExt is Ext Setter
// 扩展信息json段，用于存放APP类型，APP版本等等信息。
func (r *TaobaoAilabAicloudTopMemoMeetingListAPIRequest) SetExt(_ext string) error {
	r._ext = _ext
	r.Set("ext", _ext)
	return nil
}

// GetExt Ext Getter
func (r TaobaoAilabAicloudTopMemoMeetingListAPIRequest) GetExt() string {
	return r._ext
}

// SetMemoId is MemoId Setter
// 闹钟ID
func (r *TaobaoAilabAicloudTopMemoMeetingListAPIRequest) SetMemoId(_memoId int64) error {
	r._memoId = _memoId
	r.Set("memo_id", _memoId)
	return nil
}

// GetMemoId MemoId Getter
func (r TaobaoAilabAicloudTopMemoMeetingListAPIRequest) GetMemoId() int64 {
	return r._memoId
}

var poolTaobaoAilabAicloudTopMemoMeetingListAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAilabAicloudTopMemoMeetingListRequest()
	},
}

// GetTaobaoAilabAicloudTopMemoMeetingListRequest 从 sync.Pool 获取 TaobaoAilabAicloudTopMemoMeetingListAPIRequest
func GetTaobaoAilabAicloudTopMemoMeetingListAPIRequest() *TaobaoAilabAicloudTopMemoMeetingListAPIRequest {
	return poolTaobaoAilabAicloudTopMemoMeetingListAPIRequest.Get().(*TaobaoAilabAicloudTopMemoMeetingListAPIRequest)
}

// ReleaseTaobaoAilabAicloudTopMemoMeetingListAPIRequest 将 TaobaoAilabAicloudTopMemoMeetingListAPIRequest 放入 sync.Pool
func ReleaseTaobaoAilabAicloudTopMemoMeetingListAPIRequest(v *TaobaoAilabAicloudTopMemoMeetingListAPIRequest) {
	v.Reset()
	poolTaobaoAilabAicloudTopMemoMeetingListAPIRequest.Put(v)
}
