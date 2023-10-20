package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsAligenieOpencontentPushAPIRequest 天猫精灵内容接入标准接口 API请求
// alibaba.ailabs.aligenie.opencontent.push
//
// 第三方内容接入天猫精灵内容库，供相关技能使用
type AlibabaAilabsAligenieOpencontentPushAPIRequest struct {
	model.Params
	// 在Aligenie开放平台创建的技能的ID
	_skillId int64
	// 详细内容列表
	_contents *BatchContent
}

// NewAlibabaAilabsAligenieOpencontentPushRequest 初始化AlibabaAilabsAligenieOpencontentPushAPIRequest对象
func NewAlibabaAilabsAligenieOpencontentPushRequest() *AlibabaAilabsAligenieOpencontentPushAPIRequest {
	return &AlibabaAilabsAligenieOpencontentPushAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAilabsAligenieOpencontentPushAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.aligenie.opencontent.push"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAilabsAligenieOpencontentPushAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAilabsAligenieOpencontentPushAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSkillId is SkillId Setter
// 在Aligenie开放平台创建的技能的ID
func (r *AlibabaAilabsAligenieOpencontentPushAPIRequest) SetSkillId(_skillId int64) error {
	r._skillId = _skillId
	r.Set("skill_id", _skillId)
	return nil
}

// GetSkillId SkillId Getter
func (r AlibabaAilabsAligenieOpencontentPushAPIRequest) GetSkillId() int64 {
	return r._skillId
}

// SetContents is Contents Setter
// 详细内容列表
func (r *AlibabaAilabsAligenieOpencontentPushAPIRequest) SetContents(_contents *BatchContent) error {
	r._contents = _contents
	r.Set("contents", _contents)
	return nil
}

// GetContents Contents Getter
func (r AlibabaAilabsAligenieOpencontentPushAPIRequest) GetContents() *BatchContent {
	return r._contents
}
