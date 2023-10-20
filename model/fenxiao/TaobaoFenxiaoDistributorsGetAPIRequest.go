package fenxiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofenxiaodistributorsgetAPIRequest 获取分销商信息 API请求
// taobao.fenxiao.distributors.get
//
// 查询和当前登录供应商有合作关系的分销商的信息
type TaobaofenxiaodistributorsgetAPIRequest struct {
	model.Params
	// 分销商用户名列表。多个之间以“,”分隔;最多支持50个分销商用户名。
	_nicks string
}

// NewTaobaofenxiaodistributorsgetRequest 初始化TaobaofenxiaodistributorsgetAPIRequest对象
func NewTaobaofenxiaodistributorsgetRequest() *TaobaofenxiaodistributorsgetAPIRequest {
	return &TaobaofenxiaodistributorsgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofenxiaodistributorsgetAPIRequest) GetApiMethodName() string {
	return "taobao.fenxiao.distributors.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofenxiaodistributorsgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofenxiaodistributorsgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetNicks is Nicks Setter
// 分销商用户名列表。多个之间以“,”分隔;最多支持50个分销商用户名。
func (r *TaobaofenxiaodistributorsgetAPIRequest) SetNicks(_nicks string) error {
	r._nicks = _nicks
	r.Set("nicks", _nicks)
	return nil
}

// GetNicks Nicks Getter
func (r TaobaofenxiaodistributorsgetAPIRequest) GetNicks() string {
	return r._nicks
}
