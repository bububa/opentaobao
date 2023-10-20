package interact

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractActivityRegisterAPIRequest ISV互动应用活动注册服务 API请求
// alibaba.interact.activity.register
//
// 为支持卖家由ISV互动应用可以在手淘店铺首页透出，提供ISV互动应用创建的活动注册到手淘的服务
type AlibabaInteractActivityRegisterAPIRequest struct {
	model.Params
	// 页面内容链接，仅允许ascii字符
	_entryUrl string
	// 活动ID
	_bizId string
	// 活动描述文字
	_description string
	// 活动结束时间
	_endTime string
	// 活动名称
	_name string
	// 活动封面图片（非必填）
	_picture string
	// 活动开始时间
	_startTime string
	// 是否有有效时间，若为真开始时间和结束时间必填，默认为真
	_hasValidTime bool
}

// NewAlibabaInteractActivityRegisterRequest 初始化AlibabaInteractActivityRegisterAPIRequest对象
func NewAlibabaInteractActivityRegisterRequest() *AlibabaInteractActivityRegisterAPIRequest {
	return &AlibabaInteractActivityRegisterAPIRequest{
		Params: model.NewParams(8),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaInteractActivityRegisterAPIRequest) Reset() {
	r._entryUrl = ""
	r._bizId = ""
	r._description = ""
	r._endTime = ""
	r._name = ""
	r._picture = ""
	r._startTime = ""
	r._hasValidTime = false
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractActivityRegisterAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.activity.register"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractActivityRegisterAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaInteractActivityRegisterAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEntryUrl is EntryUrl Setter
// 页面内容链接，仅允许ascii字符
func (r *AlibabaInteractActivityRegisterAPIRequest) SetEntryUrl(_entryUrl string) error {
	r._entryUrl = _entryUrl
	r.Set("entry_url", _entryUrl)
	return nil
}

// GetEntryUrl EntryUrl Getter
func (r AlibabaInteractActivityRegisterAPIRequest) GetEntryUrl() string {
	return r._entryUrl
}

// SetBizId is BizId Setter
// 活动ID
func (r *AlibabaInteractActivityRegisterAPIRequest) SetBizId(_bizId string) error {
	r._bizId = _bizId
	r.Set("biz_id", _bizId)
	return nil
}

// GetBizId BizId Getter
func (r AlibabaInteractActivityRegisterAPIRequest) GetBizId() string {
	return r._bizId
}

// SetDescription is Description Setter
// 活动描述文字
func (r *AlibabaInteractActivityRegisterAPIRequest) SetDescription(_description string) error {
	r._description = _description
	r.Set("description", _description)
	return nil
}

// GetDescription Description Getter
func (r AlibabaInteractActivityRegisterAPIRequest) GetDescription() string {
	return r._description
}

// SetEndTime is EndTime Setter
// 活动结束时间
func (r *AlibabaInteractActivityRegisterAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r AlibabaInteractActivityRegisterAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetName is Name Setter
// 活动名称
func (r *AlibabaInteractActivityRegisterAPIRequest) SetName(_name string) error {
	r._name = _name
	r.Set("name", _name)
	return nil
}

// GetName Name Getter
func (r AlibabaInteractActivityRegisterAPIRequest) GetName() string {
	return r._name
}

// SetPicture is Picture Setter
// 活动封面图片（非必填）
func (r *AlibabaInteractActivityRegisterAPIRequest) SetPicture(_picture string) error {
	r._picture = _picture
	r.Set("picture", _picture)
	return nil
}

// GetPicture Picture Getter
func (r AlibabaInteractActivityRegisterAPIRequest) GetPicture() string {
	return r._picture
}

// SetStartTime is StartTime Setter
// 活动开始时间
func (r *AlibabaInteractActivityRegisterAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r AlibabaInteractActivityRegisterAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetHasValidTime is HasValidTime Setter
// 是否有有效时间，若为真开始时间和结束时间必填，默认为真
func (r *AlibabaInteractActivityRegisterAPIRequest) SetHasValidTime(_hasValidTime bool) error {
	r._hasValidTime = _hasValidTime
	r.Set("has_valid_time", _hasValidTime)
	return nil
}

// GetHasValidTime HasValidTime Getter
func (r AlibabaInteractActivityRegisterAPIRequest) GetHasValidTime() bool {
	return r._hasValidTime
}

var poolAlibabaInteractActivityRegisterAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaInteractActivityRegisterRequest()
	},
}

// GetAlibabaInteractActivityRegisterRequest 从 sync.Pool 获取 AlibabaInteractActivityRegisterAPIRequest
func GetAlibabaInteractActivityRegisterAPIRequest() *AlibabaInteractActivityRegisterAPIRequest {
	return poolAlibabaInteractActivityRegisterAPIRequest.Get().(*AlibabaInteractActivityRegisterAPIRequest)
}

// ReleaseAlibabaInteractActivityRegisterAPIRequest 将 AlibabaInteractActivityRegisterAPIRequest 放入 sync.Pool
func ReleaseAlibabaInteractActivityRegisterAPIRequest(v *AlibabaInteractActivityRegisterAPIRequest) {
	v.Reset()
	poolAlibabaInteractActivityRegisterAPIRequest.Put(v)
}
