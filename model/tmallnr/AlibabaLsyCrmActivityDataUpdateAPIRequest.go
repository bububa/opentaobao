package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLsyCrmActivityDataUpdateAPIRequest 私域导购数据回流接口 API请求
// alibaba.lsy.crm.activity.data.update
//
// 私域导购数据回流接口
type AlibabaLsyCrmActivityDataUpdateAPIRequest struct {
	model.Params
	// 入参对象
	_reqDTO *NrtCrmActivityStatisticsDataReq
}

// NewAlibabaLsyCrmActivityDataUpdateRequest 初始化AlibabaLsyCrmActivityDataUpdateAPIRequest对象
func NewAlibabaLsyCrmActivityDataUpdateRequest() *AlibabaLsyCrmActivityDataUpdateAPIRequest {
	return &AlibabaLsyCrmActivityDataUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaLsyCrmActivityDataUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.lsy.crm.activity.data.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaLsyCrmActivityDataUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetReqDTO is ReqDTO Setter
// 入参对象
func (r *AlibabaLsyCrmActivityDataUpdateAPIRequest) SetReqDTO(_reqDTO *NrtCrmActivityStatisticsDataReq) error {
	r._reqDTO = _reqDTO
	r.Set("req_d_t_o", _reqDTO)
	return nil
}

// GetReqDTO ReqDTO Getter
func (r AlibabaLsyCrmActivityDataUpdateAPIRequest) GetReqDTO() *NrtCrmActivityStatisticsDataReq {
	return r._reqDTO
}
