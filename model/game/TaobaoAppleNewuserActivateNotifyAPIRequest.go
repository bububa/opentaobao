package game

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAppleNewuserActivateNotifyAPIRequest 新用户激活通知接口 API请求
// taobao.apple.newuser.activate.notify
//
// 资和信主动通知激活结果
type TaobaoAppleNewuserActivateNotifyAPIRequest struct {
	model.Params
	// 结果对应值，00位成功，其他为失败
	_resultCode string
	// 处理结果中文描述
	_resultMsg string
	// 主业务参数
	_mainData *AppleTopActivateNotifyDo
}

// NewTaobaoAppleNewuserActivateNotifyRequest 初始化TaobaoAppleNewuserActivateNotifyAPIRequest对象
func NewTaobaoAppleNewuserActivateNotifyRequest() *TaobaoAppleNewuserActivateNotifyAPIRequest {
	return &TaobaoAppleNewuserActivateNotifyAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAppleNewuserActivateNotifyAPIRequest) Reset() {
	r._resultCode = ""
	r._resultMsg = ""
	r._mainData = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAppleNewuserActivateNotifyAPIRequest) GetApiMethodName() string {
	return "taobao.apple.newuser.activate.notify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAppleNewuserActivateNotifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAppleNewuserActivateNotifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetResultCode is ResultCode Setter
// 结果对应值，00位成功，其他为失败
func (r *TaobaoAppleNewuserActivateNotifyAPIRequest) SetResultCode(_resultCode string) error {
	r._resultCode = _resultCode
	r.Set("result_code", _resultCode)
	return nil
}

// GetResultCode ResultCode Getter
func (r TaobaoAppleNewuserActivateNotifyAPIRequest) GetResultCode() string {
	return r._resultCode
}

// SetResultMsg is ResultMsg Setter
// 处理结果中文描述
func (r *TaobaoAppleNewuserActivateNotifyAPIRequest) SetResultMsg(_resultMsg string) error {
	r._resultMsg = _resultMsg
	r.Set("result_msg", _resultMsg)
	return nil
}

// GetResultMsg ResultMsg Getter
func (r TaobaoAppleNewuserActivateNotifyAPIRequest) GetResultMsg() string {
	return r._resultMsg
}

// SetMainData is MainData Setter
// 主业务参数
func (r *TaobaoAppleNewuserActivateNotifyAPIRequest) SetMainData(_mainData *AppleTopActivateNotifyDo) error {
	r._mainData = _mainData
	r.Set("main_data", _mainData)
	return nil
}

// GetMainData MainData Getter
func (r TaobaoAppleNewuserActivateNotifyAPIRequest) GetMainData() *AppleTopActivateNotifyDo {
	return r._mainData
}

var poolTaobaoAppleNewuserActivateNotifyAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAppleNewuserActivateNotifyRequest()
	},
}

// GetTaobaoAppleNewuserActivateNotifyRequest 从 sync.Pool 获取 TaobaoAppleNewuserActivateNotifyAPIRequest
func GetTaobaoAppleNewuserActivateNotifyAPIRequest() *TaobaoAppleNewuserActivateNotifyAPIRequest {
	return poolTaobaoAppleNewuserActivateNotifyAPIRequest.Get().(*TaobaoAppleNewuserActivateNotifyAPIRequest)
}

// ReleaseTaobaoAppleNewuserActivateNotifyAPIRequest 将 TaobaoAppleNewuserActivateNotifyAPIRequest 放入 sync.Pool
func ReleaseTaobaoAppleNewuserActivateNotifyAPIRequest(v *TaobaoAppleNewuserActivateNotifyAPIRequest) {
	v.Reset()
	poolTaobaoAppleNewuserActivateNotifyAPIRequest.Put(v)
}
