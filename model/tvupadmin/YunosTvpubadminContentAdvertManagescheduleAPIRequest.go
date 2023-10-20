package tvupadmin

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosTvpubadminContentAdvertManagescheduleAPIRequest 广告牌照管控修改 API请求
// yunos.tvpubadmin.content.advert.manageschedule
//
// 广告牌照管控修改
type YunosTvpubadminContentAdvertManagescheduleAPIRequest struct {
	model.Params
	// 管理参数
	_req string
}

// NewYunosTvpubadminContentAdvertManagescheduleRequest 初始化YunosTvpubadminContentAdvertManagescheduleAPIRequest对象
func NewYunosTvpubadminContentAdvertManagescheduleRequest() *YunosTvpubadminContentAdvertManagescheduleAPIRequest {
	return &YunosTvpubadminContentAdvertManagescheduleAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosTvpubadminContentAdvertManagescheduleAPIRequest) Reset() {
	r._req = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosTvpubadminContentAdvertManagescheduleAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.advert.manageschedule"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosTvpubadminContentAdvertManagescheduleAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosTvpubadminContentAdvertManagescheduleAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReq is Req Setter
// 管理参数
func (r *YunosTvpubadminContentAdvertManagescheduleAPIRequest) SetReq(_req string) error {
	r._req = _req
	r.Set("req", _req)
	return nil
}

// GetReq Req Getter
func (r YunosTvpubadminContentAdvertManagescheduleAPIRequest) GetReq() string {
	return r._req
}

var poolYunosTvpubadminContentAdvertManagescheduleAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosTvpubadminContentAdvertManagescheduleRequest()
	},
}

// GetYunosTvpubadminContentAdvertManagescheduleRequest 从 sync.Pool 获取 YunosTvpubadminContentAdvertManagescheduleAPIRequest
func GetYunosTvpubadminContentAdvertManagescheduleAPIRequest() *YunosTvpubadminContentAdvertManagescheduleAPIRequest {
	return poolYunosTvpubadminContentAdvertManagescheduleAPIRequest.Get().(*YunosTvpubadminContentAdvertManagescheduleAPIRequest)
}

// ReleaseYunosTvpubadminContentAdvertManagescheduleAPIRequest 将 YunosTvpubadminContentAdvertManagescheduleAPIRequest 放入 sync.Pool
func ReleaseYunosTvpubadminContentAdvertManagescheduleAPIRequest(v *YunosTvpubadminContentAdvertManagescheduleAPIRequest) {
	v.Reset()
	poolYunosTvpubadminContentAdvertManagescheduleAPIRequest.Put(v)
}
