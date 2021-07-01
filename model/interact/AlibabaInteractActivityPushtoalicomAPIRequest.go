package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaInteractActivityPushtoalicomAPIRequest
小铺isv推广流量活动到流量聚乐部 API请求
alibaba.interact.activity.pushtoalicom

涉及到流量包的小铺isv，将活动推送到流量聚乐部 */
type AlibabaInteractActivityPushtoalicomAPIRequest struct {
	model.Params
	// 推送到流量聚乐部的banner图
	_bannerUrl string
	// 推送到流量聚乐部的活动bizid
	_bizId string
}

// New
