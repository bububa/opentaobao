package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaocbossworkplatformbiztypequerybyidAPIRequest 菜鸟工单平台根据业务类型id查询业务类型详细信息 API请求
// cainiao.cboss.workplatform.biztype.querybyid
//
// 菜鸟工单平台根据业务类型id查询业务类型详细信息。 目前调用者ISV
type CainiaocbossworkplatformbiztypequerybyidAPIRequest struct {
	model.Params
	// 业务类型id
	_bizTypeId string
}

// NewCainiaocbossworkplatformbiztypequerybyidRequest 初始化CainiaocbossworkplatformbiztypequerybyidAPIRequest对象
func NewCainiaocbossworkplatformbiztypequerybyidRequest() *CainiaocbossworkplatformbiztypequerybyidAPIRequest {
	return &CainiaocbossworkplatformbiztypequerybyidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaocbossworkplatformbiztypequerybyidAPIRequest) GetApiMethodName() string {
	return "cainiao.cboss.workplatform.biztype.querybyid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaocbossworkplatformbiztypequerybyidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaocbossworkplatformbiztypequerybyidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBizTypeId is BizTypeId Setter
// 业务类型id
func (r *CainiaocbossworkplatformbiztypequerybyidAPIRequest) SetBizTypeId(_bizTypeId string) error {
	r._bizTypeId = _bizTypeId
	r.Set("biz_type_id", _bizTypeId)
	return nil
}

// GetBizTypeId BizTypeId Getter
func (r CainiaocbossworkplatformbiztypequerybyidAPIRequest) GetBizTypeId() string {
	return r._bizTypeId
}
