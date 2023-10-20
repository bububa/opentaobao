package xiamitrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// XiamicontentresourceactionreportAPIRequest 曲库开放平台内容行为上报接口 API请求
// xiami.content.resource.action.report
//
// 合作方对接入的曲库开放内容上报行为日志
type XiamicontentresourceactionreportAPIRequest struct {
	model.Params
	// 资源ID
	_resourceId string
	// 行为类型（可枚举）：LISTEN（主动试听）、PASSIVE_LISTEN（被动试听）
	_action string
	// 资源类型（可枚举）: song(歌曲)
	_resourceType string
	// 来源id，如歌单id
	_fromId string
	// 用户id
	_openId string
	// 用户设备id
	_utdid string
	// 扩展信息
	_extra string
	// 音频id
	_relationId string
	// 关联类型（如音频）
	_relationType string
	// 行为数量
	_num int64
	// 1推荐2歌单3标签
	_fromType int64
}

// NewXiamicontentresourceactionreportRequest 初始化XiamicontentresourceactionreportAPIRequest对象
func NewXiamicontentresourceactionreportRequest() *XiamicontentresourceactionreportAPIRequest {
	return &XiamicontentresourceactionreportAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r XiamicontentresourceactionreportAPIRequest) GetApiMethodName() string {
	return "xiami.content.resource.action.report"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r XiamicontentresourceactionreportAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r XiamicontentresourceactionreportAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetResourceId is ResourceId Setter
// 资源ID
func (r *XiamicontentresourceactionreportAPIRequest) SetResourceId(_resourceId string) error {
	r._resourceId = _resourceId
	r.Set("resource_id", _resourceId)
	return nil
}

// GetResourceId ResourceId Getter
func (r XiamicontentresourceactionreportAPIRequest) GetResourceId() string {
	return r._resourceId
}

// SetAction is Action Setter
// 行为类型（可枚举）：LISTEN（主动试听）、PASSIVE_LISTEN（被动试听）
func (r *XiamicontentresourceactionreportAPIRequest) SetAction(_action string) error {
	r._action = _action
	r.Set("action", _action)
	return nil
}

// GetAction Action Getter
func (r XiamicontentresourceactionreportAPIRequest) GetAction() string {
	return r._action
}

// SetResourceType is ResourceType Setter
// 资源类型（可枚举）: song(歌曲)
func (r *XiamicontentresourceactionreportAPIRequest) SetResourceType(_resourceType string) error {
	r._resourceType = _resourceType
	r.Set("resource_type", _resourceType)
	return nil
}

// GetResourceType ResourceType Getter
func (r XiamicontentresourceactionreportAPIRequest) GetResourceType() string {
	return r._resourceType
}

// SetFromId is FromId Setter
// 来源id，如歌单id
func (r *XiamicontentresourceactionreportAPIRequest) SetFromId(_fromId string) error {
	r._fromId = _fromId
	r.Set("from_id", _fromId)
	return nil
}

// GetFromId FromId Getter
func (r XiamicontentresourceactionreportAPIRequest) GetFromId() string {
	return r._fromId
}

// SetOpenId is OpenId Setter
// 用户id
func (r *XiamicontentresourceactionreportAPIRequest) SetOpenId(_openId string) error {
	r._openId = _openId
	r.Set("open_id", _openId)
	return nil
}

// GetOpenId OpenId Getter
func (r XiamicontentresourceactionreportAPIRequest) GetOpenId() string {
	return r._openId
}

// SetUtdid is Utdid Setter
// 用户设备id
func (r *XiamicontentresourceactionreportAPIRequest) SetUtdid(_utdid string) error {
	r._utdid = _utdid
	r.Set("utdid", _utdid)
	return nil
}

// GetUtdid Utdid Getter
func (r XiamicontentresourceactionreportAPIRequest) GetUtdid() string {
	return r._utdid
}

// SetExtra is Extra Setter
// 扩展信息
func (r *XiamicontentresourceactionreportAPIRequest) SetExtra(_extra string) error {
	r._extra = _extra
	r.Set("extra", _extra)
	return nil
}

// GetExtra Extra Getter
func (r XiamicontentresourceactionreportAPIRequest) GetExtra() string {
	return r._extra
}

// SetRelationId is RelationId Setter
// 音频id
func (r *XiamicontentresourceactionreportAPIRequest) SetRelationId(_relationId string) error {
	r._relationId = _relationId
	r.Set("relation_id", _relationId)
	return nil
}

// GetRelationId RelationId Getter
func (r XiamicontentresourceactionreportAPIRequest) GetRelationId() string {
	return r._relationId
}

// SetRelationType is RelationType Setter
// 关联类型（如音频）
func (r *XiamicontentresourceactionreportAPIRequest) SetRelationType(_relationType string) error {
	r._relationType = _relationType
	r.Set("relation_type", _relationType)
	return nil
}

// GetRelationType RelationType Getter
func (r XiamicontentresourceactionreportAPIRequest) GetRelationType() string {
	return r._relationType
}

// SetNum is Num Setter
// 行为数量
func (r *XiamicontentresourceactionreportAPIRequest) SetNum(_num int64) error {
	r._num = _num
	r.Set("num", _num)
	return nil
}

// GetNum Num Getter
func (r XiamicontentresourceactionreportAPIRequest) GetNum() int64 {
	return r._num
}

// SetFromType is FromType Setter
// 1推荐2歌单3标签
func (r *XiamicontentresourceactionreportAPIRequest) SetFromType(_fromType int64) error {
	r._fromType = _fromType
	r.Set("from_type", _fromType)
	return nil
}

// GetFromType FromType Getter
func (r XiamicontentresourceactionreportAPIRequest) GetFromType() int64 {
	return r._fromType
}
