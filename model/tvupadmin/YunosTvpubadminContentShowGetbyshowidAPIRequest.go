package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadmincontentshowgetbyshowidAPIRequest 迎客松根据节目id获取节目元数据 API请求
// yunos.tvpubadmin.content.show.getbyshowid
//
// 迎客松根据节目id获取节目元数据
type YunostvpubadmincontentshowgetbyshowidAPIRequest struct {
	model.Params
	// 节目字符串id
	_showId string
}

// NewYunostvpubadmincontentshowgetbyshowidRequest 初始化YunostvpubadmincontentshowgetbyshowidAPIRequest对象
func NewYunostvpubadmincontentshowgetbyshowidRequest() *YunostvpubadmincontentshowgetbyshowidAPIRequest {
	return &YunostvpubadmincontentshowgetbyshowidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadmincontentshowgetbyshowidAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.show.getbyshowid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadmincontentshowgetbyshowidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadmincontentshowgetbyshowidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetShowId is ShowId Setter
// 节目字符串id
func (r *YunostvpubadmincontentshowgetbyshowidAPIRequest) SetShowId(_showId string) error {
	r._showId = _showId
	r.Set("show_id", _showId)
	return nil
}

// GetShowId ShowId Getter
func (r YunostvpubadmincontentshowgetbyshowidAPIRequest) GetShowId() string {
	return r._showId
}
