package baichuan

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaBaichuanCtgUserRelationAPIRequest 用户 API请求
// alibaba.baichuan.ctg.user.relation
//
// 提供给优酷查询道长和淘宝账户的绑定关系
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
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaBaichuanCtgUserRelationAPIRequest) Reset() {
	r._app = ""
	r._uid = ""
	r._tbUid = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaBaichuanCtgUserRelationAPIRequest) GetApiMethodName() string {
	return "alibaba.baichuan.ctg.user.relation"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaBaichuanCtgUserRelationAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaBaichuanCtgUserRelationAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApp is App Setter
// 调用的业务方
func (r *AlibabaBaichuanCtgUserRelationAPIRequest) SetApp(_app string) error {
	r._app = _app
	r.Set("app", _app)
	return nil
}

// GetApp App Getter
func (r AlibabaBaichuanCtgUserRelationAPIRequest) GetApp() string {
	return r._app
}

// SetUid is Uid Setter
// 业务方的用户ID
func (r *AlibabaBaichuanCtgUserRelationAPIRequest) SetUid(_uid string) error {
	r._uid = _uid
	r.Set("uid", _uid)
	return nil
}

// GetUid Uid Getter
func (r AlibabaBaichuanCtgUserRelationAPIRequest) GetUid() string {
	return r._uid
}

// SetTbUid is TbUid Setter
// 淘宝的用户ID
func (r *AlibabaBaichuanCtgUserRelationAPIRequest) SetTbUid(_tbUid string) error {
	r._tbUid = _tbUid
	r.Set("tb_uid", _tbUid)
	return nil
}

// GetTbUid TbUid Getter
func (r AlibabaBaichuanCtgUserRelationAPIRequest) GetTbUid() string {
	return r._tbUid
}

var poolAlibabaBaichuanCtgUserRelationAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaBaichuanCtgUserRelationRequest()
	},
}

// GetAlibabaBaichuanCtgUserRelationRequest 从 sync.Pool 获取 AlibabaBaichuanCtgUserRelationAPIRequest
func GetAlibabaBaichuanCtgUserRelationAPIRequest() *AlibabaBaichuanCtgUserRelationAPIRequest {
	return poolAlibabaBaichuanCtgUserRelationAPIRequest.Get().(*AlibabaBaichuanCtgUserRelationAPIRequest)
}

// ReleaseAlibabaBaichuanCtgUserRelationAPIRequest 将 AlibabaBaichuanCtgUserRelationAPIRequest 放入 sync.Pool
func ReleaseAlibabaBaichuanCtgUserRelationAPIRequest(v *AlibabaBaichuanCtgUserRelationAPIRequest) {
	v.Reset()
	poolAlibabaBaichuanCtgUserRelationAPIRequest.Put(v)
}
