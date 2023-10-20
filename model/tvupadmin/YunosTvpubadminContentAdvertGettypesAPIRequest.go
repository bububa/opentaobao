package tvupadmin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunostvpubadmincontentadvertgettypesAPIRequest 获取广告位类型 API请求
// yunos.tvpubadmin.content.advert.gettypes
//
// 获取广告位类型
type YunostvpubadmincontentadvertgettypesAPIRequest struct {
	model.Params
}

// NewYunostvpubadmincontentadvertgettypesRequest 初始化YunostvpubadmincontentadvertgettypesAPIRequest对象
func NewYunostvpubadmincontentadvertgettypesRequest() *YunostvpubadmincontentadvertgettypesAPIRequest {
	return &YunostvpubadmincontentadvertgettypesAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunostvpubadmincontentadvertgettypesAPIRequest) GetApiMethodName() string {
	return "yunos.tvpubadmin.content.advert.gettypes"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunostvpubadmincontentadvertgettypesAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunostvpubadmincontentadvertgettypesAPIRequest) GetRawParams() model.Params {
	return r.Params
}
