package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemAdgroupAdzoneBindAPIRequest 信息流单元内绑定资源位 API请求
// taobao.feedflow.item.adgroup.adzone.bind
//
// 信息流单元内绑定资源位
type TaobaoFeedflowItemAdgroupAdzoneBindAPIRequest struct {
	model.Params
	// 新增的绑定资源位
	_bindAdzoneList []AdzoneBindDto
	// 单元id
	_adgroupId int64
}

// NewTaobaoFeedflowItemAdgroupAdzoneBindRequest 初始化TaobaoFeedflowItemAdgroupAdzoneBindAPIRequest对象
func NewTaobaoFeedflowItemAdgroupAdzoneBindRequest() *TaobaoFeedflowItemAdgroupAdzoneBindAPIRequest {
	return &TaobaoFeedflowItemAdgroupAdzoneBindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemAdgroupAdzoneBindAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.adgroup.adzone.bind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemAdgroupAdzoneBindAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is BindAdzoneList Setter
// 新增的绑定资源位
func (r *TaobaoFeedflowItemAdgroupAdzoneBindAPIRequest) SetBindAdzoneList(_bindAdzoneList []AdzoneBindDto) error {
	r._bindAdzoneList = _bindAdzoneList
	r.Set("bind_adzone_list", _bindAdzoneList)
	return nil
}

// Get BindAdzoneList Getter
func (r TaobaoFeedflowItemAdgroupAdzoneBindAPIRequest) GetBindAdzoneList() []AdzoneBindDto {
	return r._bindAdzoneList
}

// Set is AdgroupId Setter
// 单元id
func (r *TaobaoFeedflowItemAdgroupAdzoneBindAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// Get AdgroupId Getter
func (r TaobaoFeedflowItemAdgroupAdzoneBindAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}
