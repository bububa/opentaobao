package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotbkscucrowddeleteAPIRequest 淘宝客-服务商-删除人群标签 API请求
// taobao.tbk.sc.ucrowd.delete
//
// 服务商使用。支持淘宝客删除人群标签id，被删除的人群标签id将失效，并不可重新生效。
type TaobaotbkscucrowddeleteAPIRequest struct {
	model.Params
	// 人群标签id
	_ucrowdId int64
}

// NewTaobaotbkscucrowddeleteRequest 初始化TaobaotbkscucrowddeleteAPIRequest对象
func NewTaobaotbkscucrowddeleteRequest() *TaobaotbkscucrowddeleteAPIRequest {
	return &TaobaotbkscucrowddeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotbkscucrowddeleteAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.sc.ucrowd.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotbkscucrowddeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotbkscucrowddeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUcrowdId is UcrowdId Setter
// 人群标签id
func (r *TaobaotbkscucrowddeleteAPIRequest) SetUcrowdId(_ucrowdId int64) error {
	r._ucrowdId = _ucrowdId
	r.Set("ucrowd_id", _ucrowdId)
	return nil
}

// GetUcrowdId UcrowdId Getter
func (r TaobaotbkscucrowddeleteAPIRequest) GetUcrowdId() int64 {
	return r._ucrowdId
}
