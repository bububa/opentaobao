package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminAdmOttQueryAPIRequest 优酷OTT端广告素材查询 API请求
// yunos.tvpubadmin.adm.ott.query
//
// 查询广告素材
type YunosTvpubadminAdmOttQueryAPIRequest struct {
	model.Params
	// 查询参数json格式
	_query string
}

// NewYunosTvpubadminAdmOttQueryRequest 初始化YunosTvpubadminAdmOttQueryAPIRequest对象
func NewYunosTvpubadminAdmOttQueryRequest() *YunosTvpubadminAdmOttQueryAPIRequest {
	return &YunosTvpubadminAdmOttQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminAdmOttQueryAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.adm.ott.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminAdmOttQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminAdmOttQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuery is Query Setter
// 查询参数json格式
func (r *YunosTvpubadminAdmOttQueryAPIRequest) SetQuery(_query string) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r YunosTvpubadminAdmOttQueryAPIRequest) GetQuery() string {
	return r._query
}
