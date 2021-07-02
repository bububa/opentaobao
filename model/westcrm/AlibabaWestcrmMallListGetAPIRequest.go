package westcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWestcrmMallListGetAPIRequest 获取商场列表 API请求
// alibaba.westcrm.mall.list.get
//
// 根据园区id获取商场列表
type AlibabaWestcrmMallListGetAPIRequest struct {
	model.Params
	// 园区id
	_campusId int64
}

// NewAlibabaWestcrmMallListGetRequest 初始化AlibabaWestcrmMallListGetAPIRequest对象
func NewAlibabaWestcrmMallListGetRequest() *AlibabaWestcrmMallListGetAPIRequest {
	return &AlibabaWestcrmMallListGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWestcrmMallListGetAPIRequest) GetApiMethodName() string {
	return "alibaba.westcrm.mall.list.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWestcrmMallListGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is CampusId Setter
// 园区id
func (r *AlibabaWestcrmMallListGetAPIRequest) SetCampusId(_campusId int64) error {
	r._campusId = _campusId
	r.Set("campus_id", _campusId)
	return nil
}

// Get CampusId Getter
func (r AlibabaWestcrmMallListGetAPIRequest) GetCampusId() int64 {
	return r._campusId
}
