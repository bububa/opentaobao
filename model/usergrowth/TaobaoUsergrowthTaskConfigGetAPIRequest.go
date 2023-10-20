package usergrowth

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaousergrowthtaskconfiggetAPIRequest 用户增长营销玩法配置查询 API请求
// taobao.usergrowth.task.config.get
//
// 用户增长营销玩法配置查询
type TaobaousergrowthtaskconfiggetAPIRequest struct {
	model.Params
	// 业务Id,不同场景不同业务id,淘宝同学提供,
	_businessId string
	// 执行命令,不同场景不同业务id,淘宝同学提供
	_command string
	// 扩展参数,可空,Map<String,String>结构
	_extra string
	// 用户加密后的openid
	_openId string
}

// NewTaobaousergrowthtaskconfiggetRequest 初始化TaobaousergrowthtaskconfiggetAPIRequest对象
func NewTaobaousergrowthtaskconfiggetRequest() *TaobaousergrowthtaskconfiggetAPIRequest {
	return &TaobaousergrowthtaskconfiggetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaousergrowthtaskconfiggetAPIRequest) GetApiMethodName() string {
	return "taobao.usergrowth.task.config.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaousergrowthtaskconfiggetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaousergrowthtaskconfiggetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBusinessId is BusinessId Setter
// 业务Id,不同场景不同业务id,淘宝同学提供,
func (r *TaobaousergrowthtaskconfiggetAPIRequest) SetBusinessId(_businessId string) error {
	r._businessId = _businessId
	r.Set("business_id", _businessId)
	return nil
}

// GetBusinessId BusinessId Getter
func (r TaobaousergrowthtaskconfiggetAPIRequest) GetBusinessId() string {
	return r._businessId
}

// SetCommand is Command Setter
// 执行命令,不同场景不同业务id,淘宝同学提供
func (r *TaobaousergrowthtaskconfiggetAPIRequest) SetCommand(_command string) error {
	r._command = _command
	r.Set("command", _command)
	return nil
}

// GetCommand Command Getter
func (r TaobaousergrowthtaskconfiggetAPIRequest) GetCommand() string {
	return r._command
}

// SetExtra is Extra Setter
// 扩展参数,可空,Map&lt;String,String&gt;结构
func (r *TaobaousergrowthtaskconfiggetAPIRequest) SetExtra(_extra string) error {
	r._extra = _extra
	r.Set("extra", _extra)
	return nil
}

// GetExtra Extra Getter
func (r TaobaousergrowthtaskconfiggetAPIRequest) GetExtra() string {
	return r._extra
}

// SetOpenId is OpenId Setter
// 用户加密后的openid
func (r *TaobaousergrowthtaskconfiggetAPIRequest) SetOpenId(_openId string) error {
	r._openId = _openId
	r.Set("open_id", _openId)
	return nil
}

// GetOpenId OpenId Getter
func (r TaobaousergrowthtaskconfiggetAPIRequest) GetOpenId() string {
	return r._openId
}
