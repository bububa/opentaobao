package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadmincontentadvertmanagescheduleAPIRequest 广告牌照管控修改 API请求
// yunos.tvpubadmin.content.advert.manageschedule
//
// 广告牌照管控修改
type YunostvpubadmincontentadvertmanagescheduleAPIRequest struct {
	model.Params
	// 管理参数
	_req string
}

// NewYunostvpubadmincontentadvertmanagescheduleRequest 初始化YunostvpubadmincontentadvertmanagescheduleAPIRequest对象
func NewYunostvpubadmincontentadvertmanagescheduleRequest() *YunostvpubadmincontentadvertmanagescheduleAPIRequest {
	return &YunostvpubadmincontentadvertmanagescheduleAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadmincontentadvertmanagescheduleAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.advert.manageschedule"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadmincontentadvertmanagescheduleAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadmincontentadvertmanagescheduleAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReq is Req Setter
// 管理参数
func (r *YunostvpubadmincontentadvertmanagescheduleAPIRequest) SetReq(_req string) error {
	r._req = _req
	r.Set("req", _req)
	return nil
}

// GetReq Req Getter
func (r YunostvpubadmincontentadvertmanagescheduleAPIRequest) GetReq() string {
	return r._req
}
