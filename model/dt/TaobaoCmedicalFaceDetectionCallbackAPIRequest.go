package dt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaocmedicalfacedetectioncallbackAPIRequest 魔镜测肤结果数据回调接口 API请求
// taobao.cmedical.face.detection.callback
//
// 消费医疗魔镜项目，isv将异步测肤结果数据，回传给行业。
type TaobaocmedicalfacedetectioncallbackAPIRequest struct {
	model.Params
	// 一次测肤接口的唯一请求id
	_requestId string
	// isv身份识别
	_identity string
	// 场景，直播：live，店铺：shop
	_scene string
	// 测肤结果数据
	_data string
	// 错误消息
	_errorMsg string
	// true:成功，false：失败
	_success bool
}

// NewTaobaocmedicalfacedetectioncallbackRequest 初始化TaobaocmedicalfacedetectioncallbackAPIRequest对象
func NewTaobaocmedicalfacedetectioncallbackRequest() *TaobaocmedicalfacedetectioncallbackAPIRequest {
	return &TaobaocmedicalfacedetectioncallbackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaocmedicalfacedetectioncallbackAPIRequest) GetApiMethodName() string {
	return "taobao.cmedical.face.detection.callback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaocmedicalfacedetectioncallbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaocmedicalfacedetectioncallbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequestId is RequestId Setter
// 一次测肤接口的唯一请求id
func (r *TaobaocmedicalfacedetectioncallbackAPIRequest) SetRequestId(_requestId string) error {
	r._requestId = _requestId
	r.Set("request_id", _requestId)
	return nil
}

// GetRequestId RequestId Getter
func (r TaobaocmedicalfacedetectioncallbackAPIRequest) GetRequestId() string {
	return r._requestId
}

// SetIdentity is Identity Setter
// isv身份识别
func (r *TaobaocmedicalfacedetectioncallbackAPIRequest) SetIdentity(_identity string) error {
	r._identity = _identity
	r.Set("identity", _identity)
	return nil
}

// GetIdentity Identity Getter
func (r TaobaocmedicalfacedetectioncallbackAPIRequest) GetIdentity() string {
	return r._identity
}

// SetScene is Scene Setter
// 场景，直播：live，店铺：shop
func (r *TaobaocmedicalfacedetectioncallbackAPIRequest) SetScene(_scene string) error {
	r._scene = _scene
	r.Set("scene", _scene)
	return nil
}

// GetScene Scene Getter
func (r TaobaocmedicalfacedetectioncallbackAPIRequest) GetScene() string {
	return r._scene
}

// SetData is Data Setter
// 测肤结果数据
func (r *TaobaocmedicalfacedetectioncallbackAPIRequest) SetData(_data string) error {
	r._data = _data
	r.Set("data", _data)
	return nil
}

// GetData Data Getter
func (r TaobaocmedicalfacedetectioncallbackAPIRequest) GetData() string {
	return r._data
}

// SetErrorMsg is ErrorMsg Setter
// 错误消息
func (r *TaobaocmedicalfacedetectioncallbackAPIRequest) SetErrorMsg(_errorMsg string) error {
	r._errorMsg = _errorMsg
	r.Set("error_msg", _errorMsg)
	return nil
}

// GetErrorMsg ErrorMsg Getter
func (r TaobaocmedicalfacedetectioncallbackAPIRequest) GetErrorMsg() string {
	return r._errorMsg
}

// SetSuccess is Success Setter
// true:成功，false：失败
func (r *TaobaocmedicalfacedetectioncallbackAPIRequest) SetSuccess(_success bool) error {
	r._success = _success
	r.Set("success", _success)
	return nil
}

// GetSuccess Success Getter
func (r TaobaocmedicalfacedetectioncallbackAPIRequest) GetSuccess() bool {
	return r._success
}
