package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabahmmarketingexpirepromotioncreate 短保优惠创建
// alibaba.hm.marketing.expire.promotion.create
//
// 过期优惠优惠信息录入
func Alibabahmmarketingexpirepromotioncreate(clt *core.SDKClient, req *wdk.AlibabahmmarketingexpirepromotioncreateAPIRequest, session string) (*wdk.AlibabahmmarketingexpirepromotioncreateAPIResponse, error) {
	var resp wdk.AlibabahmmarketingexpirepromotioncreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
