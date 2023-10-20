package vaccin

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabahealthvaccinsubscribeinforeturnAPIRequest 自有pov预约信息回传 API请求
// alibaba.health.vaccin.subscribe.info.return
//
// 自有pov预约信息回传
type AlibabahealthvaccinsubscribeinforeturnAPIRequest struct {
	model.Params
	// 医鹿的苗种，疫苗id
	_vaccineIds string
	// 取数日期，如2023-05-25 21:02:19
	_obtainDate string
	// 当前页页号，注意页号是从1开始的
	_pageNo int64
	// 分页模型每页的大小
	_pageSize int64
}

// NewAlibabahealthvaccinsubscribeinforeturnRequest 初始化AlibabahealthvaccinsubscribeinforeturnAPIRequest对象
func NewAlibabahealthvaccinsubscribeinforeturnRequest() *AlibabahealthvaccinsubscribeinforeturnAPIRequest {
	return &AlibabahealthvaccinsubscribeinforeturnAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabahealthvaccinsubscribeinforeturnAPIRequest) GetApiMethodName() string {
	return "alibaba.health.vaccin.subscribe.info.return"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabahealthvaccinsubscribeinforeturnAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabahealthvaccinsubscribeinforeturnAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetVaccineIds is VaccineIds Setter
// 医鹿的苗种，疫苗id
func (r *AlibabahealthvaccinsubscribeinforeturnAPIRequest) SetVaccineIds(_vaccineIds string) error {
	r._vaccineIds = _vaccineIds
	r.Set("vaccine_ids", _vaccineIds)
	return nil
}

// GetVaccineIds VaccineIds Getter
func (r AlibabahealthvaccinsubscribeinforeturnAPIRequest) GetVaccineIds() string {
	return r._vaccineIds
}

// SetObtainDate is ObtainDate Setter
// 取数日期，如2023-05-25 21:02:19
func (r *AlibabahealthvaccinsubscribeinforeturnAPIRequest) SetObtainDate(_obtainDate string) error {
	r._obtainDate = _obtainDate
	r.Set("obtain_date", _obtainDate)
	return nil
}

// GetObtainDate ObtainDate Getter
func (r AlibabahealthvaccinsubscribeinforeturnAPIRequest) GetObtainDate() string {
	return r._obtainDate
}

// SetPageNo is PageNo Setter
// 当前页页号，注意页号是从1开始的
func (r *AlibabahealthvaccinsubscribeinforeturnAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r AlibabahealthvaccinsubscribeinforeturnAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 分页模型每页的大小
func (r *AlibabahealthvaccinsubscribeinforeturnAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AlibabahealthvaccinsubscribeinforeturnAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
