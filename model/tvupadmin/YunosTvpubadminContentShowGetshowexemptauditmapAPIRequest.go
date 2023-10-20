package tvupadmin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminContentShowGetshowexemptauditmapAPIRequest 迎客松批量查询节目某个牌照的免审状态 API请求
// yunos.tvpubadmin.content.show.getshowexemptauditmap
//
// 迎客松批量查询节目某个牌照的免审状态
type YunosTvpubadminContentShowGetshowexemptauditmapAPIRequest struct {
	model.Params
	// 节目longid
	_showLongIds string
	// 牌照id
	_license int64
}

// NewYunosTvpubadminContentShowGetshowexemptauditmapRequest 初始化YunosTvpubadminContentShowGetshowexemptauditmapAPIRequest对象
func NewYunosTvpubadminContentShowGetshowexemptauditmapRequest() *YunosTvpubadminContentShowGetshowexemptauditmapAPIRequest {
	return &YunosTvpubadminContentShowGetshowexemptauditmapAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosTvpubadminContentShowGetshowexemptauditmapAPIRequest) Reset() {
	r._showLongIds = ""
	r._license = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentShowGetshowexemptauditmapAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.show.getshowexemptauditmap"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentShowGetshowexemptauditmapAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminContentShowGetshowexemptauditmapAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetShowLongIds is ShowLongIds Setter
// 节目longid
func (r *YunosTvpubadminContentShowGetshowexemptauditmapAPIRequest) SetShowLongIds(_showLongIds string) error {
	r._showLongIds = _showLongIds
	r.Set("show_long_ids", _showLongIds)
	return nil
}

// GetShowLongIds ShowLongIds Getter
func (r YunosTvpubadminContentShowGetshowexemptauditmapAPIRequest) GetShowLongIds() string {
	return r._showLongIds
}

// SetLicense is License Setter
// 牌照id
func (r *YunosTvpubadminContentShowGetshowexemptauditmapAPIRequest) SetLicense(_license int64) error {
	r._license = _license
	r.Set("license", _license)
	return nil
}

// GetLicense License Getter
func (r YunosTvpubadminContentShowGetshowexemptauditmapAPIRequest) GetLicense() int64 {
	return r._license
}

var poolYunosTvpubadminContentShowGetshowexemptauditmapAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosTvpubadminContentShowGetshowexemptauditmapRequest()
	},
}

// GetYunosTvpubadminContentShowGetshowexemptauditmapRequest 从 sync.Pool 获取 YunosTvpubadminContentShowGetshowexemptauditmapAPIRequest
func GetYunosTvpubadminContentShowGetshowexemptauditmapAPIRequest() *YunosTvpubadminContentShowGetshowexemptauditmapAPIRequest {
	return poolYunosTvpubadminContentShowGetshowexemptauditmapAPIRequest.Get().(*YunosTvpubadminContentShowGetshowexemptauditmapAPIRequest)
}

// ReleaseYunosTvpubadminContentShowGetshowexemptauditmapAPIRequest 将 YunosTvpubadminContentShowGetshowexemptauditmapAPIRequest 放入 sync.Pool
func ReleaseYunosTvpubadminContentShowGetshowexemptauditmapAPIRequest(v *YunosTvpubadminContentShowGetshowexemptauditmapAPIRequest) {
	v.Reset()
	poolYunosTvpubadminContentShowGetshowexemptauditmapAPIRequest.Put(v)
}
