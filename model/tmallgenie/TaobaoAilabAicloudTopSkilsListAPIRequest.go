package tmallgenie

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAilabAicloudTopSkilsListAPIRequest 获取硬件平台设备下挂载的技能列表 API请求
// taobao.ailab.aicloud.top.skils.list
//
// 提供给在硬件平台接入设备的技能列表
type TaobaoAilabAicloudTopSkilsListAPIRequest struct {
	model.Params
	// 账户体系隔离
	_schema string
	// 用户ID，此处传入第三方账户体系的用户id
	_userId string
	// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
	_utdId string
	// 扩展信息，用于存放APP类型等
	_ext string
	// query(模糊匹配skillName)
	_query string
	// type(1000代表内容技能，3000代表自定义技能，4000代表官方技能)
	_type string
	// pageNo
	_pageNo int64
	// pageSize
	_pageSize int64
}

// NewTaobaoAilabAicloudTopSkilsListRequest 初始化TaobaoAilabAicloudTopSkilsListAPIRequest对象
func NewTaobaoAilabAicloudTopSkilsListRequest() *TaobaoAilabAicloudTopSkilsListAPIRequest {
	return &TaobaoAilabAicloudTopSkilsListAPIRequest{
		Params: model.NewParams(8),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAilabAicloudTopSkilsListAPIRequest) Reset() {
	r._schema = ""
	r._userId = ""
	r._utdId = ""
	r._ext = ""
	r._query = ""
	r._type = ""
	r._pageNo = 0
	r._pageSize = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAilabAicloudTopSkilsListAPIRequest) GetApiMethodName() string {
	return "taobao.ailab.aicloud.top.skils.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAilabAicloudTopSkilsListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAilabAicloudTopSkilsListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSchema is Schema Setter
// 账户体系隔离
func (r *TaobaoAilabAicloudTopSkilsListAPIRequest) SetSchema(_schema string) error {
	r._schema = _schema
	r.Set("schema", _schema)
	return nil
}

// GetSchema Schema Getter
func (r TaobaoAilabAicloudTopSkilsListAPIRequest) GetSchema() string {
	return r._schema
}

// SetUserId is UserId Setter
// 用户ID，此处传入第三方账户体系的用户id
func (r *TaobaoAilabAicloudTopSkilsListAPIRequest) SetUserId(_userId string) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r TaobaoAilabAicloudTopSkilsListAPIRequest) GetUserId() string {
	return r._userId
}

// SetUtdId is UtdId Setter
// 用户设备唯一识别码，长度限制32以内，建议使用系统接口获取deviceid,然后做一定的混淆处理来作为此输入参数
func (r *TaobaoAilabAicloudTopSkilsListAPIRequest) SetUtdId(_utdId string) error {
	r._utdId = _utdId
	r.Set("utd_id", _utdId)
	return nil
}

// GetUtdId UtdId Getter
func (r TaobaoAilabAicloudTopSkilsListAPIRequest) GetUtdId() string {
	return r._utdId
}

// SetExt is Ext Setter
// 扩展信息，用于存放APP类型等
func (r *TaobaoAilabAicloudTopSkilsListAPIRequest) SetExt(_ext string) error {
	r._ext = _ext
	r.Set("ext", _ext)
	return nil
}

// GetExt Ext Getter
func (r TaobaoAilabAicloudTopSkilsListAPIRequest) GetExt() string {
	return r._ext
}

// SetQuery is Query Setter
// query(模糊匹配skillName)
func (r *TaobaoAilabAicloudTopSkilsListAPIRequest) SetQuery(_query string) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r TaobaoAilabAicloudTopSkilsListAPIRequest) GetQuery() string {
	return r._query
}

// SetType is Type Setter
// type(1000代表内容技能，3000代表自定义技能，4000代表官方技能)
func (r *TaobaoAilabAicloudTopSkilsListAPIRequest) SetType(_type string) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TaobaoAilabAicloudTopSkilsListAPIRequest) GetType() string {
	return r._type
}

// SetPageNo is PageNo Setter
// pageNo
func (r *TaobaoAilabAicloudTopSkilsListAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r TaobaoAilabAicloudTopSkilsListAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// pageSize
func (r *TaobaoAilabAicloudTopSkilsListAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r TaobaoAilabAicloudTopSkilsListAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

var poolTaobaoAilabAicloudTopSkilsListAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAilabAicloudTopSkilsListRequest()
	},
}

// GetTaobaoAilabAicloudTopSkilsListRequest 从 sync.Pool 获取 TaobaoAilabAicloudTopSkilsListAPIRequest
func GetTaobaoAilabAicloudTopSkilsListAPIRequest() *TaobaoAilabAicloudTopSkilsListAPIRequest {
	return poolTaobaoAilabAicloudTopSkilsListAPIRequest.Get().(*TaobaoAilabAicloudTopSkilsListAPIRequest)
}

// ReleaseTaobaoAilabAicloudTopSkilsListAPIRequest 将 TaobaoAilabAicloudTopSkilsListAPIRequest 放入 sync.Pool
func ReleaseTaobaoAilabAicloudTopSkilsListAPIRequest(v *TaobaoAilabAicloudTopSkilsListAPIRequest) {
	v.Reset()
	poolTaobaoAilabAicloudTopSkilsListAPIRequest.Put(v)
}
