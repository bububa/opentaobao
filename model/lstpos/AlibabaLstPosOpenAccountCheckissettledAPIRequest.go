package lstpos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalstposopenaccountcheckissettledAPIRequest 校验当前用户是否入驻了零售通门店接口 API请求
// alibaba.lst.pos.open.account.checkissettled
//
// 校验当前用户是否入驻了零售通门店接口
type AlibabalstposopenaccountcheckissettledAPIRequest struct {
	model.Params
	// 当前登录主账号userId
	_userId int64
}

// NewAlibabalstposopenaccountcheckissettledRequest 初始化AlibabalstposopenaccountcheckissettledAPIRequest对象
func NewAlibabalstposopenaccountcheckissettledRequest() *AlibabalstposopenaccountcheckissettledAPIRequest {
	return &AlibabalstposopenaccountcheckissettledAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalstposopenaccountcheckissettledAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.pos.open.account.checkissettled"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalstposopenaccountcheckissettledAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalstposopenaccountcheckissettledAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserId is UserId Setter
// 当前登录主账号userId
func (r *AlibabalstposopenaccountcheckissettledAPIRequest) SetUserId(_userId int64) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabalstposopenaccountcheckissettledAPIRequest) GetUserId() int64 {
	return r._userId
}
