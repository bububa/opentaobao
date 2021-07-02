package xiamitrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// XiamiContentResourceActionReportAPIRequest 曲库开放平台内容行为上报接口 API请求
// xiami.content.resource.action.report
//
// 合作方对接入的曲库开放内容上报行为日志
type XiamiContentResourceActionReportAPIRequest struct {
	model.Params
	// 资源ID
	_resourceId string
	// 行为数量
	_num int64
	// 资源类型（可枚举）: song(歌曲)
	_resourceType string
	// 行为类型（可枚举）：LISTEN（主动试听）、PASSIVE_LISTEN（被动试听）
	_action string
	// 来源id，如歌单id
	_fromId string
	// 1推荐2歌单3标签
	_fromType int64
	// 用户id
	_openId string
	// 用户设备id
	_utdid string
}

// NewXiamiContentResourceActionReportRequest 初始化XiamiContentResourceActionReportAPIRequest对象
func NewXiamiContentResourceActionReportRequest() *XiamiContentResourceActionReportAPIRequest {
	return &XiamiContentResourceActionReportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r XiamiContentResourceActionReportAPIRequest) GetApiMethodName() string {
	return "xiami.content.resource.action.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r XiamiContentResourceActionReportAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ResourceId Setter
// 资源ID
func (r *XiamiContentResourceActionReportAPIRequest) SetResourceId(_resourceId string) error {
	r._resourceId = _resourceId
	r.Set("resource_id", _resourceId)
	return nil
}

// Get ResourceId Getter
func (r XiamiContentResourceActionReportAPIRequest) GetResourceId() string {
	return r._resourceId
}

// Set is Num Setter
// 行为数量
func (r *XiamiContentResourceActionReportAPIRequest) SetNum(_num int64) error {
	r._num = _num
	r.Set("num", _num)
	return nil
}

// Get Num Getter
func (r XiamiContentResourceActionReportAPIRequest) GetNum() int64 {
	return r._num
}

// Set is ResourceType Setter
// 资源类型（可枚举）: song(歌曲)
func (r *XiamiContentResourceActionReportAPIRequest) SetResourceType(_resourceType string) error {
	r._resourceType = _resourceType
	r.Set("resource_type", _resourceType)
	return nil
}

// Get ResourceType Getter
func (r XiamiContentResourceActionReportAPIRequest) GetResourceType() string {
	return r._resourceType
}

// Set is Action Setter
// 行为类型（可枚举）：LISTEN（主动试听）、PASSIVE_LISTEN（被动试听）
func (r *XiamiContentResourceActionReportAPIRequest) SetAction(_action string) error {
	r._action = _action
	r.Set("action", _action)
	return nil
}

// Get Action Getter
func (r XiamiContentResourceActionReportAPIRequest) GetAction() string {
	return r._action
}

// Set is FromId Setter
// 来源id，如歌单id
func (r *XiamiContentResourceActionReportAPIRequest) SetFromId(_fromId string) error {
	r._fromId = _fromId
	r.Set("from_id", _fromId)
	return nil
}

// Get FromId Getter
func (r XiamiContentResourceActionReportAPIRequest) GetFromId() string {
	return r._fromId
}

// Set is FromType Setter
// 1推荐2歌单3标签
func (r *XiamiContentResourceActionReportAPIRequest) SetFromType(_fromType int64) error {
	r._fromType = _fromType
	r.Set("from_type", _fromType)
	return nil
}

// Get FromType Getter
func (r XiamiContentResourceActionReportAPIRequest) GetFromType() int64 {
	return r._fromType
}

// Set is OpenId Setter
// 用户id
func (r *XiamiContentResourceActionReportAPIRequest) SetOpenId(_openId string) error {
	r._openId = _openId
	r.Set("open_id", _openId)
	return nil
}

// Get OpenId Getter
func (r XiamiContentResourceActionReportAPIRequest) GetOpenId() string {
	return r._openId
}

// Set is Utdid Setter
// 用户设备id
func (r *XiamiContentResourceActionReportAPIRequest) SetUtdid(_utdid string) error {
	r._utdid = _utdid
	r.Set("utdid", _utdid)
	return nil
}

// Get Utdid Getter
func (r XiamiContentResourceActionReportAPIRequest) GetUtdid() string {
	return r._utdid
}
