package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaofeedflowitemadgroupadzoneunbindAPIRequest 信息流单元内解绑资源位 API请求
// taobao.feedflow.item.adgroup.adzone.unbind
//
// 信息流单元内解绑资源位
type TaobaofeedflowitemadgroupadzoneunbindAPIRequest struct {
	model.Params
	// 广告位id
	_adzoneIdList []string
	// 单元id
	_adgroupId int64
}

// NewTaobaofeedflowitemadgroupadzoneunbindRequest 初始化TaobaofeedflowitemadgroupadzoneunbindAPIRequest对象
func NewTaobaofeedflowitemadgroupadzoneunbindRequest() *TaobaofeedflowitemadgroupadzoneunbindAPIRequest {
	return &TaobaofeedflowitemadgroupadzoneunbindAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaofeedflowitemadgroupadzoneunbindAPIRequest) GetApiMethodName() string {
	return "taobao.feedflow.item.adgroup.adzone.unbind"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaofeedflowitemadgroupadzoneunbindAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaofeedflowitemadgroupadzoneunbindAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAdzoneIdList is AdzoneIdList Setter
// 广告位id
func (r *TaobaofeedflowitemadgroupadzoneunbindAPIRequest) SetAdzoneIdList(_adzoneIdList []string) error {
	r._adzoneIdList = _adzoneIdList
	r.Set("adzone_id_list", _adzoneIdList)
	return nil
}

// GetAdzoneIdList AdzoneIdList Getter
func (r TaobaofeedflowitemadgroupadzoneunbindAPIRequest) GetAdzoneIdList() []string {
	return r._adzoneIdList
}

// SetAdgroupId is AdgroupId Setter
// 单元id
func (r *TaobaofeedflowitemadgroupadzoneunbindAPIRequest) SetAdgroupId(_adgroupId int64) error {
	r._adgroupId = _adgroupId
	r.Set("adgroup_id", _adgroupId)
	return nil
}

// GetAdgroupId AdgroupId Getter
func (r TaobaofeedflowitemadgroupadzoneunbindAPIRequest) GetAdgroupId() int64 {
	return r._adgroupId
}
