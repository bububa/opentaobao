package game

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoapplenewusersignnotifynewversionAPIRequest 新用户签约结果通知接口v2 API请求
// taobao.apple.newuser.sign.notify.newversion
//
// 资和信主动通知签约结果
type TaobaoapplenewusersignnotifynewversionAPIRequest struct {
	model.Params
	// 结果code
	_resultCode string
	// 结果信息说明
	_resultMsg string
	// 业务参数
	_mainData *AppleTopNewSignNotifyDo
}

// NewTaobaoapplenewusersignnotifynewversionRequest 初始化TaobaoapplenewusersignnotifynewversionAPIRequest对象
func NewTaobaoapplenewusersignnotifynewversionRequest() *TaobaoapplenewusersignnotifynewversionAPIRequest {
	return &TaobaoapplenewusersignnotifynewversionAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoapplenewusersignnotifynewversionAPIRequest) GetApiMethodName() string {
	return "taobao.apple.newuser.sign.notify.newversion"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoapplenewusersignnotifynewversionAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoapplenewusersignnotifynewversionAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetResultCode is ResultCode Setter
// 结果code
func (r *TaobaoapplenewusersignnotifynewversionAPIRequest) SetResultCode(_resultCode string) error {
	r._resultCode = _resultCode
	r.Set("result_code", _resultCode)
	return nil
}

// GetResultCode ResultCode Getter
func (r TaobaoapplenewusersignnotifynewversionAPIRequest) GetResultCode() string {
	return r._resultCode
}

// SetResultMsg is ResultMsg Setter
// 结果信息说明
func (r *TaobaoapplenewusersignnotifynewversionAPIRequest) SetResultMsg(_resultMsg string) error {
	r._resultMsg = _resultMsg
	r.Set("result_msg", _resultMsg)
	return nil
}

// GetResultMsg ResultMsg Getter
func (r TaobaoapplenewusersignnotifynewversionAPIRequest) GetResultMsg() string {
	return r._resultMsg
}

// SetMainData is MainData Setter
// 业务参数
func (r *TaobaoapplenewusersignnotifynewversionAPIRequest) SetMainData(_mainData *AppleTopNewSignNotifyDo) error {
	r._mainData = _mainData
	r.Set("main_data", _mainData)
	return nil
}

// GetMainData MainData Getter
func (r TaobaoapplenewusersignnotifynewversionAPIRequest) GetMainData() *AppleTopNewSignNotifyDo {
	return r._mainData
}
