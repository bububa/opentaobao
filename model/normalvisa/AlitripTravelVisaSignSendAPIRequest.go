package normalvisa

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitriptravelvisasignsendAPIRequest 签证批量申请人送签接口 API请求
// alitrip.travel.visa.sign.send
//
// 签证批量申请人送签接口，用于商家批量送签。
type AlitriptravelvisasignsendAPIRequest struct {
	model.Params
	// 申请人ids
	_applyIds []string
	// 国家id。目前只支持越南，越南国家id:27027
	_nationId int64
	// 送签类型：1-非加急，2-加急，默认非加急
	_signType int64
}

// NewAlitriptravelvisasignsendRequest 初始化AlitriptravelvisasignsendAPIRequest对象
func NewAlitriptravelvisasignsendRequest() *AlitriptravelvisasignsendAPIRequest {
	return &AlitriptravelvisasignsendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitriptravelvisasignsendAPIRequest) GetApiMethodName() string {
	return "alitrip.travel.visa.sign.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitriptravelvisasignsendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitriptravelvisasignsendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetApplyIds is ApplyIds Setter
// 申请人ids
func (r *AlitriptravelvisasignsendAPIRequest) SetApplyIds(_applyIds []string) error {
	r._applyIds = _applyIds
	r.Set("apply_ids", _applyIds)
	return nil
}

// GetApplyIds ApplyIds Getter
func (r AlitriptravelvisasignsendAPIRequest) GetApplyIds() []string {
	return r._applyIds
}

// SetNationId is NationId Setter
// 国家id。目前只支持越南，越南国家id:27027
func (r *AlitriptravelvisasignsendAPIRequest) SetNationId(_nationId int64) error {
	r._nationId = _nationId
	r.Set("nation_id", _nationId)
	return nil
}

// GetNationId NationId Getter
func (r AlitriptravelvisasignsendAPIRequest) GetNationId() int64 {
	return r._nationId
}

// SetSignType is SignType Setter
// 送签类型：1-非加急，2-加急，默认非加急
func (r *AlitriptravelvisasignsendAPIRequest) SetSignType(_signType int64) error {
	r._signType = _signType
	r.Set("sign_type", _signType)
	return nil
}

// GetSignType SignType Getter
func (r AlitriptravelvisasignsendAPIRequest) GetSignType() int64 {
	return r._signType
}
