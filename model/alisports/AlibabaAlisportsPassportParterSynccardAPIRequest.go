package alisports

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlisportsPassportParterSynccardAPIRequest
阿里体育-卡信息同步接口 API请求
alibaba.alisports.passport.parter.synccard

运享通修改卡号的时候，通知更新到阿里体育和支付宝卡包中 */
type AlibabaAlisportsPassportParterSynccardAPIRequest struct {
	model.Params
	// 用户的id
	_aliuid string
	// 类型：1.修改，2.删除
	_type string
	// 用户的老卡号(修改或删除之前的卡号)
	_oldCardNum string
	// 时间戳
	_alispTime string
	// appkey
	_alispAppKey string
	// 签名字符串
	_alispSign string
	// 改卡的中心id，如果卡号唯一则不需要传
	_centerNum string
}

// New
