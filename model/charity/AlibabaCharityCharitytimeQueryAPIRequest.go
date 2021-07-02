package charity

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaCharityCharitytimeQueryAPIRequest 查询公益3小时公益时汇总 API请求
// alibaba.charity.charitytime.query
//
// 查询公益3小时公益时汇总
type AlibabaCharityCharitytimeQueryAPIRequest struct {
	model.Params
	// 公益类型
	_charityTypeList []string
	// 结束时间戳-毫秒时间
	_endDate int64
	// 开始时间戳-毫秒时间
	_startDate int64
	// 淘宝Uid
	_tbUid int64
	// 活动ID
	_activityId int64
	// 扩展参数
	_extParam string
}

// NewAlibabaCharityCharitytimeQueryRequest 初始化AlibabaCharityCharitytimeQueryAPIRequest对象
func NewAlibabaCharityCharitytimeQueryRequest() *AlibabaCharityCharitytimeQueryAPIRequest {
	return &AlibabaCharityCharitytimeQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCharityCharitytimeQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.charity.charitytime.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCharityCharitytimeQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is CharityTypeList Setter
// 公益类型
func (r *AlibabaCharityCharitytimeQueryAPIRequest) SetCharityTypeList(_charityTypeList []string) error {
	r._charityTypeList = _charityTypeList
	r.Set("charity_type_list", _charityTypeList)
	return nil
}

// Get CharityTypeList Getter
func (r AlibabaCharityCharitytimeQueryAPIRequest) GetCharityTypeList() []string {
	return r._charityTypeList
}

// Set is EndDate Setter
// 结束时间戳-毫秒时间
func (r *AlibabaCharityCharitytimeQueryAPIRequest) SetEndDate(_endDate int64) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// Get EndDate Getter
func (r AlibabaCharityCharitytimeQueryAPIRequest) GetEndDate() int64 {
	return r._endDate
}

// Set is StartDate Setter
// 开始时间戳-毫秒时间
func (r *AlibabaCharityCharitytimeQueryAPIRequest) SetStartDate(_startDate int64) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// Get StartDate Getter
func (r AlibabaCharityCharitytimeQueryAPIRequest) GetStartDate() int64 {
	return r._startDate
}

// Set is TbUid Setter
// 淘宝Uid
func (r *AlibabaCharityCharitytimeQueryAPIRequest) SetTbUid(_tbUid int64) error {
	r._tbUid = _tbUid
	r.Set("tb_uid", _tbUid)
	return nil
}

// Get TbUid Getter
func (r AlibabaCharityCharitytimeQueryAPIRequest) GetTbUid() int64 {
	return r._tbUid
}

// Set is ActivityId Setter
// 活动ID
func (r *AlibabaCharityCharitytimeQueryAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// Get ActivityId Getter
func (r AlibabaCharityCharitytimeQueryAPIRequest) GetActivityId() int64 {
	return r._activityId
}

// Set is ExtParam Setter
// 扩展参数
func (r *AlibabaCharityCharitytimeQueryAPIRequest) SetExtParam(_extParam string) error {
	r._extParam = _extParam
	r.Set("ext_param", _extParam)
	return nil
}

// Get ExtParam Getter
func (r AlibabaCharityCharitytimeQueryAPIRequest) GetExtParam() string {
	return r._extParam
}
