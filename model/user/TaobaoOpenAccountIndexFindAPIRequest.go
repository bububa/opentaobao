package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopenaccountindexfindAPIRequest Open Account索引查询 API请求
// taobao.open.account.index.find
//
// Open Account索引查询
type TaobaoopenaccountindexfindAPIRequest struct {
	model.Params
	// 具体值，当索引类型是 OPEN_ID 是，格式为 oauthPlatform|openId，即使用竖线分隔的组合值
	_indexValue string
	// int MOBILE         = 1;int EMAIL          = 2;int ISV_ACCOUNT_ID = 3;int LOGIN_ID       = 4;int OPEN_ID        = 5;
	_indexType int64
}

// NewTaobaoopenaccountindexfindRequest 初始化TaobaoopenaccountindexfindAPIRequest对象
func NewTaobaoopenaccountindexfindRequest() *TaobaoopenaccountindexfindAPIRequest {
	return &TaobaoopenaccountindexfindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoopenaccountindexfindAPIRequest) GetApiMethodName() string {
	return "taobao.open.account.index.find"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoopenaccountindexfindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoopenaccountindexfindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetIndexValue is IndexValue Setter
// 具体值，当索引类型是 OPEN_ID 是，格式为 oauthPlatform|openId，即使用竖线分隔的组合值
func (r *TaobaoopenaccountindexfindAPIRequest) SetIndexValue(_indexValue string) error {
	r._indexValue = _indexValue
	r.Set("index_value", _indexValue)
	return nil
}

// GetIndexValue IndexValue Getter
func (r TaobaoopenaccountindexfindAPIRequest) GetIndexValue() string {
	return r._indexValue
}

// SetIndexType is IndexType Setter
// int MOBILE         = 1;int EMAIL          = 2;int ISV_ACCOUNT_ID = 3;int LOGIN_ID       = 4;int OPEN_ID        = 5;
func (r *TaobaoopenaccountindexfindAPIRequest) SetIndexType(_indexType int64) error {
	r._indexType = _indexType
	r.Set("index_type", _indexType)
	return nil
}

// GetIndexType IndexType Getter
func (r TaobaoopenaccountindexfindAPIRequest) GetIndexType() int64 {
	return r._indexType
}
