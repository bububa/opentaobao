package icbushowcase

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbushowcase"
)

/* AlibabaScbpShowcaseUpdateproduct
替换橱窗商品
alibaba.scbp.showcase.updateproduct

替换橱窗商品 */
func AlibabaScbpShowcaseUpdateproduct(clt *core.SDKClient, req *icbushowcase.AlibabaScbpShowcaseUpdateproductAPIRequest, session string) (*icbushowcase.AlibabaScbpShowcaseUpdateproductAPIResponse, error) {
	var resp icbushowcase.AlibabaScbpShowcaseUpdateproductAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
