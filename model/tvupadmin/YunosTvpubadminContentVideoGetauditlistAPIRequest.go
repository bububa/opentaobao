package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadmincontentvideogetauditlistAPIRequest 迎客松视频审核记录查询 API请求
// yunos.tvpubadmin.content.video.getauditlist
//
// 迎客松视频审核记录查询
type YunostvpubadmincontentvideogetauditlistAPIRequest struct {
	model.Params
	// 查询
	_query string
}

// NewYunostvpubadmincontentvideogetauditlistRequest 初始化YunostvpubadmincontentvideogetauditlistAPIRequest对象
func NewYunostvpubadmincontentvideogetauditlistRequest() *YunostvpubadmincontentvideogetauditlistAPIRequest {
	return &YunostvpubadmincontentvideogetauditlistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadmincontentvideogetauditlistAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.video.getauditlist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadmincontentvideogetauditlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadmincontentvideogetauditlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetQuery is Query Setter
// 查询
func (r *YunostvpubadmincontentvideogetauditlistAPIRequest) SetQuery(_query string) error {
	r._query = _query
	r.Set("query", _query)
	return nil
}

// GetQuery Query Getter
func (r YunostvpubadmincontentvideogetauditlistAPIRequest) GetQuery() string {
	return r._query
}
