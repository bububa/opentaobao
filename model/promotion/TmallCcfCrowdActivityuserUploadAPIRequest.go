package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallCcfCrowdActivityuserUploadAPIRequest
品牌营销活动用户上传 API请求
tmall.ccf.crowd.activityuser.upload

搜集ISV的活动用户信息，将其沉淀为活动人群数据 */
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

// New
