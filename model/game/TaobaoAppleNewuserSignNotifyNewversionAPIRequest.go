package game

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAppleNewuserSignNotifyNewversionAPIRequest 新用户签约结果通知接口v2 API请求
// taobao.apple.newuser.sign.notify.newversion
//
// 资和信主动通知签约结果
type TaobaoAppleNewuserSignNotifyNewversionAPIRequest struct {
	model.Params
	// 结果code
	_resultCode string
	// 结果信息说明
	_resultMsg string
	// 业务参数
	_mainData *AppleTopNewSignNotifyDo
}

// NewTaobaoAppleNewuserSignNotifyNewversionRequest 初始化TaobaoAppleNewuserSignNotifyNewversionAPIRequest对象
func NewTaobaoAppleNewuserSignNotifyNewversionRequest() *TaobaoAppleNewuserSignNotifyNewversionAPIRequest {
	return &TaobaoAppleNewuserSignNotifyNewversionAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAppleNewuserSignNotifyNewversionAPIRequest) Reset() {
	r._resultCode = ""
	r._resultMsg = ""
	r._mainData = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAppleNewuserSignNotifyNewversionAPIRequest) GetApiMethodName() string {
	return "taobao.apple.newuser.sign.notify.newversion"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAppleNewuserSignNotifyNewversionAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAppleNewuserSignNotifyNewversionAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetResultCode is ResultCode Setter
// 结果code
func (r *TaobaoAppleNewuserSignNotifyNewversionAPIRequest) SetResultCode(_resultCode string) error {
	r._resultCode = _resultCode
	r.Set("result_code", _resultCode)
	return nil
}

// GetResultCode ResultCode Getter
func (r TaobaoAppleNewuserSignNotifyNewversionAPIRequest) GetResultCode() string {
	return r._resultCode
}

// SetResultMsg is ResultMsg Setter
// 结果信息说明
func (r *TaobaoAppleNewuserSignNotifyNewversionAPIRequest) SetResultMsg(_resultMsg string) error {
	r._resultMsg = _resultMsg
	r.Set("result_msg", _resultMsg)
	return nil
}

// GetResultMsg ResultMsg Getter
func (r TaobaoAppleNewuserSignNotifyNewversionAPIRequest) GetResultMsg() string {
	return r._resultMsg
}

// SetMainData is MainData Setter
// 业务参数
func (r *TaobaoAppleNewuserSignNotifyNewversionAPIRequest) SetMainData(_mainData *AppleTopNewSignNotifyDo) error {
	r._mainData = _mainData
	r.Set("main_data", _mainData)
	return nil
}

// GetMainData MainData Getter
func (r TaobaoAppleNewuserSignNotifyNewversionAPIRequest) GetMainData() *AppleTopNewSignNotifyDo {
	return r._mainData
}

var poolTaobaoAppleNewuserSignNotifyNewversionAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAppleNewuserSignNotifyNewversionRequest()
	},
}

// GetTaobaoAppleNewuserSignNotifyNewversionRequest 从 sync.Pool 获取 TaobaoAppleNewuserSignNotifyNewversionAPIRequest
func GetTaobaoAppleNewuserSignNotifyNewversionAPIRequest() *TaobaoAppleNewuserSignNotifyNewversionAPIRequest {
	return poolTaobaoAppleNewuserSignNotifyNewversionAPIRequest.Get().(*TaobaoAppleNewuserSignNotifyNewversionAPIRequest)
}

// ReleaseTaobaoAppleNewuserSignNotifyNewversionAPIRequest 将 TaobaoAppleNewuserSignNotifyNewversionAPIRequest 放入 sync.Pool
func ReleaseTaobaoAppleNewuserSignNotifyNewversionAPIRequest(v *TaobaoAppleNewuserSignNotifyNewversionAPIRequest) {
	v.Reset()
	poolTaobaoAppleNewuserSignNotifyNewversionAPIRequest.Put(v)
}
