package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallCcfCrowdActivityuserUploadAPIRequest 品牌营销活动用户上传 API请求
// tmall.ccf.crowd.activityuser.upload
//
// 搜集ISV的活动用户信息，将其沉淀为活动人群数据
type TmallCcfCrowdActivityuserUploadAPIRequest struct {
	model.Params
	// 活动id
	_activityId int64
	// 人群类型
	_crowdTypes []string
	// 淘宝小程序的openid
	_taobaoOpenId string
	// 小程序对应的appKey
	_taobaoAppKey string
}

// NewTmallCcfCrowdActivityuserUploadRequest 初始化TmallCcfCrowdActivityuserUploadAPIRequest对象
func NewTmallCcfCrowdActivityuserUploadRequest() *TmallCcfCrowdActivityuserUploadAPIRequest {
	return &TmallCcfCrowdActivityuserUploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallCcfCrowdActivityuserUploadAPIRequest) GetApiMethodName() string {
	return "tmall.ccf.crowd.activityuser.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallCcfCrowdActivityuserUploadAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetActivityId is ActivityId Setter
// 活动id
func (r *TmallCcfCrowdActivityuserUploadAPIRequest) SetActivityId(_activityId int64) error {
	r._activityId = _activityId
	r.Set("activity_id", _activityId)
	return nil
}

// GetActivityId ActivityId Getter
func (r TmallCcfCrowdActivityuserUploadAPIRequest) GetActivityId() int64 {
	return r._activityId
}

// SetCrowdTypes is CrowdTypes Setter
// 人群类型
func (r *TmallCcfCrowdActivityuserUploadAPIRequest) SetCrowdTypes(_crowdTypes []string) error {
	r._crowdTypes = _crowdTypes
	r.Set("crowd_types", _crowdTypes)
	return nil
}

// GetCrowdTypes CrowdTypes Getter
func (r TmallCcfCrowdActivityuserUploadAPIRequest) GetCrowdTypes() []string {
	return r._crowdTypes
}

// SetTaobaoOpenId is TaobaoOpenId Setter
// 淘宝小程序的openid
func (r *TmallCcfCrowdActivityuserUploadAPIRequest) SetTaobaoOpenId(_taobaoOpenId string) error {
	r._taobaoOpenId = _taobaoOpenId
	r.Set("taobao_open_id", _taobaoOpenId)
	return nil
}

// GetTaobaoOpenId TaobaoOpenId Getter
func (r TmallCcfCrowdActivityuserUploadAPIRequest) GetTaobaoOpenId() string {
	return r._taobaoOpenId
}

// SetTaobaoAppKey is TaobaoAppKey Setter
// 小程序对应的appKey
func (r *TmallCcfCrowdActivityuserUploadAPIRequest) SetTaobaoAppKey(_taobaoAppKey string) error {
	r._taobaoAppKey = _taobaoAppKey
	r.Set("taobao_app_key", _taobaoAppKey)
	return nil
}

// GetTaobaoAppKey TaobaoAppKey Getter
func (r TmallCcfCrowdActivityuserUploadAPIRequest) GetTaobaoAppKey() string {
	return r._taobaoAppKey
}
