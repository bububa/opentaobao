package feedflow

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoFeedflowItemAdgroupAdzoneBindAPIRequest) Reset() {
	r._bindAdzoneList = r._bindAdzoneList[:0]
	r._adgroupId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoFeedflowItemAdgroupAdzoneBindAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.adgroup.adzone.bind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoFeedflowItemAdgroupAdzoneBindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoFeedflowItemAdgroupAdzoneBindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBindAdzoneList is BindAdzoneList Setter
// 新增的绑定资源位
func (r *TaobaoFeedflowItemAdgroupAdzoneBindAPIRequest) SetBindAdzoneList(_bindAdzoneList []AdzoneBindDto) error {
	r._bindAdzoneList = _bindAdzoneList
	r.Set("bind_adzone_list", _bindAdzoneList)
	return nil
}

// GetBindAdzoneList BindAdzoneList Getter
func (r TaobaoFeedflowItemAdgroupAdzoneBindAPIRequest) GetBindAdzoneList() []AdzoneBindDto {
	return r._bindAdzoneList
}

// SetAdgroupId is AdgroupId Setter
// 单元id
func (r *TaobaoFeedflowItemAdgroupAdzoneBindAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaoFeedflowItemAdgroupAdzoneBindAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}

var poolTaobaoFeedflowItemAdgroupAdzoneBindAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoFeedflowItemAdgroupAdzoneBindRequest()
	},
}

// GetTaobaoFeedflowItemAdgroupAdzoneBindRequest 从 sync.Pool 获取 TaobaoFeedflowItemAdgroupAdzoneBindAPIRequest
func GetTaobaoFeedflowItemAdgroupAdzoneBindAPIRequest() *TaobaoFeedflowItemAdgroupAdzoneBindAPIRequest {
	return poolTaobaoFeedflowItemAdgroupAdzoneBindAPIRequest.Get().(*TaobaoFeedflowItemAdgroupAdzoneBindAPIRequest)
}

// ReleaseTaobaoFeedflowItemAdgroupAdzoneBindAPIRequest 将 TaobaoFeedflowItemAdgroupAdzoneBindAPIRequest 放入 sync.Pool
func ReleaseTaobaoFeedflowItemAdgroupAdzoneBindAPIRequest(v *TaobaoFeedflowItemAdgroupAdzoneBindAPIRequest) {
	v.Reset()
	poolTaobaoFeedflowItemAdgroupAdzoneBindAPIRequest.Put(v)
}
