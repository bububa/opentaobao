package alihealthcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthPregnancyBaseinfoGetAPIRequest 拉取备孕初始化信息 API请求
// alibaba.alihealth.pregnancy.baseinfo.get
//
// 拉取备孕初始化信息
type AlibabaAlihealthPregnancyBaseinfoGetAPIRequest struct {
	model.Params
	// 用户id
	_userId int64
}

// NewAlibabaAlihealthPregnancyBaseinfoGetRequest 初始化AlibabaAlihealthPregnancyBaseinfoGetAPIRequest对象
func NewAlibabaAlihealthPregnancyBaseinfoGetRequest() *AlibabaAlihealthPregnancyBaseinfoGetAPIRequest {
	return &AlibabaAlihealthPregnancyBaseinfoGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthPregnancyBaseinfoGetAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.pregnancy.baseinfo.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthPregnancyBaseinfoGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthPregnancyBaseinfoGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserId is UserId Setter
// 用户id
func (r *AlibabaAlihealthPregnancyBaseinfoGetAPIRequest) SetUserId(_userId int64) error {
	r._userId = _userId
	r.Set("user_id", _userId)
	return nil
}

// GetUserId UserId Getter
func (r AlibabaAlihealthPregnancyBaseinfoGetAPIRequest) GetUserId() int64 {
	return r._userId
}
