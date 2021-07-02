package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemAdgroupCreativeAddBindAPIRequest 信息流新增并且绑定创意 API请求
// taobao.feedflow.item.adgroup.creative.add.bind
//
// 信息流新增并且绑定创意
type TaobaoFeedflowItemAdgroupCreativeAddBindAPIRequest struct {
	model.Params
	// 新增绑定的创意，一次最多2个
	_creativeBindList []CreativeBindDto
	// 单元id
	_adgroupId int64
}

// NewTaobaoFeedflowItemAdgroupCreativeAddBindRequest 初始化TaobaoFeedflowItemAdgroupCreativeAddBindAPIRequest对象
func NewTaobaoFeedflowItemAdgroupCreativeAddBindRequest() *TaobaoFeedflowItemAdgroupCreativeAddBindAPIRequest {
	return &TaobaoFeedflowItemAdgroupCreativeAddBindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemAdgroupCreativeAddBindAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.adgroup.creative.add.bind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemAdgroupCreativeAddBindAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is CreativeBindList Setter
// 新增绑定的创意，一次最多2个
func (r *TaobaoFeedflowItemAdgroupCreativeAddBindAPIRequest) SetCreativeBindList(_creativeBindList []CreativeBindDto) error {
	r._creativeBindList = _creativeBindList
	r.Set("creative_bind_list", _creativeBindList)
	return nil
}

// Get CreativeBindList Getter
func (r TaobaoFeedflowItemAdgroupCreativeAddBindAPIRequest) GetCreativeBindList() []CreativeBindDto {
	return r._creativeBindList
}

// Set is AdgroupId Setter
// 单元id
func (r *TaobaoFeedflowItemAdgroupCreativeAddBindAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// Get AdgroupId Getter
func (r TaobaoFeedflowItemAdgroupCreativeAddBindAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}
