package iot

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopFreelistenChildrenalbumAPIRequest 儿童音频列表 API请求
// taobao.ailab.aicloud.top.freelisten.childrenalbum
//
// 儿童音频列表
type TaobaoAilabAicloudTopFreelistenChildrenalbumAPIRequest struct {
	model.Params
	// 账户体系隔离
	_schema string
	// 用户ID，此处传入第三方账户体系的用户id
	_userId string
	// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
	_utdId string
	// 扩展信息，用于存放APP类型等
	_ext string
	// 音频类型，目前只支持以下几种类型：英语儿歌 英语故事 双语故事 国学启蒙 古典名著 成语故事 寓言故事 神话故事 诗词朗读 诗词婉唱 谚语故事 胎教音乐 经典儿歌 摇篮曲 睡前故事 绘本故事 儿童故事 儿童百科 经典故事 公主故事 名人故事 胎教故事
	_param1 string
	// 页数
	_param2 int64
	// 每页条目数
	_param3 int64
}

// NewTaobaoAilabAicloudTopFreelistenChildrenalbumRequest 初始化TaobaoAilabAicloudTopFreelistenChildrenalbumAPIRequest对象
func NewTaobaoAilabAicloudTopFreelistenChildrenalbumRequest() *TaobaoAilabAicloudTopFreelistenChildrenalbumAPIRequest {
	return &TaobaoAilabAicloudTopFreelistenChildrenalbumAPIRequest{
		Params: model.NewParams(7),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAilabAicloudTopFreelistenChildrenalbumAPIRequest) Reset() {
	r._schema = ""
	r._userId = ""
	r._utdId = ""
	r._ext = ""
	r._param1 = ""
	r._param2 = 0
	r._param3 = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopFreelistenChildrenalbumAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.freelisten.childrenalbum"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopFreelistenChildrenalbumAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAilabAicloudTopFreelistenChildrenalbumAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSchema is Schema Setter
// 账户体系隔离
func (r *TaobaoAilabAicloudTopFreelistenChildrenalbumAPIRequest) SetSchema(_schema string) error {
	r._schema = _schema
	r.Set("schema", _schema)
	return nil
}

// GetSchema Schema Getter
func (r TaobaoAilabAicloudTopFreelistenChildrenalbumAPIRequest) GetSchema() string {
	return r._schema
}

// SetUserId is UserId Setter
// 用户ID，此处传入第三方账户体系的用户id
func (r *TaobaoAilabAicloudTopFreelistenChildrenalbumAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r TaobaoAilabAicloudTopFreelistenChildrenalbumAPIRequest) GetUserId() string {
	return r._userId
}

// SetUtdId is UtdId Setter
// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
func (r *TaobaoAilabAicloudTopFreelistenChildrenalbumAPIRequest) SetUtdId(_utdId string) error {
	r._utdId = _utdId
	r.Set("utd_id", _utdId)
	return nil
}

// GetUtdId UtdId Getter
func (r TaobaoAilabAicloudTopFreelistenChildrenalbumAPIRequest) GetUtdId() string {
	return r._utdId
}

// SetExt is Ext Setter
// 扩展信息，用于存放APP类型等
func (r *TaobaoAilabAicloudTopFreelistenChildrenalbumAPIRequest) SetExt(_ext string) error {
	r._ext = _ext
	r.Set("ext", _ext)
	return nil
}

// GetExt Ext Getter
func (r TaobaoAilabAicloudTopFreelistenChildrenalbumAPIRequest) GetExt() string {
	return r._ext
}

// SetParam1 is Param1 Setter
// 音频类型，目前只支持以下几种类型：英语儿歌 英语故事 双语故事 国学启蒙 古典名著 成语故事 寓言故事 神话故事 诗词朗读 诗词婉唱 谚语故事 胎教音乐 经典儿歌 摇篮曲 睡前故事 绘本故事 儿童故事 儿童百科 经典故事 公主故事 名人故事 胎教故事
func (r *TaobaoAilabAicloudTopFreelistenChildrenalbumAPIRequest) SetParam1(_param1 string) error {
	r._param1 = _param1
	r.Set("param1", _param1)
	return nil
}

// GetParam1 Param1 Getter
func (r TaobaoAilabAicloudTopFreelistenChildrenalbumAPIRequest) GetParam1() string {
	return r._param1
}

// SetParam2 is Param2 Setter
// 页数
func (r *TaobaoAilabAicloudTopFreelistenChildrenalbumAPIRequest) SetParam2(_param2 int64) error {
	r._param2 = _param2
	r.Set("param2", _param2)
	return nil
}

// GetParam2 Param2 Getter
func (r TaobaoAilabAicloudTopFreelistenChildrenalbumAPIRequest) GetParam2() int64 {
	return r._param2
}

// SetParam3 is Param3 Setter
// 每页条目数
func (r *TaobaoAilabAicloudTopFreelistenChildrenalbumAPIRequest) SetParam3(_param3 int64) error {
	r._param3 = _param3
	r.Set("param3", _param3)
	return nil
}

// GetParam3 Param3 Getter
func (r TaobaoAilabAicloudTopFreelistenChildrenalbumAPIRequest) GetParam3() int64 {
	return r._param3
}

var poolTaobaoAilabAicloudTopFreelistenChildrenalbumAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAilabAicloudTopFreelistenChildrenalbumRequest()
	},
}

// GetTaobaoAilabAicloudTopFreelistenChildrenalbumRequest 从 sync.Pool 获取 TaobaoAilabAicloudTopFreelistenChildrenalbumAPIRequest
func GetTaobaoAilabAicloudTopFreelistenChildrenalbumAPIRequest() *TaobaoAilabAicloudTopFreelistenChildrenalbumAPIRequest {
	return poolTaobaoAilabAicloudTopFreelistenChildrenalbumAPIRequest.Get().(*TaobaoAilabAicloudTopFreelistenChildrenalbumAPIRequest)
}

// ReleaseTaobaoAilabAicloudTopFreelistenChildrenalbumAPIRequest 将 TaobaoAilabAicloudTopFreelistenChildrenalbumAPIRequest 放入 sync.Pool
func ReleaseTaobaoAilabAicloudTopFreelistenChildrenalbumAPIRequest(v *TaobaoAilabAicloudTopFreelistenChildrenalbumAPIRequest) {
	v.Reset()
	poolTaobaoAilabAicloudTopFreelistenChildrenalbumAPIRequest.Put(v)
}
