package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkmarketingexpirepromotiondelete 短保优惠删除
// alibaba.wdk.marketing.expire.promotion.delete
//
// 短保优惠删除
func Alibabawdkmarketingexpirepromotiondelete(clt *core.SDKClient, req *wdk.AlibabawdkmarketingexpirepromotiondeleteAPIRequest, session string) (*wdk.AlibabawdkmarketingexpirepromotiondeleteAPIResponse, error) {
	var resp wdk.AlibabawdkmarketingexpirepromotiondeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
