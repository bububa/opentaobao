package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibababaichuanctguserrelationAPIRequest 用户 API请求
// alibaba.baichuan.ctg.user.relation
//
// 提供给优酷查询道长和淘宝账户的绑定关系
type AlibababaichuanctguserrelationAPIRequest struct {
	model.Params
	// 调用的业务方
	_app string
	// 业务方的用户ID
	_uid string
	// 淘宝的用户ID
	_tbUid string
}

// NewAlibababaichuanctguserrelationRequest 初始化AlibababaichuanctguserrelationAPIRequest对象
func NewAlibababaichuanctguserrelationRequest() *AlibababaichuanctguserrelationAPIRequest {
	return &AlibababaichuanctguserrelationAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibababaichuanctguserrelationAPIRequest) GetApiMethodName() string {
	return "alibaba.baichuan.ctg.user.relation"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibababaichuanctguserrelationAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibababaichuanctguserrelationAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApp is App Setter
// 调用的业务方
func (r *AlibababaichuanctguserrelationAPIRequest) SetApp(_app string) error {
	r._app = _app
	r.Set("app", _app)
	return nil
}

// GetApp App Getter
func (r AlibababaichuanctguserrelationAPIRequest) GetApp() string {
	return r._app
}

// SetUid is Uid Setter
// 业务方的用户ID
func (r *AlibababaichuanctguserrelationAPIRequest) SetUid(_uid string) error {
	r._uid = _uid
	r.Set("uid", _uid)
	return nil
}

// GetUid Uid Getter
func (r AlibababaichuanctguserrelationAPIRequest) GetUid() string {
	return r._uid
}

// SetTbUid is TbUid Setter
// 淘宝的用户ID
func (r *AlibababaichuanctguserrelationAPIRequest) SetTbUid(_tbUid string) error {
	r._tbUid = _tbUid
	r.Set("tb_uid", _tbUid)
	return nil
}

// GetTbUid TbUid Getter
func (r AlibababaichuanctguserrelationAPIRequest) GetTbUid() string {
	return r._tbUid
}
