package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallservicecenterworkcardcompleteAPIRequest 工单完结 API请求
// tmall.servicecenter.workcard.complete
//
// 工单完结
type TmallservicecenterworkcardcompleteAPIRequest struct {
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

// NewTmallservicecenterworkcardcompleteRequest 初始化TmallservicecenterworkcardcompleteAPIRequest对象
func NewTmallservicecenterworkcardcompleteRequest() *TmallservicecenterworkcardcompleteAPIRequest {
	return &TmallservicecenterworkcardcompleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallservicecenterworkcardcompleteAPIRequest) GetApiMethodName() string {
	return "tmall.servicecenter.workcard.complete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallservicecenterworkcardcompleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallservicecenterworkcardcompleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetExtJson is ExtJson Setter
// 扩展信息
func (r *TmallservicecenterworkcardcompleteAPIRequest) SetExtJson(_extJson string) error {
	r._extJson = _extJson
	r.Set("ext_json", _extJson)
	return nil
}

// GetExtJson ExtJson Getter
func (r TmallservicecenterworkcardcompleteAPIRequest) GetExtJson() string {
	return r._extJson
}

// SetLatitude is Latitude Setter
// 核销地纬度
func (r *TmallservicecenterworkcardcompleteAPIRequest) SetLatitude(_latitude string) error {
	r._latitude = _latitude
	r.Set("latitude", _latitude)
	return nil
}

// GetLatitude Latitude Getter
func (r TmallservicecenterworkcardcompleteAPIRequest) GetLatitude() string {
	return r._latitude
}

// SetLongitude is Longitude Setter
// 核销地经度
func (r *TmallservicecenterworkcardcompleteAPIRequest) SetLongitude(_longitude string) error {
	r._longitude = _longitude
	r.Set("longitude", _longitude)
	return nil
}

// GetLongitude Longitude Getter
func (r TmallservicecenterworkcardcompleteAPIRequest) GetLongitude() string {
	return r._longitude
}

// SetPicUrls is PicUrls Setter
// 图片地址回传集合
func (r *TmallservicecenterworkcardcompleteAPIRequest) SetPicUrls(_picUrls string) error {
	r._picUrls = _picUrls
	r.Set("pic_urls", _picUrls)
	return nil
}

// GetPicUrls PicUrls Getter
func (r TmallservicecenterworkcardcompleteAPIRequest) GetPicUrls() string {
	return r._picUrls
}

// SetWorkcardId is WorkcardId Setter
// 工单id
func (r *TmallservicecenterworkcardcompleteAPIRequest) SetWorkcardId(_workcardId int64) error {
	r._workcardId = _workcardId
	r.Set("workcard_id", _workcardId)
	return nil
}

// GetWorkcardId WorkcardId Getter
func (r TmallservicecenterworkcardcompleteAPIRequest) GetWorkcardId() int64 {
	return r._workcardId
}

// SetCompleteCount is CompleteCount Setter
// 完结次数
func (r *TmallservicecenterworkcardcompleteAPIRequest) SetCompleteCount(_completeCount int64) error {
	r._completeCount = _completeCount
	r.Set("complete_count", _completeCount)
	return nil
}

// GetCompleteCount CompleteCount Getter
func (r TmallservicecenterworkcardcompleteAPIRequest) GetCompleteCount() int64 {
	return r._completeCount
}

// SetSequence is Sequence Setter
// 工单完结号
func (r *TmallservicecenterworkcardcompleteAPIRequest) SetSequence(_sequence int64) error {
	r._sequence = _sequence
	r.Set("sequence", _sequence)
	return nil
}

// GetSequence Sequence Getter
func (r TmallservicecenterworkcardcompleteAPIRequest) GetSequence() int64 {
	return r._sequence
}
