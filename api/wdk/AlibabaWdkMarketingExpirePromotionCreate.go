package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkmarketingexpirepromotioncreate 短保优惠创建
// alibaba.wdk.marketing.expire.promotion.create
//
// 过期优惠优惠信息录入
func Alibabawdkmarketingexpirepromotioncreate(clt *core.SDKClient, req *wdk.AlibabawdkmarketingexpirepromotioncreateAPIRequest, session string) (*wdk.AlibabawdkmarketingexpirepromotioncreateAPIResponse, error) {
	var resp wdk.AlibabawdkmarketingexpirepromotioncreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
