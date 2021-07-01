package c2m

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSebpIsvUserSignAPIRequest
淘小铺三方签约同步 API请求
taobao.sebp.isv.user.sign

同步淘小铺三方服务签约信息 */
type TaobaoSebpIsvUserSignAPIRequest struct {
	model.Params
	// 淘宝账号
	_userName string
	// 身份证
	_identity string
	// 到期日期
	_endTime string
	// 签约日期
	_startTime string
}

// NewTaobaoSebpIsvUserSignRequest 初始化TaobaoSebpIsvUserSignAPIRequest对象
func NewTaobaoSebpIsvUserSignRequest() *TaobaoSebpIsvUserSignAPIRequest {
	return &TaobaoSebpIsvUserSignAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoSebpIsvUserSignAPIRequest) GetApiMethodName() string {
	return "taobao.sebp.isv.user.sign"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoSebpIsvUserSignAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is UserName Setter
// 淘宝账号
func (r *TaobaoSebpIsvUserSignAPIRequest) SetUserName(_userName string) error {
	r._userName = _userName
	r.Set("user_name", _userName)
	return nil
}

// Get UserName Getter
func (r TaobaoSebpIsvUserSignAPIRequest) GetUserName() string {
	return r._userName
}

// Set is Identity Setter
// 身份证
func (r *TaobaoSebpIsvUserSignAPIRequest) SetIdentity(_identity string) error {
	r._identity = _identity
	r.Set("identity", _identity)
	return nil
}

// Get Identity Getter
func (r TaobaoSebpIsvUserSignAPIRequest) GetIdentity() string {
	return r._identity
}

// Set is EndTime Setter
// 到期日期
func (r *TaobaoSebpIsvUserSignAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// Get EndTime Getter
func (r TaobaoSebpIsvUserSignAPIRequest) GetEndTime() string {
	return r._endTime
}

// Set is StartTime Setter
// 签约日期
func (r *TaobaoSebpIsvUserSignAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// Get StartTime Getter
func (r TaobaoSebpIsvUserSignAPIRequest) GetStartTime() string {
	return r._startTime
}
