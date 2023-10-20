package westcrm

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWestcrmGradeGetAPIRequest) Reset() {
	r._campusId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWestcrmGradeGetAPIRequest) GetApiMethodName() string {
	return "alibaba.westcrm.grade.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWestcrmGradeGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWestcrmGradeGetAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaWestcrmGradeGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWestcrmGradeGetRequest()
	},
}

// GetAlibabaWestcrmGradeGetRequest 从 sync.Pool 获取 AlibabaWestcrmGradeGetAPIRequest
func GetAlibabaWestcrmGradeGetAPIRequest() *AlibabaWestcrmGradeGetAPIRequest {
	return poolAlibabaWestcrmGradeGetAPIRequest.Get().(*AlibabaWestcrmGradeGetAPIRequest)
}

// ReleaseAlibabaWestcrmGradeGetAPIRequest 将 AlibabaWestcrmGradeGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaWestcrmGradeGetAPIRequest(v *AlibabaWestcrmGradeGetAPIRequest) {
	v.Reset()
	poolAlibabaWestcrmGradeGetAPIRequest.Put(v)
}
