package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkCouponTemplateQueryumpactidAPIRequest
通过券模板查询券活动id接口 API请求
alibaba.wdk.coupon.template.queryumpactid

当前大润发商家根据券模板创建券id，但订单回流的核销是根据券活动id回流的，大润发侧无法建立券模板和券活动的关联关系，导致大润发计算核销率比较困难，营销域增加通过券模板查询券活动id接口 */
type AlibabaWdkCouponTemplateQueryumpactidAPIRequest struct {
	model.Params
	// 券模版id列表
	_sourceIds []int64
	// 优惠券类型
	_wdkCouponType int64
}

// New
