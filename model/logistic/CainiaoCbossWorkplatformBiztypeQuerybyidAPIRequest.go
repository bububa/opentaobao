package logistic

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaoCbossWorkplatformBiztypeQuerybyidAPIRequest 菜鸟工单平台根据业务类型id查询业务类型详细信息 API请求
// cainiao.cboss.workplatform.biztype.querybyid
//
// 菜鸟工单平台根据业务类型id查询业务类型详细信息。 目前调用者ISV
type CainiaoCbossWorkplatformBiztypeQuerybyidAPIRequest struct {
	model.Params
	// 业务类型id
	_bizTypeId string
}

// NewCainiaoCbossWorkplatformBiztypeQuerybyidRequest 初始化CainiaoCbossWorkplatformBiztypeQuerybyidAPIRequest对象
func NewCainiaoCbossWorkplatformBiztypeQuerybyidRequest() *CainiaoCbossWorkplatformBiztypeQuerybyidAPIRequest {
	return &CainiaoCbossWorkplatformBiztypeQuerybyidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoCbossWorkplatformBiztypeQuerybyidAPIRequest) GetApiMethodName() string {
	return "cainiao.cboss.workplatform.biztype.querybyid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoCbossWorkplatformBiztypeQuerybyidAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetBizTypeId is BizTypeId Setter
// 业务类型id
func (r *CainiaoCbossWorkplatformBiztypeQuerybyidAPIRequest) SetBizTypeId(_bizTypeId string) error {
	r._bizTypeId = _bizTypeId
	r.Set("biz_type_id", _bizTypeId)
	return nil
}

// GetBizTypeId BizTypeId Getter
func (r CainiaoCbossWorkplatformBiztypeQuerybyidAPIRequest) GetBizTypeId() string {
	return r._bizTypeId
}
