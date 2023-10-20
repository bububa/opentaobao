package game

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAppleNewuserSignNotifyAPIRequest 新用户签约通知接口 API请求
// taobao.apple.newuser.sign.notify
//
// 用户付款成功后，资和信主动通知签约结果
type TaobaoAppleNewuserSignNotifyAPIRequest struct {
	model.Params
	// 结果code
	_resultCode string
	// 结果信息说明
	_resultMsg string
	// 业务参数
	_mainData *AppleTopNewSignNotifyDo
}

// NewTaobaoAppleNewuserSignNotifyRequest 初始化TaobaoAppleNewuserSignNotifyAPIRequest对象
func NewTaobaoAppleNewuserSignNotifyRequest() *TaobaoAppleNewuserSignNotifyAPIRequest {
	return &TaobaoAppleNewuserSignNotifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAppleNewuserSignNotifyAPIRequest) GetApiMethodName() string {
	return "taobao.apple.newuser.sign.notify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAppleNewuserSignNotifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAppleNewuserSignNotifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetResultCode is ResultCode Setter
// 结果code
func (r *TaobaoAppleNewuserSignNotifyAPIRequest) SetResultCode(_resultCode string) error {
	r._resultCode = _resultCode
	r.Set("result_code", _resultCode)
	return nil
}

// GetResultCode ResultCode Getter
func (r TaobaoAppleNewuserSignNotifyAPIRequest) GetResultCode() string {
	return r._resultCode
}

// SetResultMsg is ResultMsg Setter
// 结果信息说明
func (r *TaobaoAppleNewuserSignNotifyAPIRequest) SetResultMsg(_resultMsg string) error {
	r._resultMsg = _resultMsg
	r.Set("result_msg", _resultMsg)
	return nil
}

// GetResultMsg ResultMsg Getter
func (r TaobaoAppleNewuserSignNotifyAPIRequest) GetResultMsg() string {
	return r._resultMsg
}

// SetMainData is MainData Setter
// 业务参数
func (r *TaobaoAppleNewuserSignNotifyAPIRequest) SetMainData(_mainData *AppleTopNewSignNotifyDo) error {
	r._mainData = _mainData
	r.Set("main_data", _mainData)
	return nil
}

// GetMainData MainData Getter
func (r TaobaoAppleNewuserSignNotifyAPIRequest) GetMainData() *AppleTopNewSignNotifyDo {
	return r._mainData
}
