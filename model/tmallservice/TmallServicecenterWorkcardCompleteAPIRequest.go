package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallServicecenterWorkcardCompleteAPIRequest 工单完结 API请求
// tmall.servicecenter.workcard.complete
//
// 工单完结
type TmallServicecenterWorkcardCompleteAPIRequest struct {
	model.Params
	// 扩展信息
	_extJson string
	// 核销地纬度
	_latitude string
	// 核销地经度
	_longitude string
	// 图片地址回传集合
	_picUrls string
	// 工单id
	_workcardId int64
	// 完结次数
	_completeCount int64
	// 工单完结号
	_sequence int64
}

// NewTmallServicecenterWorkcardCompleteRequest 初始化TmallServicecenterWorkcardCompleteAPIRequest对象
func NewTmallServicecenterWorkcardCompleteRequest() *TmallServicecenterWorkcardCompleteAPIRequest {
	return &TmallServicecenterWorkcardCompleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallServicecenterWorkcardCompleteAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.workcard.complete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallServicecenterWorkcardCompleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallServicecenterWorkcardCompleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetExtJson is ExtJson Setter
// 扩展信息
func (r *TmallServicecenterWorkcardCompleteAPIRequest) SetExtJson(_extJson string) error {
	r._extJson = _extJson
	r.Set("ext_json", _extJson)
	return nil
}

// GetExtJson ExtJson Getter
func (r TmallServicecenterWorkcardCompleteAPIRequest) GetExtJson() string {
	return r._extJson
}

// SetLatitude is Latitude Setter
// 核销地纬度
func (r *TmallServicecenterWorkcardCompleteAPIRequest) SetLatitude(_latitude string) error {
	r._latitude = _latitude
	r.Set("latitude", _latitude)
	return nil
}

// GetLatitude Latitude Getter
func (r TmallServicecenterWorkcardCompleteAPIRequest) GetLatitude() string {
	return r._latitude
}

// SetLongitude is Longitude Setter
// 核销地经度
func (r *TmallServicecenterWorkcardCompleteAPIRequest) SetLongitude(_longitude string) error {
	r._longitude = _longitude
	r.Set("longitude", _longitude)
	return nil
}

// GetLongitude Longitude Getter
func (r TmallServicecenterWorkcardCompleteAPIRequest) GetLongitude() string {
	return r._longitude
}

// SetPicUrls is PicUrls Setter
// 图片地址回传集合
func (r *TmallServicecenterWorkcardCompleteAPIRequest) SetPicUrls(_picUrls string) error {
	r._picUrls = _picUrls
	r.Set("pic_urls", _picUrls)
	return nil
}

// GetPicUrls PicUrls Getter
func (r TmallServicecenterWorkcardCompleteAPIRequest) GetPicUrls() string {
	return r._picUrls
}

// SetWorkcardId is WorkcardId Setter
// 工单id
func (r *TmallServicecenterWorkcardCompleteAPIRequest) SetWorkcardId(_workcardId int64) error {
	r._workcardId = _workcardId
	r.Set("workcard_id", _workcardId)
	return nil
}

// GetWorkcardId WorkcardId Getter
func (r TmallServicecenterWorkcardCompleteAPIRequest) GetWorkcardId() int64 {
	return r._workcardId
}

// SetCompleteCount is CompleteCount Setter
// 完结次数
func (r *TmallServicecenterWorkcardCompleteAPIRequest) SetCompleteCount(_completeCount int64) error {
	r._completeCount = _completeCount
	r.Set("complete_count", _completeCount)
	return nil
}

// GetCompleteCount CompleteCount Getter
func (r TmallServicecenterWorkcardCompleteAPIRequest) GetCompleteCount() int64 {
	return r._completeCount
}

// SetSequence is Sequence Setter
// 工单完结号
func (r *TmallServicecenterWorkcardCompleteAPIRequest) SetSequence(_sequence int64) error {
	r._sequence = _sequence
	r.Set("sequence", _sequence)
	return nil
}

// GetSequence Sequence Getter
func (r TmallServicecenterWorkcardCompleteAPIRequest) GetSequence() int64 {
	return r._sequence
}
