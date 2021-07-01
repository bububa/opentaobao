package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

/* AlibabaWdkReverseCreatrefund
逆向提交
alibaba.wdk.reverse.creatrefund

逆向申请提交 */
func AlibabaWdkReverseCreatrefund(clt *core.SDKClient, req *wdk.AlibabaWdkReverseCreatrefundAPIRequest, session string) (*wdk.AlibabaWdkReverseCreatrefundAPIResponse, error) {
	var resp wdk.AlibabaWdkReverseCreatrefundAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
