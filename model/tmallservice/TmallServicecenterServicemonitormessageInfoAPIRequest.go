package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenterservicemonitormessageinfoAPIRequest 服务预警单查询接口 API请求
// tmall.servicecenter.servicemonitormessage.info
//
// 服务预警单查询接口,用于查询预警单最新状态
type TmallservicecenterservicemonitormessageinfoAPIRequest struct {
	model.Params
	// 预警单ID
	_id int64
}

// NewTmallservicecenterservicemonitormessageinfoRequest 初始化TmallservicecenterservicemonitormessageinfoAPIRequest对象
func NewTmallservicecenterservicemonitormessageinfoRequest() *TmallservicecenterservicemonitormessageinfoAPIRequest {
	return &TmallservicecenterservicemonitormessageinfoAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicecenterservicemonitormessageinfoAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.servicemonitormessage.info"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicecenterservicemonitormessageinfoAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicecenterservicemonitormessageinfoAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetId is Id Setter
// 预警单ID
func (r *TmallservicecenterservicemonitormessageinfoAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r TmallservicecenterservicemonitormessageinfoAPIRequest) GetId() int64 {
	return r._id
}
