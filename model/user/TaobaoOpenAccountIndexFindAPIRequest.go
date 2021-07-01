package user

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoOpenAccountIndexFindAPIRequest
Open Account索引查询 API请求
taobao.open.account.index.find

Open Account索引查询 */
type TaobaoOpenAccountIndexFindAPIRequest struct {
	model.Params
	// int MOBILE         = 1;int EMAIL          = 2;int ISV_ACCOUNT_ID = 3;int LOGIN_ID       = 4;int OPEN_ID        = 5;
	_indexType int64
	// 具体值，当索引类型是 OPEN_ID 是，格式为 oauthPlatform|openId，即使用竖线分隔的组合值
	_indexValue string
}

// NewTaobaoOpenAccountIndexFindRequest 初始化TaobaoOpenAccountIndexFindAPIRequest对象
func NewTaobaoOpenAccountIndexFindRequest() *TaobaoOpenAccountIndexFindAPIRequest {
	return &TaobaoOpenAccountIndexFindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoOpenAccountIndexFindAPIRequest) GetApiMethodName() string {
	return "taobao.open.account.index.find"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoOpenAccountIndexFindAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is IndexType Setter
// int MOBILE         = 1;int EMAIL          = 2;int ISV_ACCOUNT_ID = 3;int LOGIN_ID       = 4;int OPEN_ID        = 5;
func (r *TaobaoOpenAccountIndexFindAPIRequest) SetIndexType(_indexType int64) error {
	r._indexType = _indexType
	r.Set("index_type", _indexType)
	return nil
}

// Get IndexType Getter
func (r TaobaoOpenAccountIndexFindAPIRequest) GetIndexType() int64 {
	return r._indexType
}

// Set is IndexValue Setter
// 具体值，当索引类型是 OPEN_ID 是，格式为 oauthPlatform|openId，即使用竖线分隔的组合值
func (r *TaobaoOpenAccountIndexFindAPIRequest) SetIndexValue(_indexValue string) error {
	r._indexValue = _indexValue
	r.Set("index_value", _indexValue)
	return nil
}

// Get IndexValue Getter
func (r TaobaoOpenAccountIndexFindAPIRequest) GetIndexValue() string {
	return r._indexValue
}
