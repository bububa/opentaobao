package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofeedflowitemadgroupadzonebindAPIRequest 信息流单元内绑定资源位 API请求
// taobao.feedflow.item.adgroup.adzone.bind
//
// 信息流单元内绑定资源位
type TaobaofeedflowitemadgroupadzonebindAPIRequest struct {
	model.Params
	// 新增的绑定资源位
	_bindAdzoneList []AdzoneBindDto
	// 单元id
	_adgroupId int64
}

// NewTaobaofeedflowitemadgroupadzonebindRequest 初始化TaobaofeedflowitemadgroupadzonebindAPIRequest对象
func NewTaobaofeedflowitemadgroupadzonebindRequest() *TaobaofeedflowitemadgroupadzonebindAPIRequest {
	return &TaobaofeedflowitemadgroupadzonebindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofeedflowitemadgroupadzonebindAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.adgroup.adzone.bind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofeedflowitemadgroupadzonebindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofeedflowitemadgroupadzonebindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBindAdzoneList is BindAdzoneList Setter
// 新增的绑定资源位
func (r *TaobaofeedflowitemadgroupadzonebindAPIRequest) SetBindAdzoneList(_bindAdzoneList []AdzoneBindDto) error {
	r._bindAdzoneList = _bindAdzoneList
	r.Set("bind_adzone_list", _bindAdzoneList)
	return nil
}

// GetBindAdzoneList BindAdzoneList Getter
func (r TaobaofeedflowitemadgroupadzonebindAPIRequest) GetBindAdzoneList() []AdzoneBindDto {
	return r._bindAdzoneList
}

// SetAdgroupId is AdgroupId Setter
// 单元id
func (r *TaobaofeedflowitemadgroupadzonebindAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaofeedflowitemadgroupadzonebindAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}
