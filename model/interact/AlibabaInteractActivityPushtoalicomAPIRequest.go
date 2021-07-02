package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractActivityPushtoalicomAPIRequest 小铺isv推广流量活动到流量聚乐部 API请求
// alibaba.interact.activity.pushtoalicom
//
// 涉及到流量包的小铺isv，将活动推送到流量聚乐部
type AlibabaInteractActivityPushtoalicomAPIRequest struct {
	model.Params
	// 推送到流量聚乐部的banner图
	_bannerUrl string
	// 推送到流量聚乐部的活动bizid
	_bizId string
}

// NewAlibabaInteractActivityPushtoalicomRequest 初始化AlibabaInteractActivityPushtoalicomAPIRequest对象
func NewAlibabaInteractActivityPushtoalicomRequest() *AlibabaInteractActivityPushtoalicomAPIRequest {
	return &AlibabaInteractActivityPushtoalicomAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractActivityPushtoalicomAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.activity.pushtoalicom"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractActivityPushtoalicomAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetBannerUrl is BannerUrl Setter
// 推送到流量聚乐部的banner图
func (r *AlibabaInteractActivityPushtoalicomAPIRequest) SetBannerUrl(_bannerUrl string) error {
	r._bannerUrl = _bannerUrl
	r.Set("banner_url", _bannerUrl)
	return nil
}

// GetBannerUrl BannerUrl Getter
func (r AlibabaInteractActivityPushtoalicomAPIRequest) GetBannerUrl() string {
	return r._bannerUrl
}

// SetBizId is BizId Setter
// 推送到流量聚乐部的活动bizid
func (r *AlibabaInteractActivityPushtoalicomAPIRequest) SetBizId(_bizId string) error {
	r._bizId = _bizId
	r.Set("biz_id", _bizId)
	return nil
}

// GetBizId BizId Getter
func (r AlibabaInteractActivityPushtoalicomAPIRequest) GetBizId() string {
	return r._bizId
}
