package inventory

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLocationRelationEditAPIRequest 地点关联关系增量编辑 API请求
// taobao.location.relation.edit
//
// 地点关联关系增量编辑
type TaobaoLocationRelationEditAPIRequest struct {
	model.Params
	// 关系对象列表
	_locationRelationList []LocationRelationDto
}

// NewTaobaoLocationRelationEditRequest 初始化TaobaoLocationRelationEditAPIRequest对象
func NewTaobaoLocationRelationEditRequest() *TaobaoLocationRelationEditAPIRequest {
	return &TaobaoLocationRelationEditAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoLocationRelationEditAPIRequest) Reset() {
	r._locationRelationList = r._locationRelationList[:0]
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLocationRelationEditAPIRequest) GetApiMethodName() string {
	return "taobao.location.relation.edit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLocationRelationEditAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoLocationRelationEditAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLocationRelationList is LocationRelationList Setter
// 关系对象列表
func (r *TaobaoLocationRelationEditAPIRequest) SetLocationRelationList(_locationRelationList []LocationRelationDto) error {
	r._locationRelationList = _locationRelationList
	r.Set("location_relation_list", _locationRelationList)
	return nil
}

// GetLocationRelationList LocationRelationList Getter
func (r TaobaoLocationRelationEditAPIRequest) GetLocationRelationList() []LocationRelationDto {
	return r._locationRelationList
}

var poolTaobaoLocationRelationEditAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoLocationRelationEditRequest()
	},
}

// GetTaobaoLocationRelationEditRequest 从 sync.Pool 获取 TaobaoLocationRelationEditAPIRequest
func GetTaobaoLocationRelationEditAPIRequest() *TaobaoLocationRelationEditAPIRequest {
	return poolTaobaoLocationRelationEditAPIRequest.Get().(*TaobaoLocationRelationEditAPIRequest)
}

// ReleaseTaobaoLocationRelationEditAPIRequest 将 TaobaoLocationRelationEditAPIRequest 放入 sync.Pool
func ReleaseTaobaoLocationRelationEditAPIRequest(v *TaobaoLocationRelationEditAPIRequest) {
	v.Reset()
	poolTaobaoLocationRelationEditAPIRequest.Put(v)
}
