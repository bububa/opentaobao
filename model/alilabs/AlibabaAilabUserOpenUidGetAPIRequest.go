package alilabs

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabUserOpenUidGetAPIRequest access token 获取精灵用户 id API请求
// alibaba.ailab.user.open.uid.get
//
// access token 获取精灵用户 id
type AlibabaAilabUserOpenUidGetAPIRequest struct {
	model.Params
	// access token
	_skillAccessToken string
	// skill id
	_skillId int64
}

// NewAlibabaAilabUserOpenUidGetRequest 初始化AlibabaAilabUserOpenUidGetAPIRequest对象
func NewAlibabaAilabUserOpenUidGetRequest() *AlibabaAilabUserOpenUidGetAPIRequest {
	return &AlibabaAilabUserOpenUidGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAilabUserOpenUidGetAPIRequest) Reset() {
	r._skillAccessToken = ""
	r._skillId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAilabUserOpenUidGetAPIRequest) GetApiMethodName() string {
	return "alibaba.ailab.user.open.uid.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAilabUserOpenUidGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAilabUserOpenUidGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSkillAccessToken is SkillAccessToken Setter
// access token
func (r *AlibabaAilabUserOpenUidGetAPIRequest) SetSkillAccessToken(_skillAccessToken string) error {
	r._skillAccessToken = _skillAccessToken
	r.Set("skill_access_token", _skillAccessToken)
	return nil
}

// GetSkillAccessToken SkillAccessToken Getter
func (r AlibabaAilabUserOpenUidGetAPIRequest) GetSkillAccessToken() string {
	return r._skillAccessToken
}

// SetSkillId is SkillId Setter
// skill id
func (r *AlibabaAilabUserOpenUidGetAPIRequest) SetSkillId(_skillId int64) error {
	r._skillId = _skillId
	r.Set("skill_id", _skillId)
	return nil
}

// GetSkillId SkillId Getter
func (r AlibabaAilabUserOpenUidGetAPIRequest) GetSkillId() int64 {
	return r._skillId
}

var poolAlibabaAilabUserOpenUidGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAilabUserOpenUidGetRequest()
	},
}

// GetAlibabaAilabUserOpenUidGetRequest 从 sync.Pool 获取 AlibabaAilabUserOpenUidGetAPIRequest
func GetAlibabaAilabUserOpenUidGetAPIRequest() *AlibabaAilabUserOpenUidGetAPIRequest {
	return poolAlibabaAilabUserOpenUidGetAPIRequest.Get().(*AlibabaAilabUserOpenUidGetAPIRequest)
}

// ReleaseAlibabaAilabUserOpenUidGetAPIRequest 将 AlibabaAilabUserOpenUidGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaAilabUserOpenUidGetAPIRequest(v *AlibabaAilabUserOpenUidGetAPIRequest) {
	v.Reset()
	poolAlibabaAilabUserOpenUidGetAPIRequest.Put(v)
}
