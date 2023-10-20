package tmallsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenteranomalyrecoursecloseAPIRequest 服务投诉问题单关单 API请求
// tmall.servicecenter.anomalyrecourse.close
//
// 提供给服务商在投诉单完结时调用，关闭投诉问题工单。
type TmallservicecenteranomalyrecoursecloseAPIRequest struct {
	model.Params
	// 完结时服务商自定义回复给消费者内容
	_remark string
	// 投诉单号
	_id int64
}

// NewTmallservicecenteranomalyrecoursecloseRequest 初始化TmallservicecenteranomalyrecoursecloseAPIRequest对象
func NewTmallservicecenteranomalyrecoursecloseRequest() *TmallservicecenteranomalyrecoursecloseAPIRequest {
	return &TmallservicecenteranomalyrecoursecloseAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicecenteranomalyrecoursecloseAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.anomalyrecourse.close"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicecenteranomalyrecoursecloseAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicecenteranomalyrecoursecloseAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRemark is Remark Setter
// 完结时服务商自定义回复给消费者内容
func (r *TmallservicecenteranomalyrecoursecloseAPIRequest) SetRemark(_remark string) error {
	r._remark = _remark
	r.Set("remark", _remark)
	return nil
}

// GetRemark Remark Getter
func (r TmallservicecenteranomalyrecoursecloseAPIRequest) GetRemark() string {
	return r._remark
}

// SetId is Id Setter
// 投诉单号
func (r *TmallservicecenteranomalyrecoursecloseAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r TmallservicecenteranomalyrecoursecloseAPIRequest) GetId() int64 {
	return r._id
}
