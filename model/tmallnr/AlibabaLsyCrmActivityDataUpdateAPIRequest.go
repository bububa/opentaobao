package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabalsycrmactivitydataupdateAPIRequest 私域导购数据回流接口 API请求
// alibaba.lsy.crm.activity.data.update
//
// 私域导购数据回流接口
type AlibabalsycrmactivitydataupdateAPIRequest struct {
	model.Params
	// 入参对象
	_reqDTO *NrtCrmActivityStatisticsDataReq
}

// NewAlibabalsycrmactivitydataupdateRequest 初始化AlibabalsycrmactivitydataupdateAPIRequest对象
func NewAlibabalsycrmactivitydataupdateRequest() *AlibabalsycrmactivitydataupdateAPIRequest {
	return &AlibabalsycrmactivitydataupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabalsycrmactivitydataupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.lsy.crm.activity.data.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabalsycrmactivitydataupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabalsycrmactivitydataupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReqDTO is ReqDTO Setter
// 入参对象
func (r *AlibabalsycrmactivitydataupdateAPIRequest) SetReqDTO(_reqDTO *NrtCrmActivityStatisticsDataReq) error {
	r._reqDTO = _reqDTO
	r.Set("req_d_t_o", _reqDTO)
	return nil
}

// GetReqDTO ReqDTO Getter
func (r AlibabalsycrmactivitydataupdateAPIRequest) GetReqDTO() *NrtCrmActivityStatisticsDataReq {
	return r._reqDTO
}
