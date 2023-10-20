package feedflow

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFeedflowItemAdgroupCreativeAddBindAPIRequest) Reset() {
	r._creativeBindList = r._creativeBindList[:0]
	r._adgroupId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemAdgroupCreativeAddBindAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.adgroup.creative.add.bind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemAdgroupCreativeAddBindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFeedflowItemAdgroupCreativeAddBindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCreativeBindList is CreativeBindList Setter
// 新增绑定的创意，一次最多2个
func (r *TaobaoFeedflowItemAdgroupCreativeAddBindAPIRequest) SetCreativeBindList(_creativeBindList []CreativeBindDto) error {
	r._creativeBindList = _creativeBindList
	r.Set("creative_bind_list", _creativeBindList)
	return nil
}

// GetCreativeBindList CreativeBindList Getter
func (r TaobaoFeedflowItemAdgroupCreativeAddBindAPIRequest) GetCreativeBindList() []CreativeBindDto {
	return r._creativeBindList
}

// SetAdgroupId is AdgroupId Setter
// 单元id
func (r *TaobaoFeedflowItemAdgroupCreativeAddBindAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaoFeedflowItemAdgroupCreativeAddBindAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}

var poolTaobaoFeedflowItemAdgroupCreativeAddBindAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFeedflowItemAdgroupCreativeAddBindRequest()
	},
}

// GetTaobaoFeedflowItemAdgroupCreativeAddBindRequest 从 sync.Pool 获取 TaobaoFeedflowItemAdgroupCreativeAddBindAPIRequest
func GetTaobaoFeedflowItemAdgroupCreativeAddBindAPIRequest() *TaobaoFeedflowItemAdgroupCreativeAddBindAPIRequest {
	return poolTaobaoFeedflowItemAdgroupCreativeAddBindAPIRequest.Get().(*TaobaoFeedflowItemAdgroupCreativeAddBindAPIRequest)
}

// ReleaseTaobaoFeedflowItemAdgroupCreativeAddBindAPIRequest 将 TaobaoFeedflowItemAdgroupCreativeAddBindAPIRequest 放入 sync.Pool
func ReleaseTaobaoFeedflowItemAdgroupCreativeAddBindAPIRequest(v *TaobaoFeedflowItemAdgroupCreativeAddBindAPIRequest) {
	v.Reset()
	poolTaobaoFeedflowItemAdgroupCreativeAddBindAPIRequest.Put(v)
}
