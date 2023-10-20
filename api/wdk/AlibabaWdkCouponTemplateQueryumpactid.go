package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkcoupontemplatequeryumpactid 通过券模板查询券活动id接口
// alibaba.wdk.coupon.template.queryumpactid
//
// 当前大润发商家根据券模板创建券id，但订单回流的核销是根据券活动id回流的，大润发侧无法建立券模板和券活动的关联关系，导致大润发计算核销率比较困难，营销域增加通过券模板查询券活动id接口
func Alibabawdkcoupontemplatequeryumpactid(clt *core.SDKClient, req *wdk.AlibabawdkcoupontemplatequeryumpactidAPIRequest, session string) (*wdk.AlibabawdkcoupontemplatequeryumpactidAPIResponse, error) {
	var resp wdk.AlibabawdkcoupontemplatequeryumpactidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
