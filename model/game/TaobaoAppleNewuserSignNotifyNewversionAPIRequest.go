package game

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAppleNewuserSignNotifyNewversionAPIRequest) GetApiMethodName() string {
	return "taobao.apple.newuser.sign.notify.newversion"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAppleNewuserSignNotifyNewversionAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
