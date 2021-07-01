package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

/* AlibabaWdkOldposRefundCreate
五道口外部商户老pos机产生的退款单同步进盒马
alibaba.wdk.oldpos.refund.create

淘鲜达外部商户老pos机产生的退款单同步进淘鲜达 */
func AlibabaWdkOldposRefundCreate(clt *core.SDKClient, req *wdk.AlibabaWdkOldposRefundCreateAPIRequest, session string) (*wdk.AlibabaWdkOldposRefundCreateAPIResponse, error) {
	var resp wdk.AlibabaWdkOldposRefundCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
