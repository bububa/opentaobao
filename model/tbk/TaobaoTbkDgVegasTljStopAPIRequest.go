package tbk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTbkDgVegasTljStopAPIRequest 淘宝客-推广者-淘礼金暂停发放 API请求
// taobao.tbk.dg.vegas.tlj.stop
//
// 淘宝客推广者可对已经创建的淘礼金暂停发放
type TaobaoTbkDgVegasTljStopAPIRequest struct {
	model.Params
	// 创建淘礼金时返回的rightsId
	_rightsId string
	// adzoneId
	_adzoneId int64
}

// NewTaobaoTbkDgVegasTljStopRequest 初始化TaobaoTbkDgVegasTljStopAPIRequest对象
func NewTaobaoTbkDgVegasTljStopRequest() *TaobaoTbkDgVegasTljStopAPIRequest {
	return &TaobaoTbkDgVegasTljStopAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTbkDgVegasTljStopAPIRequest) Reset() {
	r._rightsId = ""
	r._adzoneId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTbkDgVegasTljStopAPIRequest) GetApiMethodName() string {
	return "taobao.tbk.dg.vegas.tlj.stop"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTbkDgVegasTljStopAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTbkDgVegasTljStopAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRightsId is RightsId Setter
// 创建淘礼金时返回的rightsId
func (r *TaobaoTbkDgVegasTljStopAPIRequest) SetRightsId(_rightsId string) error {
	r._rightsId = _rightsId
	r.Set("rights_id", _rightsId)
	return nil
}

// GetRightsId RightsId Getter
func (r TaobaoTbkDgVegasTljStopAPIRequest) GetRightsId() string {
	return r._rightsId
}

// SetAdzoneId is AdzoneId Setter
// adzoneId
func (r *TaobaoTbkDgVegasTljStopAPIRequest) SetAdzoneId(_adzoneId int64) error {
	r._adzoneId = _adzoneId
	r.Set("adzone_id", _adzoneId)
	return nil
}

// GetAdzoneId AdzoneId Getter
func (r TaobaoTbkDgVegasTljStopAPIRequest) GetAdzoneId() int64 {
	return r._adzoneId
}

var poolTaobaoTbkDgVegasTljStopAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTbkDgVegasTljStopRequest()
	},
}

// GetTaobaoTbkDgVegasTljStopRequest 从 sync.Pool 获取 TaobaoTbkDgVegasTljStopAPIRequest
func GetTaobaoTbkDgVegasTljStopAPIRequest() *TaobaoTbkDgVegasTljStopAPIRequest {
	return poolTaobaoTbkDgVegasTljStopAPIRequest.Get().(*TaobaoTbkDgVegasTljStopAPIRequest)
}

// ReleaseTaobaoTbkDgVegasTljStopAPIRequest 将 TaobaoTbkDgVegasTljStopAPIRequest 放入 sync.Pool
func ReleaseTaobaoTbkDgVegasTljStopAPIRequest(v *TaobaoTbkDgVegasTljStopAPIRequest) {
	v.Reset()
	poolTaobaoTbkDgVegasTljStopAPIRequest.Put(v)
}
