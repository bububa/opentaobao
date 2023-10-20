package game

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoappleolduserchargenotifyAPIRequest 老用户激活并兑换通知接口 API请求
// taobao.apple.olduser.charge.notify
//
// 老用户激活并兑换通知接口
type TaobaoappleolduserchargenotifyAPIRequest struct {
	model.Params
	// 结果code
	_resultCode string
	// 结果信息说明
	_resultMsg string
	// 业务参数
	_mainData *AppleTopOldSignNotifyDo
}

// NewTaobaoappleolduserchargenotifyRequest 初始化TaobaoappleolduserchargenotifyAPIRequest对象
func NewTaobaoappleolduserchargenotifyRequest() *TaobaoappleolduserchargenotifyAPIRequest {
	return &TaobaoappleolduserchargenotifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoappleolduserchargenotifyAPIRequest) GetApiMethodName() string {
	return "taobao.apple.olduser.charge.notify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoappleolduserchargenotifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoappleolduserchargenotifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetResultCode is ResultCode Setter
// 结果code
func (r *TaobaoappleolduserchargenotifyAPIRequest) SetResultCode(_resultCode string) error {
	r._resultCode = _resultCode
	r.Set("result_code", _resultCode)
	return nil
}

// GetResultCode ResultCode Getter
func (r TaobaoappleolduserchargenotifyAPIRequest) GetResultCode() string {
	return r._resultCode
}

// SetResultMsg is ResultMsg Setter
// 结果信息说明
func (r *TaobaoappleolduserchargenotifyAPIRequest) SetResultMsg(_resultMsg string) error {
	r._resultMsg = _resultMsg
	r.Set("result_msg", _resultMsg)
	return nil
}

// GetResultMsg ResultMsg Getter
func (r TaobaoappleolduserchargenotifyAPIRequest) GetResultMsg() string {
	return r._resultMsg
}

// SetMainData is MainData Setter
// 业务参数
func (r *TaobaoappleolduserchargenotifyAPIRequest) SetMainData(_mainData *AppleTopOldSignNotifyDo) error {
	r._mainData = _mainData
	r.Set("main_data", _mainData)
	return nil
}

// GetMainData MainData Getter
func (r TaobaoappleolduserchargenotifyAPIRequest) GetMainData() *AppleTopOldSignNotifyDo {
	return r._mainData
}
