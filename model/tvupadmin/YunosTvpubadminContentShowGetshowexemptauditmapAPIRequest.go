package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadmincontentshowgetshowexemptauditmapAPIRequest 迎客松批量查询节目某个牌照的免审状态 API请求
// yunos.tvpubadmin.content.show.getshowexemptauditmap
//
// 迎客松批量查询节目某个牌照的免审状态
type YunostvpubadmincontentshowgetshowexemptauditmapAPIRequest struct {
	model.Params
	// 节目longid
	_showLongIds string
	// 牌照id
	_license int64
}

// NewYunostvpubadmincontentshowgetshowexemptauditmapRequest 初始化YunostvpubadmincontentshowgetshowexemptauditmapAPIRequest对象
func NewYunostvpubadmincontentshowgetshowexemptauditmapRequest() *YunostvpubadmincontentshowgetshowexemptauditmapAPIRequest {
	return &YunostvpubadmincontentshowgetshowexemptauditmapAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadmincontentshowgetshowexemptauditmapAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.show.getshowexemptauditmap"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadmincontentshowgetshowexemptauditmapAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadmincontentshowgetshowexemptauditmapAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetShowLongIds is ShowLongIds Setter
// 节目longid
func (r *YunostvpubadmincontentshowgetshowexemptauditmapAPIRequest) SetShowLongIds(_showLongIds string) error {
	r._showLongIds = _showLongIds
	r.Set("show_long_ids", _showLongIds)
	return nil
}

// GetShowLongIds ShowLongIds Getter
func (r YunostvpubadmincontentshowgetshowexemptauditmapAPIRequest) GetShowLongIds() string {
	return r._showLongIds
}

// SetLicense is License Setter
// 牌照id
func (r *YunostvpubadmincontentshowgetshowexemptauditmapAPIRequest) SetLicense(_license int64) error {
	r._license = _license
	r.Set("license", _license)
	return nil
}

// GetLicense License Getter
func (r YunostvpubadmincontentshowgetshowexemptauditmapAPIRequest) GetLicense() int64 {
	return r._license
}
