package itpolicy

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripItFareDeleteAPIRequest 【国际机票自有政策】单条删除 API请求
// taobao.alitrip.it.fare.delete
//
// 自有政策删除接口，可以根据fareId或outId删除，根据outId删除时，如果outId不唯一，返回失败
type TaobaoAlitripItFareDeleteAPIRequest struct {
	model.Params
	// json格式的字符串，扩展属性，预留
	_extendAttributes string
	// 外部id，为新增时请求参数中的外部政策id
	_outId string
	// 运价id，单条新增成功时返回运价id，fareId和outId必填一个，fareId优先
	_fareId int64
}

// NewTaobaoAlitripItFareDeleteRequest 初始化TaobaoAlitripItFareDeleteAPIRequest对象
func NewTaobaoAlitripItFareDeleteRequest() *TaobaoAlitripItFareDeleteAPIRequest {
	return &TaobaoAlitripItFareDeleteAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripItFareDeleteAPIRequest) Reset() {
	r._extendAttributes = ""
	r._outId = ""
	r._fareId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripItFareDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.it.fare.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripItFareDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripItFareDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetExtendAttributes is ExtendAttributes Setter
// json格式的字符串，扩展属性，预留
func (r *TaobaoAlitripItFareDeleteAPIRequest) SetExtendAttributes(_extendAttributes string) error {
	r._extendAttributes = _extendAttributes
	r.Set("extendAttributes", _extendAttributes)
	return nil
}

// GetExtendAttributes ExtendAttributes Getter
func (r TaobaoAlitripItFareDeleteAPIRequest) GetExtendAttributes() string {
	return r._extendAttributes
}

// SetOutId is OutId Setter
// 外部id，为新增时请求参数中的外部政策id
func (r *TaobaoAlitripItFareDeleteAPIRequest) SetOutId(_outId string) error {
	r._outId = _outId
	r.Set("outId", _outId)
	return nil
}

// GetOutId OutId Getter
func (r TaobaoAlitripItFareDeleteAPIRequest) GetOutId() string {
	return r._outId
}

// SetFareId is FareId Setter
// 运价id，单条新增成功时返回运价id，fareId和outId必填一个，fareId优先
func (r *TaobaoAlitripItFareDeleteAPIRequest) SetFareId(_fareId int64) error {
	r._fareId = _fareId
	r.Set("fareId", _fareId)
	return nil
}

// GetFareId FareId Getter
func (r TaobaoAlitripItFareDeleteAPIRequest) GetFareId() int64 {
	return r._fareId
}

var poolTaobaoAlitripItFareDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripItFareDeleteRequest()
	},
}

// GetTaobaoAlitripItFareDeleteRequest 从 sync.Pool 获取 TaobaoAlitripItFareDeleteAPIRequest
func GetTaobaoAlitripItFareDeleteAPIRequest() *TaobaoAlitripItFareDeleteAPIRequest {
	return poolTaobaoAlitripItFareDeleteAPIRequest.Get().(*TaobaoAlitripItFareDeleteAPIRequest)
}

// ReleaseTaobaoAlitripItFareDeleteAPIRequest 将 TaobaoAlitripItFareDeleteAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripItFareDeleteAPIRequest(v *TaobaoAlitripItFareDeleteAPIRequest) {
	v.Reset()
	poolTaobaoAlitripItFareDeleteAPIRequest.Put(v)
}
