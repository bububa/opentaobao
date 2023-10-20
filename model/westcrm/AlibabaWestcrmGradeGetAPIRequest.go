package westcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawestcrmgradegetAPIRequest 获取等级列表 API请求
// alibaba.westcrm.grade.get
//
// 获取会员卡等级列表
type AlibabawestcrmgradegetAPIRequest struct {
	model.Params
	// 园区id
	_campusId int64
}

// NewAlibabawestcrmgradegetRequest 初始化AlibabawestcrmgradegetAPIRequest对象
func NewAlibabawestcrmgradegetRequest() *AlibabawestcrmgradegetAPIRequest {
	return &AlibabawestcrmgradegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawestcrmgradegetAPIRequest) GetApiMethodName() string {
	return "alibaba.westcrm.grade.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawestcrmgradegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawestcrmgradegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCampusId is CampusId Setter
// 园区id
func (r *AlibabawestcrmgradegetAPIRequest) SetCampusId(_campusId int64) error {
	r._campusId = _campusId
	r.Set("campus_id", _campusId)
	return nil
}

// GetCampusId CampusId Getter
func (r AlibabawestcrmgradegetAPIRequest) GetCampusId() int64 {
	return r._campusId
}
