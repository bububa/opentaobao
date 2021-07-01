package alihealthcrm

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlihealthBabyBaseinfoGetAPIRequest
三方从我们这获取宝宝信息 API请求
alibaba.alihealth.baby.baseinfo.get

三方从我们这获取宝宝信息 */
type AlibabaAlihealthBabyBaseinfoGetAPIRequest struct {
	model.Params
	// 宝宝id
	_babyId int64
	// 宝宝所属的用户
	_tpUserId int64
}

// New
