package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoailabaicloudtopmemomeetingdeleteAPIRequest 天猫精灵会议删除 API请求
// taobao.ailab.aicloud.top.memo.meeting.delete
//
// 天猫精灵会议删除
type TaobaoailabaicloudtopmemomeetingdeleteAPIRequest struct {
	model.Params
	// schema
	_schema string
	// 企业用户ID
	_userId string
	// 手持设备ID
	_utdId string
	// 扩展信息json段，用于存放APP类型，APP版本等等信息。
	_ext string
	// 会议ID
	_memoId int64
}

// NewTaobaoailabaicloudtopmemomeetingdeleteRequest 初始化TaobaoailabaicloudtopmemomeetingdeleteAPIRequest对象
func NewTaobaoailabaicloudtopmemomeetingdeleteRequest() *TaobaoailabaicloudtopmemomeetingdeleteAPIRequest {
	return &TaobaoailabaicloudtopmemomeetingdeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoailabaicloudtopmemomeetingdeleteAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.memo.meeting.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoailabaicloudtopmemomeetingdeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoailabaicloudtopmemomeetingdeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSchema is Schema Setter
// schema
func (r *TaobaoailabaicloudtopmemomeetingdeleteAPIRequest) SetSchema(_schema string) error {
	r._schema = _schema
	r.Set("schema", _schema)
	return nil
}

// GetSchema Schema Getter
func (r TaobaoailabaicloudtopmemomeetingdeleteAPIRequest) GetSchema() string {
	return r._schema
}

// SetUserId is UserId Setter
// 企业用户ID
func (r *TaobaoailabaicloudtopmemomeetingdeleteAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r TaobaoailabaicloudtopmemomeetingdeleteAPIRequest) GetUserId() string {
	return r._userId
}

// SetUtdId is UtdId Setter
// 手持设备ID
func (r *TaobaoailabaicloudtopmemomeetingdeleteAPIRequest) SetUtdId(_utdId string) error {
	r._utdId = _utdId
	r.Set("utd_id", _utdId)
	return nil
}

// GetUtdId UtdId Getter
func (r TaobaoailabaicloudtopmemomeetingdeleteAPIRequest) GetUtdId() string {
	return r._utdId
}

// SetExt is Ext Setter
// 扩展信息json段，用于存放APP类型，APP版本等等信息。
func (r *TaobaoailabaicloudtopmemomeetingdeleteAPIRequest) SetExt(_ext string) error {
	r._ext = _ext
	r.Set("ext", _ext)
	return nil
}

// GetExt Ext Getter
func (r TaobaoailabaicloudtopmemomeetingdeleteAPIRequest) GetExt() string {
	return r._ext
}

// SetMemoId is MemoId Setter
// 会议ID
func (r *TaobaoailabaicloudtopmemomeetingdeleteAPIRequest) SetMemoId(_memoId int64) error {
	r._memoId = _memoId
	r.Set("memo_id", _memoId)
	return nil
}

// GetMemoId MemoId Getter
func (r TaobaoailabaicloudtopmemomeetingdeleteAPIRequest) GetMemoId() int64 {
	return r._memoId
}
