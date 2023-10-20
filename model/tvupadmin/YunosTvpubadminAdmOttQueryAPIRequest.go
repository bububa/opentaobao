package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadminadmottqueryAPIRequest 优酷OTT端广告素材查询 API请求
// yunos.tvpubadmin.adm.ott.query
//
// 查询广告素材
type YunostvpubadminadmottqueryAPIRequest struct {
	model.Params
	// 查询参数json格式
	_query string
}

// NewYunostvpubadminadmottqueryRequest 初始化YunostvpubadminadmottqueryAPIRequest对象
func NewYunostvpubadminadmottqueryRequest() *YunostvpubadminadmottqueryAPIRequest {
	return &YunostvpubadminadmottqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadminadmottqueryAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.adm.ott.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadminadmottqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadminadmottqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuery is Query Setter
// 查询参数json格式
func (r *YunostvpubadminadmottqueryAPIRequest) SetQuery(_query string) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r YunostvpubadminadmottqueryAPIRequest) GetQuery() string {
	return r._query
}
