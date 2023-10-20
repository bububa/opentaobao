package tvupadmin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminContentShowGetbyshowlongidAPIRequest 迎客松根据节目longid获取节目元数据 API请求
// yunos.tvpubadmin.content.show.getbyshowlongid
//
// 迎客松根据节目longid获取节目元数据
type YunosTvpubadminContentShowGetbyshowlongidAPIRequest struct {
	model.Params
	// 节目longid
	_showLongId int64
}

// NewYunosTvpubadminContentShowGetbyshowlongidRequest 初始化YunosTvpubadminContentShowGetbyshowlongidAPIRequest对象
func NewYunosTvpubadminContentShowGetbyshowlongidRequest() *YunosTvpubadminContentShowGetbyshowlongidAPIRequest {
	return &YunosTvpubadminContentShowGetbyshowlongidAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosTvpubadminContentShowGetbyshowlongidAPIRequest) Reset() {
	r._showLongId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentShowGetbyshowlongidAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.show.getbyshowlongid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentShowGetbyshowlongidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminContentShowGetbyshowlongidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetShowLongId is ShowLongId Setter
// 节目longid
func (r *YunosTvpubadminContentShowGetbyshowlongidAPIRequest) SetShowLongId(_showLongId int64) error {
	r._showLongId = _showLongId
	r.Set("show_long_id", _showLongId)
	return nil
}

// GetShowLongId ShowLongId Getter
func (r YunosTvpubadminContentShowGetbyshowlongidAPIRequest) GetShowLongId() int64 {
	return r._showLongId
}

var poolYunosTvpubadminContentShowGetbyshowlongidAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosTvpubadminContentShowGetbyshowlongidRequest()
	},
}

// GetYunosTvpubadminContentShowGetbyshowlongidRequest 从 sync.Pool 获取 YunosTvpubadminContentShowGetbyshowlongidAPIRequest
func GetYunosTvpubadminContentShowGetbyshowlongidAPIRequest() *YunosTvpubadminContentShowGetbyshowlongidAPIRequest {
	return poolYunosTvpubadminContentShowGetbyshowlongidAPIRequest.Get().(*YunosTvpubadminContentShowGetbyshowlongidAPIRequest)
}

// ReleaseYunosTvpubadminContentShowGetbyshowlongidAPIRequest 将 YunosTvpubadminContentShowGetbyshowlongidAPIRequest 放入 sync.Pool
func ReleaseYunosTvpubadminContentShowGetbyshowlongidAPIRequest(v *YunosTvpubadminContentShowGetbyshowlongidAPIRequest) {
	v.Reset()
	poolYunosTvpubadminContentShowGetbyshowlongidAPIRequest.Put(v)
}
