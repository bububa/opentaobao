package user

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpenAccountIndexFindAPIRequest Open Account索引查询 API请求
// taobao.open.account.index.find
//
// Open Account索引查询
type TaobaoOpenAccountIndexFindAPIRequest struct {
	model.Params
	// 具体值，当索引类型是 OPEN_ID 是，格式为 oauthPlatform|openId，即使用竖线分隔的组合值
	_indexValue string
	// int MOBILE         = 1;int EMAIL          = 2;int ISV_ACCOUNT_ID = 3;int LOGIN_ID       = 4;int OPEN_ID        = 5;
	_indexType int64
}

// NewTaobaoOpenAccountIndexFindRequest 初始化TaobaoOpenAccountIndexFindAPIRequest对象
func NewTaobaoOpenAccountIndexFindRequest() *TaobaoOpenAccountIndexFindAPIRequest {
	return &TaobaoOpenAccountIndexFindAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoOpenAccountIndexFindAPIRequest) Reset() {
	r._indexValue = ""
	r._indexType = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenAccountIndexFindAPIRequest) GetApiMethodName() string {
	return "taobao.open.account.index.find"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenAccountIndexFindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoOpenAccountIndexFindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIndexValue is IndexValue Setter
// 具体值，当索引类型是 OPEN_ID 是，格式为 oauthPlatform|openId，即使用竖线分隔的组合值
func (r *TaobaoOpenAccountIndexFindAPIRequest) SetIndexValue(_indexValue string) error {
	r._indexValue = _indexValue
	r.Set("index_value", _indexValue)
	return nil
}

// GetIndexValue IndexValue Getter
func (r TaobaoOpenAccountIndexFindAPIRequest) GetIndexValue() string {
	return r._indexValue
}

// SetIndexType is IndexType Setter
// int MOBILE         = 1;int EMAIL          = 2;int ISV_ACCOUNT_ID = 3;int LOGIN_ID       = 4;int OPEN_ID        = 5;
func (r *TaobaoOpenAccountIndexFindAPIRequest) SetIndexType(_indexType int64) error {
	r._indexType = _indexType
	r.Set("index_type", _indexType)
	return nil
}

// GetIndexType IndexType Getter
func (r TaobaoOpenAccountIndexFindAPIRequest) GetIndexType() int64 {
	return r._indexType
}

var poolTaobaoOpenAccountIndexFindAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoOpenAccountIndexFindRequest()
	},
}

// GetTaobaoOpenAccountIndexFindRequest 从 sync.Pool 获取 TaobaoOpenAccountIndexFindAPIRequest
func GetTaobaoOpenAccountIndexFindAPIRequest() *TaobaoOpenAccountIndexFindAPIRequest {
	return poolTaobaoOpenAccountIndexFindAPIRequest.Get().(*TaobaoOpenAccountIndexFindAPIRequest)
}

// ReleaseTaobaoOpenAccountIndexFindAPIRequest 将 TaobaoOpenAccountIndexFindAPIRequest 放入 sync.Pool
func ReleaseTaobaoOpenAccountIndexFindAPIRequest(v *TaobaoOpenAccountIndexFindAPIRequest) {
	v.Reset()
	poolTaobaoOpenAccountIndexFindAPIRequest.Put(v)
}
