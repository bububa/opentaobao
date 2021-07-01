package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaInteractActivityApplyAPIRequest
ISV报名官方活动(中心化流量) API请求
alibaba.interact.activity.apply

支持商家将使用isv创建的活动所对应的权益信息同步到手淘，供过滤是否在中心化流量入口透出 */
type AlibabaInteractActivityApplyAPIRequest struct {
	model.Params
	// 活动发放的权益类型，1:支付宝红包2:流量包3:淘金币4:集分宝5:优惠卷
	_benefitType string
	// 权益对应的面额
	_benefitDenomination string
	// 报名参加的中心化流量活动的banner 地址
	_bannerUrl string
	// 报名参加中心化流量活动的活动id
	_activityBizId string
	// 该活动参与的中心化流量活动。1:代表每日赢宝箱2:微淘广场
	_bizType string
	// 权益总额
	_benefitAmount string
	// 权益属性(如红包，则为relationid)
	_benefitAttribute string
}

// New
