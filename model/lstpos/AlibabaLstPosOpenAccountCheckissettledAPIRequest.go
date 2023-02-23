package lstpos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstPosOpenAccountCheckissettledAPIRequest 校验当前用户是否入驻了零售通门店接口 API请求
// alibaba.lst.pos.open.account.checkissettled
//
// 校验当前用户是否入驻了零售通门店接口
type AlibabaLstPosOpenAccountCheckissettledAPIRequest struct {
	model.Params
	// 当前登录主账号userId
	_userId int64
}

// NewAlibabaLstPosOpenAccountCheckissettledRequest 初始化AlibabaLstPosOpenAccountCheckissettledAPIRequest对象
func NewAlibabaLstPosOpenAccountCheckissettledRequest() *AlibabaLstPosOpenAccountCheckissettledAPIRequest {
	return &AlibabaLstPosOpenAccountCheckissettledAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLstPosOpenAccountCheckissettledAPIRequest) GetApiMethodName() string {
	return "alibaba.lst.pos.open.account.checkissettled"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLstPosOpenAccountCheckissettledAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaLstPosOpenAccountCheckissettledAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserId is UserId Setter
// 当前登录主账号userId
func (r *AlibabaLstPosOpenAccountCheckissettledAPIRequest) SetUserId(_userId int64) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabaLstPosOpenAccountCheckissettledAPIRequest) GetUserId() int64 {
	return r._userId
}
