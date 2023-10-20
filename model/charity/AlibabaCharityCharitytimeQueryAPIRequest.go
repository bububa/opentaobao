package charity

import (
	"net/url"
	"sync"

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
	// 扩展参数
	_extParam string
	// 活动ID
	_activityId int64
	// 结束时间戳-毫秒时间
	_endDate int64
	// 开始时间戳-毫秒时间
	_startDate int64
	// 淘宝Uid
	_tbUid int64
}

// NewAlibabaCharityCharitytimeQueryRequest 初始化AlibabaCharityCharitytimeQueryAPIRequest对象
func NewAlibabaCharityCharitytimeQueryRequest() *AlibabaCharityCharitytimeQueryAPIRequest {
	return &AlibabaCharityCharitytimeQueryAPIRequest{
		Params: model.NewParams(6),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaCharityCharitytimeQueryAPIRequest) Reset() {
	r._charityTypeList = r._charityTypeList[:0]
	r._extParam = ""
	r._activityId = 0
	r._endDate = 0
	r._startDate = 0
	r._tbUid = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaCharityCharitytimeQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.charity.charitytime.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaCharityCharitytimeQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaCharityCharitytimeQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCharityTypeList is CharityTypeList Setter
// 公益类型
func (r *AlibabaCharityCharitytimeQueryAPIRequest) SetCharityTypeList(_charityTypeList []string) error {
	r._charityTypeList = _charityTypeList
	r.Set("charity_type_list", _charityTypeList)
	return nil
}

// GetCharityTypeList CharityTypeList Getter
func (r AlibabaCharityCharitytimeQueryAPIRequest) GetCharityTypeList() []string {
	return r._charityTypeList
}

// SetExtParam is ExtParam Setter
// 扩展参数
func (r *AlibabaCharityCharitytimeQueryAPIRequest) SetExtParam(_extParam string) error {
	r._extParam = _extParam
	r.Set("ext_param", _extParam)
	return nil
}

// GetExtParam ExtParam Getter
func (r AlibabaCharityCharitytimeQueryAPIRequest) GetExtParam() string {
	return r._extParam
}

// SetActivityId is ActivityId Setter
// 活动ID
func (r *AlibabaCharityCharitytimeQueryAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r AlibabaCharityCharitytimeQueryAPIRequest) GetActivityId() int64 {
	return r._activityId
}

// SetEndDate is EndDate Setter
// 结束时间戳-毫秒时间
func (r *AlibabaCharityCharitytimeQueryAPIRequest) SetEndDate(_endDate int64) error {
	r._endDate = _endDate
	r.Set("end_date", _endDate)
	return nil
}

// GetEndDate EndDate Getter
func (r AlibabaCharityCharitytimeQueryAPIRequest) GetEndDate() int64 {
	return r._endDate
}

// SetStartDate is StartDate Setter
// 开始时间戳-毫秒时间
func (r *AlibabaCharityCharitytimeQueryAPIRequest) SetStartDate(_startDate int64) error {
	r._startDate = _startDate
	r.Set("start_date", _startDate)
	return nil
}

// GetStartDate StartDate Getter
func (r AlibabaCharityCharitytimeQueryAPIRequest) GetStartDate() int64 {
	return r._startDate
}

// SetTbUid is TbUid Setter
// 淘宝Uid
func (r *AlibabaCharityCharitytimeQueryAPIRequest) SetTbUid(_tbUid int64) error {
	r._tbUid = _tbUid
	r.Set("tb_uid", _tbUid)
	return nil
}

// GetTbUid TbUid Getter
func (r AlibabaCharityCharitytimeQueryAPIRequest) GetTbUid() int64 {
	return r._tbUid
}

var poolAlibabaCharityCharitytimeQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaCharityCharitytimeQueryRequest()
	},
}

// GetAlibabaCharityCharitytimeQueryRequest 从 sync.Pool 获取 AlibabaCharityCharitytimeQueryAPIRequest
func GetAlibabaCharityCharitytimeQueryAPIRequest() *AlibabaCharityCharitytimeQueryAPIRequest {
	return poolAlibabaCharityCharitytimeQueryAPIRequest.Get().(*AlibabaCharityCharitytimeQueryAPIRequest)
}

// ReleaseAlibabaCharityCharitytimeQueryAPIRequest 将 AlibabaCharityCharitytimeQueryAPIRequest 放入 sync.Pool
func ReleaseAlibabaCharityCharitytimeQueryAPIRequest(v *AlibabaCharityCharitytimeQueryAPIRequest) {
	v.Reset()
	poolAlibabaCharityCharitytimeQueryAPIRequest.Put(v)
}
