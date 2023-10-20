package tblogistics

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaowlbstoresbaseinfogetAPIRequest 商家按照仓的类型获取仓库接口 API请求
// taobao.wlb.stores.baseinfo.get
//
// 通过USERID和仓库类型，获取商家自有仓库或菜鸟仓库或全部仓库
type TaobaowlbstoresbaseinfogetAPIRequest struct {
	model.Params
	// 0.商家仓库.1.菜鸟仓库.2全部被划分的仓库
	_type int64
}

// NewTaobaowlbstoresbaseinfogetRequest 初始化TaobaowlbstoresbaseinfogetAPIRequest对象
func NewTaobaowlbstoresbaseinfogetRequest() *TaobaowlbstoresbaseinfogetAPIRequest {
	return &TaobaowlbstoresbaseinfogetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaowlbstoresbaseinfogetAPIRequest) GetApiMethodName() string {
	return "taobao.wlb.stores.baseinfo.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaowlbstoresbaseinfogetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaowlbstoresbaseinfogetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetType is Type Setter
// 0.商家仓库.1.菜鸟仓库.2全部被划分的仓库
func (r *TaobaowlbstoresbaseinfogetAPIRequest) SetType(_type int64) error {
	r._type = _type
	r.Set("type", _type)
	return nil
}

// GetType Type Getter
func (r TaobaowlbstoresbaseinfogetAPIRequest) GetType() int64 {
	return r._type
}
