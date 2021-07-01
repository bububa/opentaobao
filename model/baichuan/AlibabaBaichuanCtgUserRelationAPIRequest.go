package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaBaichuanCtgUserRelationAPIRequest
用户 API请求
alibaba.baichuan.ctg.user.relation

提供给优酷查询道长和淘宝账户的绑定关系 */
type AlibabaBaichuanCtgUserRelationAPIRequest struct {
	model.Params
	// 调用的业务方
	_app string
	// 业务方的用户ID
	_uid string
	// 淘宝的用户ID
	_tbUid string
}

// NewAlibabaBaichuanCtgUserRelationRequest 初始化AlibabaBaichuanCtgUserRelationAPIRequest对象
func NewAlibabaBaichuanCtgUserRelationRequest() *AlibabaBaichuanCtgUserRelationAPIRequest {
	return &AlibabaBaichuanCtgUserRelationAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaBaichuanCtgUserRelationAPIRequest) GetApiMethodName() string {
	return "alibaba.baichuan.ctg.user.relation"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaBaichuanCtgUserRelationAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is App Setter
// 调用的业务方
func (r *AlibabaBaichuanCtgUserRelationAPIRequest) SetApp(_app string) error {
	r._app = _app
	r.Set("app", _app)
	return nil
}

// Get App Getter
func (r AlibabaBaichuanCtgUserRelationAPIRequest) GetApp() string {
	return r._app
}

// Set is Uid Setter
// 业务方的用户ID
func (r *AlibabaBaichuanCtgUserRelationAPIRequest) SetUid(_uid string) error {
	r._uid = _uid
	r.Set("uid", _uid)
	return nil
}

// Get Uid Getter
func (r AlibabaBaichuanCtgUserRelationAPIRequest) GetUid() string {
	return r._uid
}

// Set is TbUid Setter
// 淘宝的用户ID
func (r *AlibabaBaichuanCtgUserRelationAPIRequest) SetTbUid(_tbUid string) error {
	r._tbUid = _tbUid
	r.Set("tb_uid", _tbUid)
	return nil
}

// Get TbUid Getter
func (r AlibabaBaichuanCtgUserRelationAPIRequest) GetTbUid() string {
	return r._tbUid
}
