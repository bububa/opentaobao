package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbkscucrowdgetAPIRequest 淘宝客-服务商-获取人群标签 API请求
// taobao.tbk.sc.ucrowd.get
//
// 服务商使用。支持淘宝客通过入参人群标签id，获得人群信息，包括人群名称描述及覆盖会员数。
type TaobaotbkscucrowdgetAPIRequest struct {
	model.Params
	// 人群标签id
	_ucrowdId int64
}

// NewTaobaotbkscucrowdgetRequest 初始化TaobaotbkscucrowdgetAPIRequest对象
func NewTaobaotbkscucrowdgetRequest() *TaobaotbkscucrowdgetAPIRequest {
	return &TaobaotbkscucrowdgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotbkscucrowdgetAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.sc.ucrowd.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotbkscucrowdgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotbkscucrowdgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUcrowdId is UcrowdId Setter
// 人群标签id
func (r *TaobaotbkscucrowdgetAPIRequest) SetUcrowdId(_ucrowdId int64) error {
	r._ucrowdId = _ucrowdId
	r.Set("ucrowd_id", _ucrowdId)
	return nil
}

// GetUcrowdId UcrowdId Getter
func (r TaobaotbkscucrowdgetAPIRequest) GetUcrowdId() int64 {
	return r._ucrowdId
}
