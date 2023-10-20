package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadmincontentappquerybylicenceAPIRequest 按牌照查询应用 API请求
// yunos.tvpubadmin.content.app.querybylicence
//
// 按牌照查询应用
type YunostvpubadmincontentappquerybylicenceAPIRequest struct {
	model.Params
	// 查询条件
	_query string
}

// NewYunostvpubadmincontentappquerybylicenceRequest 初始化YunostvpubadmincontentappquerybylicenceAPIRequest对象
func NewYunostvpubadmincontentappquerybylicenceRequest() *YunostvpubadmincontentappquerybylicenceAPIRequest {
	return &YunostvpubadmincontentappquerybylicenceAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadmincontentappquerybylicenceAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.app.querybylicence"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadmincontentappquerybylicenceAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadmincontentappquerybylicenceAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuery is Query Setter
// 查询条件
func (r *YunostvpubadmincontentappquerybylicenceAPIRequest) SetQuery(_query string) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r YunostvpubadmincontentappquerybylicenceAPIRequest) GetQuery() string {
	return r._query
}
