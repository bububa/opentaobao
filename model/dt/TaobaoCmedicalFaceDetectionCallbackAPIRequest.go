package dt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCmedicalFaceDetectionCallbackAPIRequest 魔镜测肤结果数据回调接口 API请求
// taobao.cmedical.face.detection.callback
//
// 消费医疗魔镜项目，isv将异步测肤结果数据，回传给行业。
type TaobaoCmedicalFaceDetectionCallbackAPIRequest struct {
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

// NewTaobaoCmedicalFaceDetectionCallbackRequest 初始化TaobaoCmedicalFaceDetectionCallbackAPIRequest对象
func NewTaobaoCmedicalFaceDetectionCallbackRequest() *TaobaoCmedicalFaceDetectionCallbackAPIRequest {
	return &TaobaoCmedicalFaceDetectionCallbackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoCmedicalFaceDetectionCallbackAPIRequest) GetApiMethodName() string {
	return "taobao.cmedical.face.detection.callback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoCmedicalFaceDetectionCallbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoCmedicalFaceDetectionCallbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequestId is RequestId Setter
// 一次测肤接口的唯一请求id
func (r *TaobaoCmedicalFaceDetectionCallbackAPIRequest) SetRequestId(_requestId string) error {
	r._requestId = _requestId
	r.Set("request_id", _requestId)
	return nil
}

// GetRequestId RequestId Getter
func (r TaobaoCmedicalFaceDetectionCallbackAPIRequest) GetRequestId() string {
	return r._requestId
}

// SetIdentity is Identity Setter
// isv身份识别
func (r *TaobaoCmedicalFaceDetectionCallbackAPIRequest) SetIdentity(_identity string) error {
	r._identity = _identity
	r.Set("identity", _identity)
	return nil
}

// GetIdentity Identity Getter
func (r TaobaoCmedicalFaceDetectionCallbackAPIRequest) GetIdentity() string {
	return r._identity
}

// SetScene is Scene Setter
// 场景，直播：live，店铺：shop
func (r *TaobaoCmedicalFaceDetectionCallbackAPIRequest) SetScene(_scene string) error {
	r._scene = _scene
	r.Set("scene", _scene)
	return nil
}

// GetScene Scene Getter
func (r TaobaoCmedicalFaceDetectionCallbackAPIRequest) GetScene() string {
	return r._scene
}

// SetData is Data Setter
// 测肤结果数据
func (r *TaobaoCmedicalFaceDetectionCallbackAPIRequest) SetData(_data string) error {
	r._data = _data
	r.Set("data", _data)
	return nil
}

// GetData Data Getter
func (r TaobaoCmedicalFaceDetectionCallbackAPIRequest) GetData() string {
	return r._data
}

// SetErrorMsg is ErrorMsg Setter
// 错误消息
func (r *TaobaoCmedicalFaceDetectionCallbackAPIRequest) SetErrorMsg(_errorMsg string) error {
	r._errorMsg = _errorMsg
	r.Set("error_msg", _errorMsg)
	return nil
}

// GetErrorMsg ErrorMsg Getter
func (r TaobaoCmedicalFaceDetectionCallbackAPIRequest) GetErrorMsg() string {
	return r._errorMsg
}

// SetSuccess is Success Setter
// true:成功，false：失败
func (r *TaobaoCmedicalFaceDetectionCallbackAPIRequest) SetSuccess(_success bool) error {
	r._success = _success
	r.Set("success", _success)
	return nil
}

// GetSuccess Success Getter
func (r TaobaoCmedicalFaceDetectionCallbackAPIRequest) GetSuccess() bool {
	return r._success
}
