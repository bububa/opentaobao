package westcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWestcrmGradeGetAPIRequest 获取等级列表 API请求
// alibaba.westcrm.grade.get
//
// 获取会员卡等级列表
type AlibabaWestcrmGradeGetAPIRequest struct {
	model.Params
	// 园区id
	_campusId int64
}

// NewAlibabaWestcrmGradeGetRequest 初始化AlibabaWestcrmGradeGetAPIRequest对象
func NewAlibabaWestcrmGradeGetRequest() *AlibabaWestcrmGradeGetAPIRequest {
	return &AlibabaWestcrmGradeGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWestcrmGradeGetAPIRequest) GetApiMethodName() string {
	return "alibaba.westcrm.grade.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWestcrmGradeGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetCampusId is CampusId Setter
// 园区id
func (r *AlibabaWestcrmGradeGetAPIRequest) SetCampusId(_campusId int64) error {
	r._campusId = _campusId
	r.Set("campus_id", _campusId)
	return nil
}

// GetCampusId CampusId Getter
func (r AlibabaWestcrmGradeGetAPIRequest) GetCampusId() int64 {
	return r._campusId
}
