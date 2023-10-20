package usergrowth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaousergrowthadmaterialupdateAPIRequest 素材更新 API请求
// taobao.usergrowth.ad.material.update
//
// 素材更新
type TaobaousergrowthadmaterialupdateAPIRequest struct {
	model.Params
	// 渠道-创意id
	_outerCreativeId string
	// 长文案
	_title string
	// 图片url
	_imgUrl string
	// H5落地页链接
	_h5Url string
	// 视频url
	_videoUrl string
	// 短文案
	_subTitle string
	// dp链接
	_dpUrl string
	// 淘宝频道
	_bizType int64
	// 标的类型
	_scenarioType int64
	// 广告类型
	_adType int64
	// 创意类型
	_creativeType int64
	// 应用id
	_appId int64
	// 渠道id
	_channelId int64
	// 创意id
	_id int64
	// 热点事件
	_hotEvent int64
}

// NewTaobaousergrowthadmaterialupdateRequest 初始化TaobaousergrowthadmaterialupdateAPIRequest对象
func NewTaobaousergrowthadmaterialupdateRequest() *TaobaousergrowthadmaterialupdateAPIRequest {
	return &TaobaousergrowthadmaterialupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaousergrowthadmaterialupdateAPIRequest) GetApiMethodName() string {
	return "taobao.usergrowth.ad.material.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaousergrowthadmaterialupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaousergrowthadmaterialupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOuterCreativeId is OuterCreativeId Setter
// 渠道-创意id
func (r *TaobaousergrowthadmaterialupdateAPIRequest) SetOuterCreativeId(_outerCreativeId string) error {
	r._outerCreativeId = _outerCreativeId
	r.Set("outer_creative_id", _outerCreativeId)
	return nil
}

// GetOuterCreativeId OuterCreativeId Getter
func (r TaobaousergrowthadmaterialupdateAPIRequest) GetOuterCreativeId() string {
	return r._outerCreativeId
}

// SetTitle is Title Setter
// 长文案
func (r *TaobaousergrowthadmaterialupdateAPIRequest) SetTitle(_title string) error {
	r._title = _title
	r.Set("title", _title)
	return nil
}

// GetTitle Title Getter
func (r TaobaousergrowthadmaterialupdateAPIRequest) GetTitle() string {
	return r._title
}

// SetImgUrl is ImgUrl Setter
// 图片url
func (r *TaobaousergrowthadmaterialupdateAPIRequest) SetImgUrl(_imgUrl string) error {
	r._imgUrl = _imgUrl
	r.Set("img_url", _imgUrl)
	return nil
}

// GetImgUrl ImgUrl Getter
func (r TaobaousergrowthadmaterialupdateAPIRequest) GetImgUrl() string {
	return r._imgUrl
}

// SetH5Url is H5Url Setter
// H5落地页链接
func (r *TaobaousergrowthadmaterialupdateAPIRequest) SetH5Url(_h5Url string) error {
	r._h5Url = _h5Url
	r.Set("h5_url", _h5Url)
	return nil
}

// GetH5Url H5Url Getter
func (r TaobaousergrowthadmaterialupdateAPIRequest) GetH5Url() string {
	return r._h5Url
}

// SetVideoUrl is VideoUrl Setter
// 视频url
func (r *TaobaousergrowthadmaterialupdateAPIRequest) SetVideoUrl(_videoUrl string) error {
	r._videoUrl = _videoUrl
	r.Set("video_url", _videoUrl)
	return nil
}

// GetVideoUrl VideoUrl Getter
func (r TaobaousergrowthadmaterialupdateAPIRequest) GetVideoUrl() string {
	return r._videoUrl
}

// SetSubTitle is SubTitle Setter
// 短文案
func (r *TaobaousergrowthadmaterialupdateAPIRequest) SetSubTitle(_subTitle string) error {
	r._subTitle = _subTitle
	r.Set("sub_title", _subTitle)
	return nil
}

// GetSubTitle SubTitle Getter
func (r TaobaousergrowthadmaterialupdateAPIRequest) GetSubTitle() string {
	return r._subTitle
}

// SetDpUrl is DpUrl Setter
// dp链接
func (r *TaobaousergrowthadmaterialupdateAPIRequest) SetDpUrl(_dpUrl string) error {
	r._dpUrl = _dpUrl
	r.Set("dp_url", _dpUrl)
	return nil
}

// GetDpUrl DpUrl Getter
func (r TaobaousergrowthadmaterialupdateAPIRequest) GetDpUrl() string {
	return r._dpUrl
}

// SetBizType is BizType Setter
// 淘宝频道
func (r *TaobaousergrowthadmaterialupdateAPIRequest) SetBizType(_bizType int64) error {
	r._bizType = _bizType
	r.Set("biz_type", _bizType)
	return nil
}

// GetBizType BizType Getter
func (r TaobaousergrowthadmaterialupdateAPIRequest) GetBizType() int64 {
	return r._bizType
}

// SetScenarioType is ScenarioType Setter
// 标的类型
func (r *TaobaousergrowthadmaterialupdateAPIRequest) SetScenarioType(_scenarioType int64) error {
	r._scenarioType = _scenarioType
	r.Set("scenario_type", _scenarioType)
	return nil
}

// GetScenarioType ScenarioType Getter
func (r TaobaousergrowthadmaterialupdateAPIRequest) GetScenarioType() int64 {
	return r._scenarioType
}

// SetAdType is AdType Setter
// 广告类型
func (r *TaobaousergrowthadmaterialupdateAPIRequest) SetAdType(_adType int64) error {
	r._adType = _adType
	r.Set("ad_type", _adType)
	return nil
}

// GetAdType AdType Getter
func (r TaobaousergrowthadmaterialupdateAPIRequest) GetAdType() int64 {
	return r._adType
}

// SetCreativeType is CreativeType Setter
// 创意类型
func (r *TaobaousergrowthadmaterialupdateAPIRequest) SetCreativeType(_creativeType int64) error {
	r._creativeType = _creativeType
	r.Set("creative_type", _creativeType)
	return nil
}

// GetCreativeType CreativeType Getter
func (r TaobaousergrowthadmaterialupdateAPIRequest) GetCreativeType() int64 {
	return r._creativeType
}

// SetAppId is AppId Setter
// 应用id
func (r *TaobaousergrowthadmaterialupdateAPIRequest) SetAppId(_appId int64) error {
	r._appId = _appId
	r.Set("app_id", _appId)
	return nil
}

// GetAppId AppId Getter
func (r TaobaousergrowthadmaterialupdateAPIRequest) GetAppId() int64 {
	return r._appId
}

// SetChannelId is ChannelId Setter
// 渠道id
func (r *TaobaousergrowthadmaterialupdateAPIRequest) SetChannelId(_channelId int64) error {
	r._channelId = _channelId
	r.Set("channel_id", _channelId)
	return nil
}

// GetChannelId ChannelId Getter
func (r TaobaousergrowthadmaterialupdateAPIRequest) GetChannelId() int64 {
	return r._channelId
}

// SetId is Id Setter
// 创意id
func (r *TaobaousergrowthadmaterialupdateAPIRequest) SetId(_id int64) error {
	r._id = _id
	r.Set("id", _id)
	return nil
}

// GetId Id Getter
func (r TaobaousergrowthadmaterialupdateAPIRequest) GetId() int64 {
	return r._id
}

// SetHotEvent is HotEvent Setter
// 热点事件
func (r *TaobaousergrowthadmaterialupdateAPIRequest) SetHotEvent(_hotEvent int64) error {
	r._hotEvent = _hotEvent
	r.Set("hot_event", _hotEvent)
	return nil
}

// GetHotEvent HotEvent Getter
func (r TaobaousergrowthadmaterialupdateAPIRequest) GetHotEvent() int64 {
	return r._hotEvent
}
