package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabahmmarketingexpirepromotiondelete 短保优惠删除
// alibaba.hm.marketing.expire.promotion.delete
//
// 短保优惠删除
func Alibabahmmarketingexpirepromotiondelete(clt *core.SDKClient, req *wdk.AlibabahmmarketingexpirepromotiondeleteAPIRequest, session string) (*wdk.AlibabahmmarketingexpirepromotiondeleteAPIResponse, error) {
	var resp wdk.AlibabahmmarketingexpirepromotiondeleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
