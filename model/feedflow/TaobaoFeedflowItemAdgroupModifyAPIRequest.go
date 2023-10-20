package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofeedflowitemadgroupmodifyAPIRequest 信息流单元修改 API请求
// taobao.feedflow.item.adgroup.modify
//
// 信息流单元修改
type TaobaofeedflowitemadgroupmodifyAPIRequest struct {
	model.Params
	// 单元信息
	_adgroup *AdgroupDto
}

// NewTaobaofeedflowitemadgroupmodifyRequest 初始化TaobaofeedflowitemadgroupmodifyAPIRequest对象
func NewTaobaofeedflowitemadgroupmodifyRequest() *TaobaofeedflowitemadgroupmodifyAPIRequest {
	return &TaobaofeedflowitemadgroupmodifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofeedflowitemadgroupmodifyAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.adgroup.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofeedflowitemadgroupmodifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofeedflowitemadgroupmodifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAdgroup is Adgroup Setter
// 单元信息
func (r *TaobaofeedflowitemadgroupmodifyAPIRequest) SetAdgroup(_adgroup *AdgroupDto) error {
	r._adgroup = _adgroup
	r.Set("adgroup", _adgroup)
	return nil
}

// GetAdgroup Adgroup Getter
func (r TaobaofeedflowitemadgroupmodifyAPIRequest) GetAdgroup() *AdgroupDto {
	return r._adgroup
}
