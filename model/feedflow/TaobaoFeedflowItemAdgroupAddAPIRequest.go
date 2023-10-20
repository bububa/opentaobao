package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemAdgroupAddAPIRequest 信息流增加单元 API请求
// taobao.feedflow.item.adgroup.add
//
// 信息流增加单元
type TaobaoFeedflowItemAdgroupAddAPIRequest struct {
	model.Params
	// 单元信息
	_adgroup *AdgroupDto
}

// NewTaobaoFeedflowItemAdgroupAddRequest 初始化TaobaoFeedflowItemAdgroupAddAPIRequest对象
func NewTaobaoFeedflowItemAdgroupAddRequest() *TaobaoFeedflowItemAdgroupAddAPIRequest {
	return &TaobaoFeedflowItemAdgroupAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemAdgroupAddAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.adgroup.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemAdgroupAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFeedflowItemAdgroupAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAdgroup is Adgroup Setter
// 单元信息
func (r *TaobaoFeedflowItemAdgroupAddAPIRequest) SetAdgroup(_adgroup *AdgroupDto) error {
	r._adgroup = _adgroup
	r.Set("adgroup", _adgroup)
	return nil
}

// GetAdgroup Adgroup Getter
func (r TaobaoFeedflowItemAdgroupAddAPIRequest) GetAdgroup() *AdgroupDto {
	return r._adgroup
}
