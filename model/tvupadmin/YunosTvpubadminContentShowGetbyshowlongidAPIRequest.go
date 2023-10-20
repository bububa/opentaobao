package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadmincontentshowgetbyshowlongidAPIRequest 迎客松根据节目longid获取节目元数据 API请求
// yunos.tvpubadmin.content.show.getbyshowlongid
//
// 迎客松根据节目longid获取节目元数据
type YunostvpubadmincontentshowgetbyshowlongidAPIRequest struct {
	model.Params
	// 节目longid
	_showLongId int64
}

// NewYunostvpubadmincontentshowgetbyshowlongidRequest 初始化YunostvpubadmincontentshowgetbyshowlongidAPIRequest对象
func NewYunostvpubadmincontentshowgetbyshowlongidRequest() *YunostvpubadmincontentshowgetbyshowlongidAPIRequest {
	return &YunostvpubadmincontentshowgetbyshowlongidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadmincontentshowgetbyshowlongidAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.show.getbyshowlongid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadmincontentshowgetbyshowlongidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadmincontentshowgetbyshowlongidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetShowLongId is ShowLongId Setter
// 节目longid
func (r *YunostvpubadmincontentshowgetbyshowlongidAPIRequest) SetShowLongId(_showLongId int64) error {
	r._showLongId = _showLongId
	r.Set("show_long_id", _showLongId)
	return nil
}

// GetShowLongId ShowLongId Getter
func (r YunostvpubadmincontentshowgetbyshowlongidAPIRequest) GetShowLongId() int64 {
	return r._showLongId
}
