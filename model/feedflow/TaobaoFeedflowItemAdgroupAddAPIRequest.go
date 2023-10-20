package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofeedflowitemadgroupaddAPIRequest 信息流增加单元 API请求
// taobao.feedflow.item.adgroup.add
//
// 信息流增加单元
type TaobaofeedflowitemadgroupaddAPIRequest struct {
	model.Params
	// 单元信息
	_adgroup *AdgroupDto
}

// NewTaobaofeedflowitemadgroupaddRequest 初始化TaobaofeedflowitemadgroupaddAPIRequest对象
func NewTaobaofeedflowitemadgroupaddRequest() *TaobaofeedflowitemadgroupaddAPIRequest {
	return &TaobaofeedflowitemadgroupaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofeedflowitemadgroupaddAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.adgroup.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofeedflowitemadgroupaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofeedflowitemadgroupaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAdgroup is Adgroup Setter
// 单元信息
func (r *TaobaofeedflowitemadgroupaddAPIRequest) SetAdgroup(_adgroup *AdgroupDto) error {
	r._adgroup = _adgroup
	r.Set("adgroup", _adgroup)
	return nil
}

// GetAdgroup Adgroup Getter
func (r TaobaofeedflowitemadgroupaddAPIRequest) GetAdgroup() *AdgroupDto {
	return r._adgroup
}
