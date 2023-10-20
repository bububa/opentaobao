package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugcodeuserdataAPIRequest 西安杨森同步用户行为接口 API请求
// alibaba.alihealth.drugcode.user.data
//
// 西安杨森同步用户行为接口
type AlibabaalihealthdrugcodeuserdataAPIRequest struct {
	model.Params
	// 用户信息
	_list []HaoxinqingDataDto
	// 企业ID，用户区分 appkey下不同企业数据隔离的
	_refEntId string
}

// NewAlibabaalihealthdrugcodeuserdataRequest 初始化AlibabaalihealthdrugcodeuserdataAPIRequest对象
func NewAlibabaalihealthdrugcodeuserdataRequest() *AlibabaalihealthdrugcodeuserdataAPIRequest {
	return &AlibabaalihealthdrugcodeuserdataAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugcodeuserdataAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drugcode.user.data"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugcodeuserdataAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugcodeuserdataAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetList is List Setter
// 用户信息
func (r *AlibabaalihealthdrugcodeuserdataAPIRequest) SetList(_list []HaoxinqingDataDto) error {
	r._list = _list
	r.Set("list", _list)
	return nil
}

// GetList List Getter
func (r AlibabaalihealthdrugcodeuserdataAPIRequest) GetList() []HaoxinqingDataDto {
	return r._list
}

// SetRefEntId is RefEntId Setter
// 企业ID，用户区分 appkey下不同企业数据隔离的
func (r *AlibabaalihealthdrugcodeuserdataAPIRequest) SetRefEntId(_refEntId string) error {
	r._refEntId = _refEntId
	r.Set("ref_ent_id", _refEntId)
	return nil
}

// GetRefEntId RefEntId Getter
func (r AlibabaalihealthdrugcodeuserdataAPIRequest) GetRefEntId() string {
	return r._refEntId
}
