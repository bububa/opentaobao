package xhotelitem

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoXhotelBnbcommonAddAPIRequest 通用调用top接口 API请求
// taobao.xhotel.bnbcommon.add
//
// 通用调用top接口
type TaobaoXhotelBnbcommonAddAPIRequest struct {
	model.Params
	// 参数
	_param string
	// 业务场景
	_scene string
}

// NewTaobaoXhotelBnbcommonAddRequest 初始化TaobaoXhotelBnbcommonAddAPIRequest对象
func NewTaobaoXhotelBnbcommonAddRequest() *TaobaoXhotelBnbcommonAddAPIRequest {
	return &TaobaoXhotelBnbcommonAddAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoXhotelBnbcommonAddAPIRequest) Reset() {
	r._param = ""
	r._scene = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoXhotelBnbcommonAddAPIRequest) GetApiMethodName() string {
	return "taobao.xhotel.bnbcommon.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoXhotelBnbcommonAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoXhotelBnbcommonAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 参数
func (r *TaobaoXhotelBnbcommonAddAPIRequest) SetParam(_param string) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r TaobaoXhotelBnbcommonAddAPIRequest) GetParam() string {
	return r._param
}

// SetScene is Scene Setter
// 业务场景
func (r *TaobaoXhotelBnbcommonAddAPIRequest) SetScene(_scene string) error {
	r._scene = _scene
	r.Set("scene", _scene)
	return nil
}

// GetScene Scene Getter
func (r TaobaoXhotelBnbcommonAddAPIRequest) GetScene() string {
	return r._scene
}

var poolTaobaoXhotelBnbcommonAddAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoXhotelBnbcommonAddRequest()
	},
}

// GetTaobaoXhotelBnbcommonAddRequest 从 sync.Pool 获取 TaobaoXhotelBnbcommonAddAPIRequest
func GetTaobaoXhotelBnbcommonAddAPIRequest() *TaobaoXhotelBnbcommonAddAPIRequest {
	return poolTaobaoXhotelBnbcommonAddAPIRequest.Get().(*TaobaoXhotelBnbcommonAddAPIRequest)
}

// ReleaseTaobaoXhotelBnbcommonAddAPIRequest 将 TaobaoXhotelBnbcommonAddAPIRequest 放入 sync.Pool
func ReleaseTaobaoXhotelBnbcommonAddAPIRequest(v *TaobaoXhotelBnbcommonAddAPIRequest) {
	v.Reset()
	poolTaobaoXhotelBnbcommonAddAPIRequest.Put(v)
}
