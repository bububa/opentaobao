package tmallservice

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaServicecenterSpserviceorderUpdateAPIRequest
服务供应链服务单更新 API请求
alibaba.servicecenter.spserviceorder.update

服务供应链服务单更新，服务商通过此接口将商品的sn等信息推送到服务单中 */
type AlibabaServicecenterSpserviceorderUpdateAPIRequest struct {
	model.Params
	// 服务单id
	_spServiceOrderId int64
	// 新设备sn.当action填写addSn、changeSn时必填
	_action string
	// 新设备sn.当action填写addSn、changeSn时必填
	_newSn string
	// 旧设备sn，当action填写changeSn时必填
	_oldSn string
	// 服务有效期开始时间
	_gmtEffect string
	// 服务过期时间
	_gmtExpire string
}

// New
