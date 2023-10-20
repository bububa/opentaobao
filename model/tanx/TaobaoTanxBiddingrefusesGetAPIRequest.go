package tanx

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotanxbiddingrefusesgetAPIRequest tanx竞价失败反馈api API请求
// taobao.tanx.biddingrefuses.get
//
// 竞价失败反馈根据创意id查询API提供
type TaobaotanxbiddingrefusesgetAPIRequest struct {
	model.Params
	// dsp的创意id
	_creativeIds []string
	// dsp对应的tanx的token
	_token string
	// 起始时间
	_startTime string
	// 截止时间
	_endTime string
	// dsp在tanx的memberid
	_memberId int64
	// 1970年到现在的毫秒
	_signTime int64
}

// NewTaobaotanxbiddingrefusesgetRequest 初始化TaobaotanxbiddingrefusesgetAPIRequest对象
func NewTaobaotanxbiddingrefusesgetRequest() *TaobaotanxbiddingrefusesgetAPIRequest {
	return &TaobaotanxbiddingrefusesgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotanxbiddingrefusesgetAPIRequest) GetApiMethodName() string {
	return "taobao.tanx.biddingrefuses.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotanxbiddingrefusesgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotanxbiddingrefusesgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCreativeIds is CreativeIds Setter
// dsp的创意id
func (r *TaobaotanxbiddingrefusesgetAPIRequest) SetCreativeIds(_creativeIds []string) error {
	r._creativeIds = _creativeIds
	r.Set("creative_ids", _creativeIds)
	return nil
}

// GetCreativeIds CreativeIds Getter
func (r TaobaotanxbiddingrefusesgetAPIRequest) GetCreativeIds() []string {
	return r._creativeIds
}

// SetToken is Token Setter
// dsp对应的tanx的token
func (r *TaobaotanxbiddingrefusesgetAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r TaobaotanxbiddingrefusesgetAPIRequest) GetToken() string {
	return r._token
}

// SetStartTime is StartTime Setter
// 起始时间
func (r *TaobaotanxbiddingrefusesgetAPIRequest) SetStartTime(_startTime string) error {
	r._startTime = _startTime
	r.Set("start_time", _startTime)
	return nil
}

// GetStartTime StartTime Getter
func (r TaobaotanxbiddingrefusesgetAPIRequest) GetStartTime() string {
	return r._startTime
}

// SetEndTime is EndTime Setter
// 截止时间
func (r *TaobaotanxbiddingrefusesgetAPIRequest) SetEndTime(_endTime string) error {
	r._endTime = _endTime
	r.Set("end_time", _endTime)
	return nil
}

// GetEndTime EndTime Getter
func (r TaobaotanxbiddingrefusesgetAPIRequest) GetEndTime() string {
	return r._endTime
}

// SetMemberId is MemberId Setter
// dsp在tanx的memberid
func (r *TaobaotanxbiddingrefusesgetAPIRequest) SetMemberId(_memberId int64) error {
	r._memberId = _memberId
	r.Set("member_id", _memberId)
	return nil
}

// GetMemberId MemberId Getter
func (r TaobaotanxbiddingrefusesgetAPIRequest) GetMemberId() int64 {
	return r._memberId
}

// SetSignTime is SignTime Setter
// 1970年到现在的毫秒
func (r *TaobaotanxbiddingrefusesgetAPIRequest) SetSignTime(_signTime int64) error {
	r._signTime = _signTime
	r.Set("sign_time", _signTime)
	return nil
}

// GetSignTime SignTime Getter
func (r TaobaotanxbiddingrefusesgetAPIRequest) GetSignTime() int64 {
	return r._signTime
}
