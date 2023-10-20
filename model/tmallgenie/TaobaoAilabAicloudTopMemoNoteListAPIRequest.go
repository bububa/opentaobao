package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoailabaicloudtopmemonotelistAPIRequest 查询天猫精灵用户设置的所有备忘录 API请求
// taobao.ailab.aicloud.top.memo.note.list
//
// 查询天猫精灵用户设置的所有备忘录
type TaobaoailabaicloudtopmemonotelistAPIRequest struct {
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

// NewTaobaoailabaicloudtopmemonotelistRequest 初始化TaobaoailabaicloudtopmemonotelistAPIRequest对象
func NewTaobaoailabaicloudtopmemonotelistRequest() *TaobaoailabaicloudtopmemonotelistAPIRequest {
	return &TaobaoailabaicloudtopmemonotelistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoailabaicloudtopmemonotelistAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.memo.note.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoailabaicloudtopmemonotelistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoailabaicloudtopmemonotelistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSchema is Schema Setter
// schema
func (r *TaobaoailabaicloudtopmemonotelistAPIRequest) SetSchema(_schema string) error {
	r._schema = _schema
	r.Set("schema", _schema)
	return nil
}

// GetSchema Schema Getter
func (r TaobaoailabaicloudtopmemonotelistAPIRequest) GetSchema() string {
	return r._schema
}

// SetUserId is UserId Setter
// 企业用户ID
func (r *TaobaoailabaicloudtopmemonotelistAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r TaobaoailabaicloudtopmemonotelistAPIRequest) GetUserId() string {
	return r._userId
}

// SetUtdId is UtdId Setter
// 手持设备ID
func (r *TaobaoailabaicloudtopmemonotelistAPIRequest) SetUtdId(_utdId string) error {
	r._utdId = _utdId
	r.Set("utd_id", _utdId)
	return nil
}

// GetUtdId UtdId Getter
func (r TaobaoailabaicloudtopmemonotelistAPIRequest) GetUtdId() string {
	return r._utdId
}

// SetExt is Ext Setter
// 扩展信息json段，用于存放APP类型，APP版本等等信息。
func (r *TaobaoailabaicloudtopmemonotelistAPIRequest) SetExt(_ext string) error {
	r._ext = _ext
	r.Set("ext", _ext)
	return nil
}

// GetExt Ext Getter
func (r TaobaoailabaicloudtopmemonotelistAPIRequest) GetExt() string {
	return r._ext
}

// SetMemoId is MemoId Setter
// 备忘录ID
func (r *TaobaoailabaicloudtopmemonotelistAPIRequest) SetMemoId(_memoId int64) error {
	r._memoId = _memoId
	r.Set("memo_id", _memoId)
	return nil
}

// GetMemoId MemoId Getter
func (r TaobaoailabaicloudtopmemonotelistAPIRequest) GetMemoId() int64 {
	return r._memoId
}
