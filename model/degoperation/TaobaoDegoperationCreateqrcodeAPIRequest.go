package degoperation

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaodegoperationcreateqrcodeAPIRequest 中奖生成二维码 API请求
// taobao.degoperation.createqrcode
//
// 用户中奖后，生成二维码图片链接
type TaobaodegoperationcreateqrcodeAPIRequest struct {
	model.Params
	// 设备id
	_uuid string
	// 系统信息
	_degAccessToken string
	// 奖品唯一标识
	_sequenceNo string
	// 活动名称
	_activity string
	// 奖品名称
	_title string
}

// NewTaobaodegoperationcreateqrcodeRequest 初始化TaobaodegoperationcreateqrcodeAPIRequest对象
func NewTaobaodegoperationcreateqrcodeRequest() *TaobaodegoperationcreateqrcodeAPIRequest {
	return &TaobaodegoperationcreateqrcodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaodegoperationcreateqrcodeAPIRequest) GetApiMethodName() string {
	return "taobao.degoperation.createqrcode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaodegoperationcreateqrcodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaodegoperationcreateqrcodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUuid is Uuid Setter
// 设备id
func (r *TaobaodegoperationcreateqrcodeAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r TaobaodegoperationcreateqrcodeAPIRequest) GetUuid() string {
	return r._uuid
}

// SetDegAccessToken is DegAccessToken Setter
// 系统信息
func (r *TaobaodegoperationcreateqrcodeAPIRequest) SetDegAccessToken(_degAccessToken string) error {
	r._degAccessToken = _degAccessToken
	r.Set("deg_access_token", _degAccessToken)
	return nil
}

// GetDegAccessToken DegAccessToken Getter
func (r TaobaodegoperationcreateqrcodeAPIRequest) GetDegAccessToken() string {
	return r._degAccessToken
}

// SetSequenceNo is SequenceNo Setter
// 奖品唯一标识
func (r *TaobaodegoperationcreateqrcodeAPIRequest) SetSequenceNo(_sequenceNo string) error {
	r._sequenceNo = _sequenceNo
	r.Set("sequence_no", _sequenceNo)
	return nil
}

// GetSequenceNo SequenceNo Getter
func (r TaobaodegoperationcreateqrcodeAPIRequest) GetSequenceNo() string {
	return r._sequenceNo
}

// SetActivity is Activity Setter
// 活动名称
func (r *TaobaodegoperationcreateqrcodeAPIRequest) SetActivity(_activity string) error {
	r._activity = _activity
	r.Set("activity", _activity)
	return nil
}

// GetActivity Activity Getter
func (r TaobaodegoperationcreateqrcodeAPIRequest) GetActivity() string {
	return r._activity
}

// SetTitle is Title Setter
// 奖品名称
func (r *TaobaodegoperationcreateqrcodeAPIRequest) SetTitle(_title string) error {
	r._title = _title
	r.Set("title", _title)
	return nil
}

// GetTitle Title Getter
func (r TaobaodegoperationcreateqrcodeAPIRequest) GetTitle() string {
	return r._title
}
