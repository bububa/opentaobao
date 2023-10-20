package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaailabsaligenieopencontentpushAPIRequest 天猫精灵内容接入标准接口 API请求
// alibaba.ailabs.aligenie.opencontent.push
//
// 第三方内容接入天猫精灵内容库，供相关技能使用
type AlibabaailabsaligenieopencontentpushAPIRequest struct {
	model.Params
	// 在Aligenie开放平台创建的技能的ID
	_skillId int64
	// 详细内容列表
	_contents *BatchContent
}

// NewAlibabaailabsaligenieopencontentpushRequest 初始化AlibabaailabsaligenieopencontentpushAPIRequest对象
func NewAlibabaailabsaligenieopencontentpushRequest() *AlibabaailabsaligenieopencontentpushAPIRequest {
	return &AlibabaailabsaligenieopencontentpushAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaailabsaligenieopencontentpushAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.aligenie.opencontent.push"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaailabsaligenieopencontentpushAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaailabsaligenieopencontentpushAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSkillId is SkillId Setter
// 在Aligenie开放平台创建的技能的ID
func (r *AlibabaailabsaligenieopencontentpushAPIRequest) SetSkillId(_skillId int64) error {
	r._skillId = _skillId
	r.Set("skill_id", _skillId)
	return nil
}

// GetSkillId SkillId Getter
func (r AlibabaailabsaligenieopencontentpushAPIRequest) GetSkillId() int64 {
	return r._skillId
}

// SetContents is Contents Setter
// 详细内容列表
func (r *AlibabaailabsaligenieopencontentpushAPIRequest) SetContents(_contents *BatchContent) error {
	r._contents = _contents
	r.Set("contents", _contents)
	return nil
}

// GetContents Contents Getter
func (r AlibabaailabsaligenieopencontentpushAPIRequest) GetContents() *BatchContent {
	return r._contents
}
