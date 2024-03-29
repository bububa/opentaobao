package degoperation

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoDegoperationCreateqrcodeAPIRequest 中奖生成二维码 API请求
// taobao.degoperation.createqrcode
//
// 用户中奖后，生成二维码图片链接
type TaobaoDegoperationCreateqrcodeAPIRequest struct {
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

// NewTaobaoDegoperationCreateqrcodeRequest 初始化TaobaoDegoperationCreateqrcodeAPIRequest对象
func NewTaobaoDegoperationCreateqrcodeRequest() *TaobaoDegoperationCreateqrcodeAPIRequest {
	return &TaobaoDegoperationCreateqrcodeAPIRequest{
		Params: model.NewParams(5),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoDegoperationCreateqrcodeAPIRequest) Reset() {
	r._uuid = ""
	r._degAccessToken = ""
	r._sequenceNo = ""
	r._activity = ""
	r._title = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoDegoperationCreateqrcodeAPIRequest) GetApiMethodName() string {
	return "taobao.degoperation.createqrcode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoDegoperationCreateqrcodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoDegoperationCreateqrcodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUuid is Uuid Setter
// 设备id
func (r *TaobaoDegoperationCreateqrcodeAPIRequest) SetUuid(_uuid string) error {
	r._uuid = _uuid
	r.Set("uuid", _uuid)
	return nil
}

// GetUuid Uuid Getter
func (r TaobaoDegoperationCreateqrcodeAPIRequest) GetUuid() string {
	return r._uuid
}

// SetDegAccessToken is DegAccessToken Setter
// 系统信息
func (r *TaobaoDegoperationCreateqrcodeAPIRequest) SetDegAccessToken(_degAccessToken string) error {
	r._degAccessToken = _degAccessToken
	r.Set("deg_access_token", _degAccessToken)
	return nil
}

// GetDegAccessToken DegAccessToken Getter
func (r TaobaoDegoperationCreateqrcodeAPIRequest) GetDegAccessToken() string {
	return r._degAccessToken
}

// SetSequenceNo is SequenceNo Setter
// 奖品唯一标识
func (r *TaobaoDegoperationCreateqrcodeAPIRequest) SetSequenceNo(_sequenceNo string) error {
	r._sequenceNo = _sequenceNo
	r.Set("sequence_no", _sequenceNo)
	return nil
}

// GetSequenceNo SequenceNo Getter
func (r TaobaoDegoperationCreateqrcodeAPIRequest) GetSequenceNo() string {
	return r._sequenceNo
}

// SetActivity is Activity Setter
// 活动名称
func (r *TaobaoDegoperationCreateqrcodeAPIRequest) SetActivity(_activity string) error {
	r._activity = _activity
	r.Set("activity", _activity)
	return nil
}

// GetActivity Activity Getter
func (r TaobaoDegoperationCreateqrcodeAPIRequest) GetActivity() string {
	return r._activity
}

// SetTitle is Title Setter
// 奖品名称
func (r *TaobaoDegoperationCreateqrcodeAPIRequest) SetTitle(_title string) error {
	r._title = _title
	r.Set("title", _title)
	return nil
}

// GetTitle Title Getter
func (r TaobaoDegoperationCreateqrcodeAPIRequest) GetTitle() string {
	return r._title
}

var poolTaobaoDegoperationCreateqrcodeAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoDegoperationCreateqrcodeRequest()
	},
}

// GetTaobaoDegoperationCreateqrcodeRequest 从 sync.Pool 获取 TaobaoDegoperationCreateqrcodeAPIRequest
func GetTaobaoDegoperationCreateqrcodeAPIRequest() *TaobaoDegoperationCreateqrcodeAPIRequest {
	return poolTaobaoDegoperationCreateqrcodeAPIRequest.Get().(*TaobaoDegoperationCreateqrcodeAPIRequest)
}

// ReleaseTaobaoDegoperationCreateqrcodeAPIRequest 将 TaobaoDegoperationCreateqrcodeAPIRequest 放入 sync.Pool
func ReleaseTaobaoDegoperationCreateqrcodeAPIRequest(v *TaobaoDegoperationCreateqrcodeAPIRequest) {
	v.Reset()
	poolTaobaoDegoperationCreateqrcodeAPIRequest.Put(v)
}
