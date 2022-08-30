package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkScUcrowdDeleteAPIRequest 淘宝客-服务商-删除人群标签 API请求
// taobao.tbk.sc.ucrowd.delete
//
// 服务商使用。支持淘宝客删除人群标签id，被删除的人群标签id将失效，并不可重新生效。
type TaobaoTbkScUcrowdDeleteAPIRequest struct {
	model.Params
	// 人群标签id
	_ucrowdId int64
}

// NewTaobaoTbkScUcrowdDeleteRequest 初始化TaobaoTbkScUcrowdDeleteAPIRequest对象
func NewTaobaoTbkScUcrowdDeleteRequest() *TaobaoTbkScUcrowdDeleteAPIRequest {
	return &TaobaoTbkScUcrowdDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTbkScUcrowdDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.sc.ucrowd.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTbkScUcrowdDeleteAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetUcrowdId is UcrowdId Setter
// 人群标签id
func (r *TaobaoTbkScUcrowdDeleteAPIRequest) SetUcrowdId(_ucrowdId int64) error {
	r._ucrowdId = _ucrowdId
	r.Set("ucrowd_id", _ucrowdId)
	return nil
}

// GetUcrowdId UcrowdId Getter
func (r TaobaoTbkScUcrowdDeleteAPIRequest) GetUcrowdId() int64 {
	return r._ucrowdId
}
