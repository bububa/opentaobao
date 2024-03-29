package promotion

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallCcfCrowdActivityuserUploadAPIRequest 品牌营销活动用户上传 API请求
// tmall.ccf.crowd.activityuser.upload
//
// 搜集ISV的活动用户信息，将其沉淀为活动人群数据
type TmallCcfCrowdActivityuserUploadAPIRequest struct {
	model.Params
	// 人群类型
	_crowdTypes []string
	// 淘宝小程序的openid
	_taobaoOpenId string
	// 小程序对应的appKey
	_taobaoAppKey string
	// 活动id
	_activityId int64
}

// NewTmallCcfCrowdActivityuserUploadRequest 初始化TmallCcfCrowdActivityuserUploadAPIRequest对象
func NewTmallCcfCrowdActivityuserUploadRequest() *TmallCcfCrowdActivityuserUploadAPIRequest {
	return &TmallCcfCrowdActivityuserUploadAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallCcfCrowdActivityuserUploadAPIRequest) Reset() {
	r._crowdTypes = r._crowdTypes[:0]
	r._taobaoOpenId = ""
	r._taobaoAppKey = ""
	r._activityId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallCcfCrowdActivityuserUploadAPIRequest) GetApiMethodName() string {
	return "tmall.ccf.crowd.activityuser.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallCcfCrowdActivityuserUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallCcfCrowdActivityuserUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolTmallCcfCrowdActivityuserUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallCcfCrowdActivityuserUploadRequest()
	},
}

// GetTmallCcfCrowdActivityuserUploadRequest 从 sync.Pool 获取 TmallCcfCrowdActivityuserUploadAPIRequest
func GetTmallCcfCrowdActivityuserUploadAPIRequest() *TmallCcfCrowdActivityuserUploadAPIRequest {
	return poolTmallCcfCrowdActivityuserUploadAPIRequest.Get().(*TmallCcfCrowdActivityuserUploadAPIRequest)
}

// ReleaseTmallCcfCrowdActivityuserUploadAPIRequest 将 TmallCcfCrowdActivityuserUploadAPIRequest 放入 sync.Pool
func ReleaseTmallCcfCrowdActivityuserUploadAPIRequest(v *TmallCcfCrowdActivityuserUploadAPIRequest) {
	v.Reset()
	poolTmallCcfCrowdActivityuserUploadAPIRequest.Put(v)
}
